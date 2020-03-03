# octane: MIDI adder

[Hello, Operator?](hello-operator.mp3)

# EXAMPLE

```console
$ octane -list
Polling for MIDI devices...
MIDI IN devices:

* Arturia KeyStep 32
* SQ-1 SEQ IN

MIDI OUT devices:

* Arturia KeyStep 32
* SQ-1 MIDI OUT
* SQ-1 CTRL

$ octane -in 'SQ-1 SEQ IN' -out 'Arturia KeyStep 32' -transposeNote -48
Polling for MIDI devices...
Connected to MIDI IN device: SQ-1 SEQ IN
Connected to MIDI OUT device: Arturia KeyStep 32

(Play a sequence)

#0 [4 d:4] channel.NoteOn channel 0 key 48 velocity 64
#0 [1187 d:1187] channel.NoteOff channel 0 key 48
...
```

See `octane -help` for more options.

# TIPS

* Polling may complete faster when MIDI software such as Arturia Analog Lab is running. Though be careful about such software interfering with your jam sessions.
* In a pinch, many MIDI devices can serve as adapters to reach further MIDI devices, using "thru" options.

# DOCUMENTATION

https://godoc.org/github.com/mcandre/octane

# RUNTIME REQUIREMENTS

(None)

# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.12+
* a [C++](https://en.wikipedia.org/wiki/List_of_compilers#C++_compilers) compiler

## Recommended

* [Docker](https://www.docker.com/)
* [Mage](https://magefile.org/) (e.g., `go get github.com/magefile/mage`)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)
* [golint](https://github.com/golang/lint) (e.g. `go get github.com/golang/lint/golint`)
* [errcheck](https://github.com/kisielk/errcheck) (e.g. `go get github.com/kisielk/errcheck`)
* [nakedret](https://github.com/alexkohler/nakedret) (e.g. `go get github.com/alexkohler/nakedret`)
* [shadow](golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow) (e.g. `go get -u golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow`)
* [zipc](https://github.com/mcandre/zipc) (e.g. `go get github.com/mcandre/zipc/...`)
* [karp](https://github.com/mcandre/karp) (e.g., `go get github.com/mcandre/karp/...`)

# INSTALL FROM REMOTE GIT REPOSITORY

```console
$ go get github.com/mcandre/octane/...
```

(Yes, include the ellipsis as well, it's the magic Go syntax for downloading, building, and installing all components of a package, including any libraries and command line tools.)

# INSTALL FROM LOCAL GIT REPOSITORY

```
$ mkdir -p $GOPATH/src/github.com/mcandre
$ git clone https://github.com/mcandre/octane.git $GOPATH/src/github.com/mcandre/octane
$ cd $GOPATH/src/github.com/mcandre/octane
$ git submodule update --init --recursive
$ go install ./...
```

# LINT

Keep the code tidy:

```console
$ mage lint
```
