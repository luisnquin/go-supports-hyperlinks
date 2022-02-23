// Translation of https://github.com/jamestalmage/supports-hyperlinks to Go Programming Language
// by luisnquin(https://github.com/luisnquin)

// About windows support:
// I understand that a properly implemented terminal will just ignore the special characters,
// and show the link text. But that can tend to clobber output formatting and alignments.
// The goal of this package isn't simply to avoid printing control characters (as mentioned,
// a properly implemented terminal won't display those characters anyways), but to allow authors
// to conditionally format their output based on how it will be displayed.

package supports_hyperlinks

import (
	"flag"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/caarlos0/env/v6"
	"github.com/jwalton/go-supportscolor"
)

type (
	envVars struct {
		TermProgramVersion string `env:"TERM_PROGRAM_VERSION"`
		TeamcityVersion    string `env:"TEAMCITY_VERSION"`
		ForceHyperlink     string `env:"FORCE_HYPERLINK"`
		TermProgram        string `env:"TERM_PROGRAM"`
		ForceColor         string `env:"FORCE_COLOR"`
		VteVersion         string `env:"VTE_VERSION"`
		Netlify            string `env:"NETLIFY"`
		Ci                 string `env:"CI"`
	}

	vSpec struct{ major, minor, patch int }
)

var (
	Stdout = SupportsHyperlinks(os.Stdout)
	Stderr = SupportsHyperlinks(os.Stderr)
)

func SupportsHyperlinks(stream *os.File) bool {
	envir := envVars{}

	if err := env.Parse(&envir); err != nil {
		panic(err)
	}

	if isFlagPassed("no-hyperlink") || isFlagPassed("no-hyperlinks") {
		return false
	}

	if isFlagPassed("hyperlink") {
		return true
	}

	// For now
	if runtime.GOOS == "windows" {
		return false
	}

	if envir.ForceHyperlink != "" {
		f, err := strconv.Atoi(envir.ForceHyperlink)
		checkErr(err)

		return f == 1
	}

	if !supportscolor.Stdout().SupportsColor {
		return false
	}

	if envir.Netlify != "" {
		return true
	}

	if envir.Ci != "" {
		return false
	}

	if envir.TeamcityVersion != "" {
		return false
	}

	if envir.TermProgram != "" {
		version := parseVersion(envir.TermProgramVersion)
		switch {
		case envir.TermProgram == "iTerm.app" && version.major == 3:
			return version.minor >= 1
		default:
			// Originally greater than 3
			return version.major >= 3
		}
	}

	if envir.VteVersion != "" {
		if envir.VteVersion == "0.50.0" {
			return false
		}

		version := parseVersion(envir.VteVersion)
		return version.major > 0 || version.minor >= 50
	}
	return false
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func isFlagPassed(name string) bool {
	var passed bool

	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			passed = true
		}
	})
	return passed
}

func parseVersion(version string) vSpec {
	var (
		v      vSpec
		vSpecs []string
		err    error
	)

	rgx, _ := regexp.Compile(`^\d{3,4}$`)
	if rgx.Match([]byte(version)) {
		rgx, _ = regexp.Compile(`(\d{1,2})(\d{2})`)
		vSpecs = rgx.FindAllString(version, 2)

		v.major = 0
		v.minor, err = strconv.Atoi(vSpecs[1])
		checkErr(err)
		v.patch, err = strconv.Atoi(vSpecs[2])
		checkErr(err)

		return v
	}

	vSpecs = strings.Split(version, ".")

	v.major, err = strconv.Atoi(vSpecs[0])
	checkErr(err)
	v.minor, err = strconv.Atoi(vSpecs[1])
	checkErr(err)
	v.patch, err = strconv.Atoi(vSpecs[2])
	checkErr(err)

	return v
}
