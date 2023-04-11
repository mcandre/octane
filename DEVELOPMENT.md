# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.20.2+
* a [C++](https://en.wikipedia.org/wiki/List_of_compilers#C++_compilers) compiler
* [Node.js](https://nodejs.org/en) 16.14.2+
* [Rust](https://www.rust-lang.org/) 1.68.2+
* [Docker](https://www.docker.com/) 19+
* a POSIX compliant [make](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/make.html) implementation (e.g. GNU make, BSD make, etc.)
* Provision additional dev tools with `make`
* Re-run `modvendor -copy='**/*.h **/*.c **/*.hpp **/*.cpp'` after `go mod`... commands, to work around a quirk in how the upstream gomidi Cgo project is structured.

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

# BUILD + PUSH DOCKER IMAGE

```console
$ mage dockerBuild
$ mage dockerPush
```

# PORT

```console
$ mage port
```
