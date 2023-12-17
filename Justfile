test-release:
    goreleaser release --skip=validate --skip=publish --snapshot --clean
