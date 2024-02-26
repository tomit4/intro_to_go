module github.com/tomit4/hellogo

go 1.22.0

// allows to replace package mystrings without going to github,
// and instead go to our local machine
replace github.com/tomit4/mystrings => /home/brian/Documents/Code/golang/intro_to_go/mystrings

require (
    github.com/tomit4/mystrings v0.0.0
)
