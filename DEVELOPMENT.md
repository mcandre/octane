# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.17+
* a [C++](https://en.wikipedia.org/wiki/List_of_compilers#C++_compilers) compiler

## Linux

Linux build environments have additional requirements.

* [ALSA](https://alsa-project.org/wiki/Main_Page) development headers (Debian: `libasound2-dev`, RHEL: `alsa-lib-devel`, Alpine: `alsa-lib-dev`, etc.)

## Recommended

* [Docker](https://www.docker.com/)
* [Mage](https://magefile.org/) (e.g., `go get github.com/magefile/mage`)
* [xgo](https://github.com/karalabe/xgo) (e.g., `go get github.com/karalabe/xgo`)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)
* [golint](https://github.com/golang/lint) (e.g. `go get github.com/golang/lint/golint`)
* [errcheck](https://github.com/kisielk/errcheck) (e.g. `go get github.com/kisielk/errcheck`)
* [nakedret](https://github.com/alexkohler/nakedret) (e.g. `go get github.com/alexkohler/nakedret`)
* [shadow](golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow) (e.g. `go get -u golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow`)
* [zipc](https://github.com/mcandre/zipc) (e.g. `go get github.com/mcandre/zipc/...`)
* [karp](https://github.com/mcandre/karp) (e.g., `go get github.com/mcandre/karp/...`)

# INSTALL

```
$ GO111MODULE=off go install ./...
```

# LINT

Keep the code tidy:

```console
$ mage lint
```

# PORT

```console
$ mage port
```
