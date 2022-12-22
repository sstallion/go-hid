# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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
See https://github.com/signal11/hidapi for more details.

[Unreleased]: https://github.com/sstallion/go-hid/compare/v0.12.3...HEAD
[0.12.3]: https://github.com/sstallion/go-hid/releases/tag/v0.12.3
[0.12.2]: https://github.com/sstallion/go-hid/releases/tag/v0.12.2
[0.12.1]: https://github.com/sstallion/go-hid/releases/tag/v0.12.1
[0.12.0]: https://github.com/sstallion/go-hid/releases/tag/v0.12.0
[0.8.0]: https://github.com/sstallion/go-hid/releases/tag/v0.8.0
