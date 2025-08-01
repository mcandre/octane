//go:build mage
package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg"
	mageextras "github.com/mcandre/mage-extras"
	"github.com/mcandre/tug"
)

// artifactsPath describes where artifacts are produced.
var artifactsPath = "bin"

// Default references the default build task.
var Default = Test

// Govulncheck runs govulncheck.
func Govulncheck() error { return mageextras.Govulncheck("-scan", "package", "./...") }

// Snyk runs Snyk SCA.
func Snyk() error { return mageextras.SnykTest() }

// DockerPublish publishes demo images.
func DockerPublish() error {
	cmd := exec.Command(
		"tug",
		"-t",
		"mcandre/tug-demo",
		"-exclude-arch",
		"386,arm/v6,arm/v7,ppc64le,riscv64,s390x",
	)
	cmd.Env = os.Environ()
	cmd.Dir = "example"
	return cmd.Run()
}

// DockerScout runs a Docker security audit.
func DockerScout() error { return mageextras.DockerScout("-e", "mcandre/tug-demo") }

// Audit runs security audits.
func Audit() error {
	mg.Deps(Govulncheck)
	mg.Deps(Snyk)
	return DockerScout()
}

// Test runs a test suite.
func Test() error { return mageextras.UnitTest() }

// Deadcode runs deadcode.
func Deadcode() error { return mageextras.Deadcode("./...") }

// Gofmt runs gofmt.
func GoFmt() error { return mageextras.GoFmt("-s", "-w") }

// GoImports runs goimports.
func GoImports() error { return mageextras.GoImports("-w") }

// GoVet runs default go vet analyzers.
func GoVet() error { return mageextras.GoVet() }

// Errcheck runs errcheck.
func Errcheck() error { return mageextras.Errcheck("-blank") }

// Nakedret runs nakedret.
func Nakedret() error { return mageextras.Nakedret("-l", "0") }

// Revive runs revive.
func Revive() error { return mageextras.Revive("-set_exit_status") }

// Shadow runs go vet with shadow checks enabled.
func Shadow() error { return mageextras.GoVetShadow() }

// Staticcheck runs staticcheck.
func Staticcheck() error { return mageextras.Staticcheck() }

// Unmake runs unmake.
func Unmake() error {
	err := mageextras.Unmake(".")

	if err != nil {
		return err
	}

	return mageextras.Unmake("-n", ".")
}

// Lint runs the lint suite.
func Lint() error {
	mg.Deps(Deadcode)
	mg.Deps(GoVet)
	mg.Deps(GoFmt)
	mg.Deps(GoImports)
	mg.Deps(Errcheck)
	mg.Deps(Nakedret)
	mg.Deps(Revive)
	mg.Deps(Shadow)
	mg.Deps(Staticcheck)
	mg.Deps(Unmake)
	return nil
}

// portBasename labels the artifact basename.
var portBasename = fmt.Sprintf("tug-%s", tug.Version)

// repoNamespace identifies the Go namespace for this project.
var repoNamespace = "github.com/mcandre/tug"

// Factorio cross-compiles Go binaries for a multitude of platforms.
func Factorio() error { return mageextras.Factorio(portBasename) }

// Port builds and compresses artifacts.
func Port() error { mg.Deps(Factorio); return mageextras.Archive(portBasename, artifactsPath) }

// Install builds and installs Go applications.
func Install() error { return mageextras.Install() }

// Uninstall deletes installed Go applications.
func Uninstall() error { return mageextras.Uninstall("tug") }

// Clean deletes artifacts.
func Clean() error { return os.RemoveAll(artifactsPath) }
