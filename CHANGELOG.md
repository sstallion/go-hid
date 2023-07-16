# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.14.1] - 2023-07-15

### Fixed

- Fixed MBCS string support for `windows` ([MURAOKA Taro](https://github.com/koron)).

## [0.14.0] - 2023-06-05

### Added

- Added missing optional APIs for `windows`.

### Changed

- Imported 0.14.0 sources from [libusb/hidapi](https://github.com/libusb/hidapi).

## [0.13.3] - 2023-02-27

### Added

- Added libusb backend support for `linux`.

## [0.13.2] - 2023-01-31

### Changed

- Improved `lshid` device information formatting.

### Fixed

- `go.mod` now correctly references `go1.17`.

## [0.13.1] - 2023-01-09

### Changed

- Updated `lshid` flags; `-verbose` has been renamed to `-V`.
- Imported 0.13.1 sources from [libusb/hidapi](https://github.com/libusb/hidapi).

## [0.13.0] - 2023-01-06

### Changed

- Imported 0.13.0 sources from [libusb/hidapi](https://github.com/libusb/hidapi).

## [0.12.4] - 2022-12-23

### Fixed

- Guaranteed `hid.Device` satisfies `io.ReadWriteCloser`.

## [0.12.3] - 2022-12-20

### Added

- Added `/usr/local/lib` to library search path on FreeBSD.

## [0.12.2] - 2022-12-20

### Changed

- Renamed `hid.InterfaceAny` to `hid.InterfaceNbrAny` for consistency.

## [0.12.1] - 2022-12-19

### Added

- Added missing optional APIs for `darwin` and `freebsd`.
- Improved error reporting.

## [0.12.0] - 2022-12-18

### Changed

- Imported 0.12.0 sources from [libusb/hidapi](https://github.com/libusb/hidapi).

## [0.8.0] - 2022-12-18

Historical release based on the original HIDAPI, updated to support Go Modules.
See https://github.com/signal11/hidapi for details.

[Unreleased]: https://github.com/sstallion/go-hid/compare/v0.14.1...HEAD
[0.14.1]: https://github.com/sstallion/go-hid/releases/tag/v0.14.1
[0.14.0]: https://github.com/sstallion/go-hid/releases/tag/v0.14.0
[0.13.3]: https://github.com/sstallion/go-hid/releases/tag/v0.13.3
[0.13.2]: https://github.com/sstallion/go-hid/releases/tag/v0.13.2
[0.13.1]: https://github.com/sstallion/go-hid/releases/tag/v0.13.1
[0.13.0]: https://github.com/sstallion/go-hid/releases/tag/v0.13.0
[0.12.4]: https://github.com/sstallion/go-hid/releases/tag/v0.12.4
[0.12.3]: https://github.com/sstallion/go-hid/releases/tag/v0.12.3
[0.12.2]: https://github.com/sstallion/go-hid/releases/tag/v0.12.2
[0.12.1]: https://github.com/sstallion/go-hid/releases/tag/v0.12.1
[0.12.0]: https://github.com/sstallion/go-hid/releases/tag/v0.12.0
[0.8.0]: https://github.com/sstallion/go-hid/releases/tag/v0.8.0
