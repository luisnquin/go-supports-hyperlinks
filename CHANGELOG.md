New: CHANGELOG.md(Yeah mom, me!)

Mod: supports-hyperlinks.go (
    - vSpec type renamed to versionSpecs
    - envir variable from SupportsHyperlinks function, renamed to env | at the same time, the third party package was aliased to env6
    - flagWasParsed to prevent unnecessary parsing
    - Stdout variable has been changed to function, now checks if the parsing process has occurred or not
    - Stderr variable has been changed to function, now checks if the parsing process has occurred or not
    - SupportsHyperlinks function now in the env.Term Program flow per default return the boolean result of (version greater than 3) checking
    - checkErr function removed and patched holes
)

Mod: README.md (
    - Usage section updated
    - Platform support section
    - CLI section better explained
)

Fix: supports-hyperlinks.go (
    - Flags support (
        func init with establishment of flags(just to be on the isFlagPassed function scope)        
    )
)