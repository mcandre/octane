# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.17+
* [accio](https://github.com/mcandre/accio) v0.0.2
* [Docker](https://www.docker.com/) 19+
* a [C++](https://en.wikipedia.org/wiki/List_of_compilers#C++_compilers) compiler

## Linux

Linux build environments have additional requirements.

* [ALSA](https://alsa-project.org/wiki/Main_Page) development headers (Debian: `libasound2-dev`, RHEL: `alsa-lib-devel`, Alpine: `alsa-lib-dev`, etc.)

# INSTALL

```
$ go install ./...
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
