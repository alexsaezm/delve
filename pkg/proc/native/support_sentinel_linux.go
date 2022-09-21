// This file is used to detect build on unsupported GOOS/GOARCH combinations.

//go:build linux && !amd64 && !arm64 && !386 && !ppc64le
// +build linux,!amd64,!arm64,!386,!ppc64le

package your_linux_architecture_is_not_supported_by_delve
