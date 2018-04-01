# goinit
#### A tool to create isolated development environments for golang projects

## Why?
Currently when you develop in golang, you put code inside your `GOPATH`. This is a standard location which you use for golang code, dependencies, compiled binaries, and object files across all your projects.

This limits the capability for code isolation, per-project dependency management, and swappability of toolchains.

## What is goinit?
Goinit packages a few concepts together with a quick command-line tool:
* Golang installation and organization ([gimme](https://github.com/travis-ci/gimme))
* `GOPATH` creation per-project
* Vendoring provided on new environment initialization ([glide](https://glide.sh/
)/[govendor](https://github.com/kardianos/govendor)/more to come)
* Scripts to easily swap in and out `GOPATH` + `GOROOT` env-vars
