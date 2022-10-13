# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.17+
* [accio](https://github.com/mcandre/accio) v0.0.2
* [Docker](https://www.docker.com/) 19+
* a [C++](https://en.wikipedia.org/wiki/List_of_compilers#C++_compilers) compiler

## Recommended

* [ASDF](https://asdf-vm.com/) 0.10
* [snyk](https://www.npmjs.com/package/snyk) 1.996.0 (`npm install -g snyk@1.996.0`)

## Linux

Linux build environments have additional requirements.

* [ALSA](https://alsa-project.org/wiki/Main_Page) development headers (Debian: `libasound2-dev`, RHEL: `alsa-lib-devel`, Alpine: `alsa-lib-dev`, etc.)

# SECURITY AUDIT

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
