# HIDAPI Bindings for Go

![](.github/images/gopher.png)

[![](https://github.com/sstallion/go-hid/actions/workflows/ci.yml/badge.svg?branch=master)][1]
[![](https://pkg.go.dev/badge/github.com/sstallion/go-hid)][2]
[![](https://goreportcard.com/badge/github.com/sstallion/go-hid)][3]
[![](https://img.shields.io/github/v/release/sstallion/go-hid)][4]
[![](https://img.shields.io/github/license/sstallion/go-hid.svg)][5]

Package `hid` provides an idiomatic interface to HIDAPI, a simple library for
communicating with USB and Bluetooth HID devices on FreeBSD, Linux, macOS, and
Windows.

See https://github.com/libusb/hidapi for details.

## Installation

To add package `hid` as a dependency or upgrade to its latest version, issue:

    $ go get github.com/sstallion/go-hid@latest

>**Note**: Prerequisites for building HIDAPI from source must be installed prior
> to issuing `go get`. See [Prerequisites][6] for details.

### libusb Backend Support

On Linux, the hidraw backend is enabled by default. If the libusb backend is
desired, the `libusb` build constraint must be specified:

    $ go build -tags libusb ./...

### lshid

A command named `lshid` is provided, which lists HID devices attached to the
system. `lshid` may be installed by issuing:

    $ go install github.com/sstallion/go-hid/cmd/lshid@latest

Once installed, issue `lshid -h` to show usage.

## Documentation

Up-to-date documentation can be found on [pkg.go.dev][2] or by issuing the `go
doc` command after installation:

    $ go doc -all github.com/sstallion/go-hid

## Contributing

Pull requests are welcome! See [CONTRIBUTING.md][7] for details.

## License

Source code in this repository is licensed under a Simplified BSD License. See
[LICENSE][5] for details.

[1]: https://github.com/sstallion/go-hid/actions/workflows/ci.yml
[2]: https://pkg.go.dev/github.com/sstallion/go-hid
[3]: https://goreportcard.com/report/github.com/sstallion/go-hid
[4]: https://github.com/sstallion/go-hid/releases/latest
[5]: https://github.com/sstallion/go-hid/blob/master/LICENSE
[6]: https://github.com/libusb/hidapi/blob/master/BUILD.md#prerequisites
[7]: https://github.com/sstallion/go-hid/blob/master/CONTRIBUTING.md
