//go:build mage

package main

import (
	"fmt"
	"os"
	"path"

	"github.com/magefile/mage/mg"
	mageextras "github.com/mcandre/mage-extras"
	"github.com/mcandre/octane"
)

// artifactsPath describes where artifacts are produced.
var artifactsPath = "bin"

// Default references the default build task.
var Default = Test

// portBasename labels the artifact basename.
var portBasename = "octane"

// artifactsPathDist is the parent directory of xgo artifacts.
var artifactsPathDist = path.Join(artifactsPath, portBasename)

// repoNamespace identifies the Go namespace for this project.
var repoNamespace = "github.com/mcandre/octane"

// imageXgo denotes a Docker image for building this project.
var imageXgo = "n4jm4/octane-xgo"

// Clean deletes build artifacts.
func Clean() error { mg.Deps(CleanArtifacts); return CleanPackages() }

// CleanBin deletes Go artifacts.
func CleanArtifacts() error { return os.RemoveAll(artifactsPath) }

// CleanPackages deletes OS package artifacts.
func CleanPackages() error { return mageextras.Run("rockhopper", "-c") }

// Audit runs security audits.
func Audit() error {
	mg.Deps(Govulncheck)
	return DockerScout()
}

// Deadcode runs deadcode.
func Deadcode() error { return mageextras.Run("deadcode", "./...") }

// DockerBuild generates Docker images.
func DockerBuild() error { return mageextras.Run("docker", "buildx", "bake", "all") }

// DockerPush pushes Docker images.
func DockerPush() error { return mageextras.Run("docker", "buildx", "bake", "production", "--push") }

// DockerScout runs docker scout scans.
func DockerScout() error {
	if err := mageextras.Run("docker", "scout", "cves", "-e", imageXgo); err != nil {
		return err
	}

	return mageextras.Run("docker", "scout", "cves", "-e", "fs://.")
}

// DockerTest tests pushing Docker images.
func DockerTest() error { return mageextras.Run("docker", "buildx", "bake", "test", "--push") }

// Errcheck runs errcheck.
func Errcheck() error { return mageextras.Run("errcheck", "-blank") }

// GoFix runs go fix.
func GoFix() error { return mageextras.Run("go", "fix", "./...") }

// GoImports runs goimports.
func GoImports() error { return mageextras.GoImports("-w") }

// GoVet runs default go vet analyzers.
func GoVet() error { return mageextras.GoVet() }

// Govulncheck runs govulncheck.
func Govulncheck() error { return mageextras.Run("govulncheck", "-scan", "package", "./...") }

// Install builds and installs Go applications.
func Install() error { return mageextras.Install() }

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
func Nakedret() error { return mageextras.Nakedret("-l", "0") }

// Package generates OS packages.
func Package() error { return mageextras.Run("rockhopper", "-r", fmt.Sprintf("version=%s", octane.Version)) }

// Shadow runs go vet with shadow checks enabled.
func Shadow() error { return mageextras.GoVetShadow() }

// Staticcheck runs staticcheck.
func Staticcheck() error { return mageextras.Run("staticcheck", "./...") }

// Test runs a test suite.
func Test() error { return mageextras.UnitTest() }

// Uninstall deletes installed Go applications.
func Uninstall() error { return mageextras.Uninstall("octane") }

// Upload copies packages to CloudFlare R2.
func Upload() error { mg.Deps(Install); return mageextras.Run("./upload") }

// Xgo cross-compiles (c)Go binaries with additional targets enabled.
func Xgo() error {
	return mageextras.Xgo(
		artifactsPathDist,
		"-image",
		imageXgo,
		"-targets",
		"darwin/amd64,darwin/arm64,freebsd/amd64,linux/amd64,linux/arm64,windows/amd64,windows/arm64",
		".",
	)
}
