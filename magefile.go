//go:build mage

package main

import (
	"fmt"
	"path"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/mcandre/mx"
	"github.com/mcandre/octane"
)

// artifactsPath describes where artifacts are produced.
const artifactsPath = "bin"

// portBasename labels the artifact basename.
const portBasename = "octane"

// repoNamespace identifies the Go namespace for this project.
const repoNamespace = "github.com/mcandre/octane"

// imageXgo denotes a Docker image for building this project.
const imageXgo = "n4jm4/octane-xgo"

// artifactsPathDist is the parent directory of xgo artifacts.
var artifactsPathDist = path.Join(artifactsPath, portBasename)

// Default references the default build task.
var Default = Test

// Clean deletes build artifacts.
func Clean() error { mg.Deps(CleanArtifacts); return CleanPackages() }

// CleanBin deletes Go artifacts.
func CleanArtifacts() error { return sh.Rm(artifactsPath) }

// CleanPackages deletes OS package artifacts.
func CleanPackages() error { return sh.RunV("rockhopper", "-c") }

// Audit runs security audits.
func Audit() error {
	mg.Deps(Govulncheck)
	return DockerScout()
}

// Deadcode runs deadcode.
func Deadcode() error { return sh.RunV("deadcode", "./...") }

// DockerBuild generates Docker images.
func DockerBuild() error { return sh.RunV("docker", "buildx", "bake", "all") }

// DockerPush pushes Docker images.
func DockerPush() error { return sh.RunV("docker", "buildx", "bake", "production", "--push") }

// DockerScout runs docker scout scans.
func DockerScout() error {
	if err := sh.RunV("docker", "scout", "cves", "-e", imageXgo); err != nil {
		return err
	}

	return sh.RunV("docker", "scout", "cves", "-e", "fs://.")
}

// DockerTest tests pushing Docker images.
func DockerTest() error { return sh.RunV("docker", "buildx", "bake", "test", "--push") }

// Errcheck runs errcheck.
func Errcheck() error { return sh.RunV("errcheck", "-blank") }

// GoFix runs go fix.
func GoFix() error { return sh.RunV("go", "fix", "./...") }

// GoImports runs goimports.
func GoImports() error { return mx.GoImports("-w") }

// GoVet runs default go vet analyzers.
func GoVet() error { return mx.GoVet() }

// Govulncheck runs govulncheck.
func Govulncheck() error { return sh.RunV("govulncheck", "-scan", "package", "./...") }

// Install builds and installs Go applications.
func Install() error { return mx.Install() }

// Lint runs the lint suite.
func Lint() error {
	mg.Deps(Deadcode)
	mg.Deps(GoFix)
	mg.Deps(GoImports)
	mg.Deps(GoVet)
	mg.Deps(Errcheck)
	mg.Deps(Nakedret)
	mg.Deps(Shadow)
	mg.Deps(Staticcheck)
	return nil
}

// Nakedret runs nakedret.
func Nakedret() error { return mx.Nakedret("-l", "0") }

// Package generates OS packages.
func Package() error { return sh.RunV("rockhopper", "-r", fmt.Sprintf("version=%s", octane.Version)) }

// Shadow runs go vet with shadow checks enabled.
func Shadow() error { return mx.GoVetShadow() }

// Staticcheck runs staticcheck.
func Staticcheck() error { return sh.RunV("staticcheck", "./...") }

// Test runs a test suite.
func Test() error { return mx.UnitTest() }

// Uninstall deletes installed Go applications.
func Uninstall() error { return mx.Uninstall("octane") }

// Upload copies packages to CloudFlare R2.
func Upload() error { mg.Deps(Install); return sh.RunV("./upload") }

// Xgo cross-compiles (c)Go binaries with additional targets enabled.
func Xgo() error {
	// Skip 32-bit ports
	// Skip broken ports
	ports := []string{
		"darwin/amd64",
		"darwin/arm64",
		"freebsd/amd64",
		"linux/amd64",
		"linux/arm64",
		"windows/amd64",
		"windows/arm64",
	}

	return sh.RunV(
		"xgo",
		"-dest",
		artifactsPathDist,
		"-image",
		imageXgo,
		"-targets",
		strings.Join(ports, ","),
		".",
	)
}
