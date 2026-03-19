//go:build mage

package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/mcandre/mx"
	"github.com/mcandre/octane"
)

// ArtifactsPath describes where artifacts are produced.
const ArtifactsPath = "bin"

// PortBasename labels the artifact basename.
const PortBasename = "octane"

// RepoNamespace identifies the Go namespace for this project.
const RepoNamespace = "github.com/mcandre/octane"

// ImageXgo denotes a Docker image for building this project.
const ImageXgo = "n4jm4/octane-xgo"

// ArtifactsPathDist is the parent directory of xgo artifacts.
var ArtifactsPathDist = path.Join(ArtifactsPath, PortBasename)

// Default references the default build task.
var Default = Build

// Audit runs security checks.
func Audit() error { return Govulncheck() }

// Build compiles Go projects.
func Build() error {
	dest := ArtifactsPath

	if d, ok := os.LookupEnv("DEST"); ok && d != "" {
		dest = d
	}

	if err := os.MkdirAll(dest, 0755); err != nil {
		return err
	}

	return sh.RunV("go", "build", "-o", dest, "./...")
}

// Clean deletes build artifacts.
func Clean() error { mg.Deps(CleanArtifacts); mg.Deps(CleanBuild); return CleanPackages() }

// CleanBin deletes Go artifacts.
func CleanArtifacts() error { return sh.Rm(ArtifactsPath) }

// CleanBuild removes build artifacts.
func CleanBuild() error { return os.RemoveAll(ArtifactsPath) }

// CleanPackages deletes OS package artifacts.
func CleanPackages() error { return sh.RunV("rockhopper", "-c") }

// Deadcode runs deadcode.
func Deadcode() error { return sh.RunV("deadcode", "./...") }

// DockerBuild generates Docker images.
func DockerBuild() error { return sh.RunV("docker", "buildx", "bake", "all") }

// DockerPush pushes Docker images.
func DockerPush() error { return sh.RunV("docker", "buildx", "bake", "production", "--push") }

// DockerScout runs docker scout scans.
func DockerScout() error {
	if err := sh.RunV("docker", "scout", "cves", "-e", ImageXgo); err != nil {
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

// Bucket stores OS packages
const S3Bucket = "s3://octane"

// Artifacts contains precompiled binaries
var Artifacts = path.Join(".rockhopper", "artifacts")

// Banner identifies the application version.
var Banner = fmt.Sprintf("octane-%s", octane.Version)

// S3Dest stores OS packages for this application version.
var S3Dest = fmt.Sprintf("%s/%s/", S3Bucket, Banner)

// Upload sends packages to CloudFlare R2.
func Upload() error {
	return mx.RunVSilent("aws",
		"--cli-connect-timeout", "1",
		"s3",
		"cp",
		"--recursive",
		Artifacts,
		S3Dest,
	)
}

// Xgo cross-compiles (c)Go binaries with additional targets enabled.
func Xgo() error {
	// Skip 32-bit ports
	// Skip broken ports
	// Skip fringe platforms
	ports := []string{
		"darwin/amd64",
		"darwin/arm64",
		// "freebsd/amd64",
		"linux/amd64",
		"linux/arm64",
		// "windows/amd64",
		// "windows/arm64",
	}

	return sh.RunV(
		"xgo",
		"-dest",
		ArtifactsPathDist,
		"-image",
		ImageXgo,
		"-targets",
		strings.Join(ports, ","),
		".",
	)
}
