variable "VERSION" {
    default = "test"
}

variable "PLATFORMS" {
    # Drop 32-bit support
    # Work around buildx quirks
    default = [
        # "linux/386",
        "linux/amd64",
        # "linux/arm/v6",
        # "linux/arm/v7",
        "linux/arm64/v8",
        # "linux/ppc64le",
        # "linux/riscv64",
        # "linux/s390x",
    ]
}

variable "PRODUCTION" {
    default = [
        "octane-xgo",
    ]
}

variable "TEST" {
    default = [
        "test-octane-xgo",
    ]
}

group "production" {
    targets = PRODUCTION
}

group "test" {
    targets = TEST
}

group "all" {
    targets = concat(PRODUCTION, TEST)
}

target "octane-xgo" {
    platforms = PLATFORMS
    tags = [
        "n4jm4/octane-xgo",
    ]
}

target "test-octane-xgo" {
    platforms = PLATFORMS
    tags = [
        "n4jm4/octane-xgo:test",
    ]
}
