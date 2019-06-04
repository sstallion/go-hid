# HIDAPI Bindings for Go

![](doc/gopher.png)

[![](https://travis-ci.org/sstallion/go-hid.svg?branch=master)][1]
[![](https://godoc.org/github.com/sstallion/go-hid?status.svg)][2]
[![](https://goreportcard.com/badge/github.com/sstallion/go-hid)][3]
[![](https://img.shields.io/github/license/sstallion/go-hid.svg)][4]

Package `hid` provides an idiomatic interface to HIDAPI, a simple library for
communicating with USB and Bluetooth HID devices on Linux, Mac, and Windows.

See https://github.com/signal11/hidapi for details.

## Documentation

Up-to-date documentation can be found on [GoDoc][2], or by issuing the `go doc`
command after installing the package:

    $ go doc -all github.com/sstallion/go-hid

## Installation

Package `hid` may be installed using one of two methods:

1. Via the `go get` command, which dynamically links against the system HIDAPI
   installation. This method requires HIDAPI be installed using a system package
   (eg. `hidapi-devel`) and headers are available. In practice, this works well
   for Linux and Mac, but can cause issues on Windows where HIDAPI is not
   commonly packaged:

       $ go get github.com/sstallion/go-hid

2. Use the provided Makefile to statically link against a vendored copy of
   HIDAPI (commit [a6a622f]). This method works for all supported OSes and is
   the suggested method if installing on Windows:

       $ go get -d github.com/sstallion/go-hid
       $ cd $GOPATH/src/github.com/sstallion/go-hid
       $ make all
       $ make install

   [a6a622f]: https://github.com/signal11/hidapi/commit/a6a622ffb680c55da0de787ff93b80280498330f

Note: The prerequisites for HIDAPI must also be installed regardless of the
method chosen above. See the HIDAPI [README][5] for details.

### lshid

An example command named `lshid` is provided, which displays information about
HID devices attached to the system. `lshid` may be installed by issuing:

    $ go get github.com/sstallion/go-hid/cmd/lshid

Once installed, `lshid -h` may be issued to display usage.

## Contributing

Pull requests are welcome. If a problem is encountered using this package,
please file an issue on [GitHub][6].

[1]: https://travis-ci.org/sstallion/go-hid
[2]: https://godoc.org/github.com/sstallion/go-hid
[3]: https://goreportcard.com/report/github.com/sstallion/go-hid
[4]: https://github.com/sstallion/go-hid/blob/master/LICENSE
[5]: https://github.com/signal11/hidapi/blob/master/README.txt
[6]: https://github.com/sstallion/go-hid/issues/new
