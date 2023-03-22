# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.19+ with `go install github.com/mcandre/accio/cmd/accio@v0.0.3`, `accio -install`, and `modvendor -copy='**/*.h **/*.c **/*.hpp **/*.cpp'` re-run after every `go mod vendor` execution
* [Docker](https://www.docker.com/) 19+
* a [C++](https://en.wikipedia.org/wiki/List_of_compilers#C++_compilers) compiler
* [Node.js](https://nodejs.org/en) 16.14.2+ with `npm install -g snyk@1.996.0`

## Recommended

* [ASDF](https://asdf-vm.com/) 0.10
* [direnv](https://direnv.net/) 2

## Linux

Linux build environments have additional requirements.

* [ALSA](https://alsa-project.org/wiki/Main_Page) development headers (Debian: `libasound2-dev`, RHEL: `alsa-lib-devel`, Alpine: `alsa-lib-dev`, etc.)

# CGO BUILD ERRORS

```console
$ go install ./...
file not found
```

You forgot to run the `modvendor -copy='**/*.h **/*.c **/*.hpp **/*.cpp'` after a `go mod vendor` run. This is an artifact of a messed up CGO project upstream.

# AUDIT

```console
$ mage audit
```

# INSTALL

```console
$ mage install
```

# UNINSTALL

```console
$ mage uninstall
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
