# instant-receive-server

A simple web application for transferring files over a local network.

## Purpose

This application enables easy file transfer within a local network. It is designed for personal use on trusted networks.

## How to Use

1. Run `webapp.go` to start the application. It will launch on port 8080.
2. Access the application through a web browser.
3. **Important:** This app lacks authentication and security features, as it’s intended for local, trusted networks only.
4. For security, the app will automatically shut down after 3 minutes of inactivity.
5. You can build this application as an executable for various platforms (Windows, Mac, Linux) and run it on different environments.

## Build Example

To build for Windows with AMD64 architecture, use the following command:

```bash
GOOS=windows GOARCH=amd64 go build webapp.go
```

## Cross-compilation

To see the full list of supported platforms for cross-compilation, use this command:

```bash
$ go tool dist list
```

This will list all platforms supported by the Go compiler:

```bash
$ go tool dist list
aix/ppc64
android/386
android/amd64
android/arm
android/arm64
darwin/amd64
darwin/arm64
dragonfly/amd64
freebsd/386
freebsd/amd64
freebsd/arm
freebsd/arm64
freebsd/riscv64
illumos/amd64
ios/amd64
ios/arm64
js/wasm
linux/386
linux/amd64
linux/arm
linux/arm64
linux/loong64
linux/mips
linux/mips64
linux/mips64le
linux/mipsle
linux/ppc64
linux/ppc64le
linux/riscv64
linux/s390x
netbsd/386
netbsd/amd64
netbsd/arm
netbsd/arm64
openbsd/386
openbsd/amd64
openbsd/arm
openbsd/arm64
openbsd/ppc64
openbsd/riscv64
plan9/386
plan9/amd64
plan9/arm
solaris/amd64
wasip1/wasm
windows/386
windows/amd64
windows/arm
windows/arm64
```