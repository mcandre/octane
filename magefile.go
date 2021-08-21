//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"path"

	"github.com/magefile/mage/mg"
	"github.com/mcandre/mage-extras"
	"github.com/mcandre/octane"
)

// artifactsPath describes where artifacts are produced.
var artifactsPath = "bin"

// Default references the default build task.
var Default = Install

// GoVet runs go vet with shadow checks enabled.
func GoVet() error { return mageextras.GoVetShadow() }

// GoLint runs golint.
func GoLint() error { return mageextras.GoLint() }

// Gofmt runs gofmt.
func GoFmt() error { return mageextras.GoFmt("-s", "-w") }

// GoImports runs goimports.
func GoImports() error { return mageextras.GoImports("-w") }

// Errcheck runs errcheck.
func Errcheck() error { return mageextras.Errcheck("-blank") }

// Nakedret runs nakedret.
func Nakedret() error { return mageextras.Nakedret("-l", "0") }

// Lint runs the lint suite.
func Lint() error {
	mg.Deps(GoVet)
	mg.Deps(GoLint)
	mg.Deps(GoFmt)
	mg.Deps(GoImports)
	mg.Deps(Errcheck)
	mg.Deps(Nakedret)
	return nil
}

// portBasename labels the artifact basename.
var portBasename = fmt.Sprintf("octane-%s", octane.Version)

// repoNamespace identifies the Go namespace for this project.
var repoNamespace = "github.com/mcandre/octane"

// Xgo cross-compiles (c)Go binaries with additional targets enabled.
func Xgo() error {
	artifactsPathDist := path.Join(artifactsPath, portBasename)

	return mageextras.Xgo(
		artifactsPathDist,
		"-image",
		"mcandre/octane-builder",
		"-targets",
		"darwin/amd64,linux/amd64",
		"github.com/mcandre/octane/cmd/octane",
	)
}

// Port builds and compresses artifacts.
func Port() error { mg.Deps(Xgo); return mageextras.Archive(portBasename, artifactsPath) }

// Install builds and installs Go applications.
func Install() error { return mageextras.Install() }

// Uninstall deletes installed Go applications.
func Uninstall() error { return mageextras.Uninstall("octane") }

// Clean deletes artifacts.
func Clean() error { return os.RemoveAll(artifactsPath) }
