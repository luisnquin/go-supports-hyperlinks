// Translation of https://github.com/jamestalmage/supports-hyperlinks to Go Programming Language
// by luisnquin(https://github.com/luisnquin)

// About windows support(Quote extracted from https://github.com/jamestalmage/supports-hyperlinks/pull/8):
// "I understand that a properly implemented terminal will just ignore the special characters,
// and show the link text. But that can tend to clobber output formatting and alignments.
// The goal of this package isn't simply to avoid printing control characters (as mentioned,
// a properly implemented terminal won't display those characters anyways), but to allow authors
// to conditionally format their output based on how it will be displayed."

package supports_hyperlinks

import (
	"flag"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	env6 "github.com/caarlos0/env/v6"
	sc "github.com/jwalton/go-supportscolor"
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

	versionSpecs struct{ major, minor, patch int }
)

var flagWasParsed bool

func init() {
	_ = flag.Bool("hyperlink", false, "Enables the hyperlinks support in the current program execution")
	_ = flag.Bool("hyperlinks", false, "Enables the hyperlinks support in the current program execution")
	_ = flag.Bool("no-hyperlink", false, "Disables the hyperlinks support in the current program execution")
	_ = flag.Bool("no-hyperlinks", false, "Disables the hyperlinks support in the current program execution")
}

func Stdout() bool {
	if !flagWasParsed {
		flag.Parse()
		flagWasParsed = true
	}
	return SupportsHyperlinks(os.Stdout)
}

func Stderr() bool {
	if !flagWasParsed {
		flag.Parse()
		flagWasParsed = true
	}
	return SupportsHyperlinks(os.Stderr)
}

func SupportsHyperlinks(stream *os.File) bool {
	var env envVars

	if err := env6.Parse(&env); err != nil {
		panic(err)
	}
	if !flagWasParsed {
		flag.Parse()
		flagWasParsed = true
	}
	if isFlagPassed("hyperlink") || isFlagPassed("hyperlinks") {
		return true
	}
	if isFlagPassed("no-hyperlink") || isFlagPassed("no-hyperlinks") {
		return false
	}
	// For now
	if runtime.GOOS == "windows" {
		return false
	}

	if env.ForceHyperlink != "" {
		f, err := strconv.Atoi(env.ForceHyperlink)
		if err != nil {
			panic(err)
		}
		return f == 1
	}

	if !sc.Stdout().SupportsColor {
		return false
	}

	if env.Netlify != "" {
		return true
	}

	if env.Ci != "" {
		return false
	}

	if env.TeamcityVersion != "" {
		return false
	}

	if env.TermProgram != "" {
		version := parseVersion(env.TermProgramVersion)
		switch {
		case env.TermProgram == "iTerm.app" && version.major == 3:
			return version.minor >= 1
		default:
			return version.major > 3
		}
	}

	if env.VteVersion != "" {
		if env.VteVersion == "0.50.0" {
			return false
		}

		version := parseVersion(env.VteVersion)
		return version.major > 0 || version.minor >= 50
	}
	return false
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

func parseVersion(v string) versionSpecs {
	var (
		vSliced []string
		err     error
	)
	var newV versionSpecs

	rgx, _ := regexp.Compile(`^\d{3,4}$`)
	if rgx.Match([]byte(v)) {
		rgx, _ = regexp.Compile(`(\d{1,2})(\d{2})`)
		vSliced = rgx.FindAllString(v, 2)

		newV.major = 0
		newV.minor, err = strconv.Atoi(vSliced[1])
		if err != nil {
			panic(err)
		}
		newV.patch, err = strconv.Atoi(vSliced[2])
		if err != nil {
			panic(err)
		}
		return newV
	}

	vSliced = strings.Split(v, ".")

	newV.major, err = strconv.Atoi(vSliced[0])
	if err != nil {
		panic(err)
	}
	newV.minor, err = strconv.Atoi(vSliced[1])
	if err != nil {
		panic(err)
	}
	newV.patch, err = strconv.Atoi(vSliced[2])
	if err != nil {
		panic(err)
	}
	return newV
}
