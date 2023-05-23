// Copyright (c) 2023 Steven Stallion <sstallion@gmail.com>
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE AUTHOR AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED.  IN NO EVENT SHALL THE AUTHOR OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
// OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
// HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
// LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
// OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
// SUCH DAMAGE.

// Package hid provides an idiomatic interface to HIDAPI, a simple library for
// communicating with USB and Bluetooth HID devices on FreeBSD, Linux, macOS,
// and Windows.
//
// See https://github.com/libusb/hidapi for details.
package hid

/*
#cgo darwin LDFLAGS: -framework IOKit -framework CoreFoundation -framework AppKit
#cgo freebsd CFLAGS: -I/usr/local/include
#cgo freebsd LDFLAGS: -L/usr/local/lib -lusb -liconv -pthread
#cgo linux LDFLAGS: -ludev -lrt
#cgo linux,libusb pkg-config: libusb-1.0
#cgo linux,libusb LDFLAGS: -lpthread

#include <stdint.h>
#include <stdlib.h>
#include "hidapi.h"
*/
import "C"

import (
	"errors"
	"io"
	"math"
	"time"
	"unsafe"
)

// VendorIDAny and ProductIDAny can be passed to the Enumerate function to
// match any vendor or product ID, respectively.
const (
	VendorIDAny  = 0
	ProductIDAny = 0
)

// maxStrLen is the maximum length of a string descriptor (bLength).
const maxStrLen = math.MaxUint8

// ErrTimeout is returned if a blocking operation times out before completing.
var ErrTimeout = errors.New("timeout")

func wrapErr(err error) error {
	if err == nil {
		return errors.New("unspecified error")
	}
	return err
}

// Init initializes the hid package. Calling this function is not strictly
// necessary, however it is recommended for concurrent programs.
func Init() error {
	if res := C.hid_init(); res == -1 {
		return wrapErr(Error())
	}
	return nil
}

// Exit finalizes the hid package. This function should be called after all
// device handles have been closed to avoid memory leaks.
func Exit() error {
	if res := C.hid_exit(); res == -1 {
		return wrapErr(Error())
	}
	return nil
}

// BusType describes the underlying bus type.
type BusType int

//go:generate stringer -type BusType -trimprefix=Bus
const (
	BusUnknown BusType = iota
	BusUSB
	BusBluetooth
	BusI2C
	BusSPI
)

// DeviceInfo describes a HID device attached to the system.
type DeviceInfo struct {
	Path         string  // Platform-Specific Device Path
	VendorID     uint16  // Device Vendor ID
	ProductID    uint16  // Device Product ID
	SerialNbr    string  // Serial Number
	ReleaseNbr   uint16  // Device Version Number
	MfrStr       string  // Manufacturer String
	ProductStr   string  // Product String
	UsagePage    uint16  // Usage Page for Device/Interface
	Usage        uint16  // Usage for Device/Interface
	InterfaceNbr int     // USB Interface Number
	BusType      BusType // Underlying Bus Type
}

func newDeviceInfo(p *C.struct_hid_device_info) *DeviceInfo {
	return &DeviceInfo{
		Path:         C.GoString(p.path),
		VendorID:     uint16(p.vendor_id),
		ProductID:    uint16(p.product_id),
		SerialNbr:    wcstogo(p.serial_number),
		ReleaseNbr:   uint16(p.release_number),
		MfrStr:       wcstogo(p.manufacturer_string),
		ProductStr:   wcstogo(p.product_string),
		UsagePage:    uint16(p.usage_page),
		Usage:        uint16(p.usage),
		InterfaceNbr: int(p.interface_number),
		BusType:      BusType(p.bus_type),
	}
}

// EnumFunc is the type of the function called for each HID device attached to
// the system visited by Enumerate. The information provided by the DeviceInfo
// type can be passed to Open to open the device.
type EnumFunc func(info *DeviceInfo) error

// Enumerate visits each HID device attached to the system with a matching
// vendor and product ID. To match multiple devices, VendorIDAny and
// ProductIDAny can be passed to this function. If an error is returned by
// EnumFunc, Enumerate will return immediately with the original error.
func Enumerate(vid, pid uint16, enumFn EnumFunc) error {
	p := C.hid_enumerate(C.uint16_t(vid), C.uint16_t(pid))
	defer C.hid_free_enumeration(p)

	for p != nil {
		if err := enumFn(newDeviceInfo(p)); err != nil {
			return err
		}
		p = p.next
	}
	return nil
}

// Device is a HID device attached to the system.
type Device struct {
	handle *C.hid_device
}

// Open opens a HID device attached to the system with a matching vendor ID,
// product ID, and serial number. It returns an open device handle and an
// error, if any.
func Open(vid, pid uint16, serial string) (*Device, error) {
	wcs := gotowcs(serial)
	defer C.free(unsafe.Pointer(wcs))

	handle := C.hid_open(C.uint16_t(vid), C.uint16_t(pid), wcs)
	if handle == nil {
		return nil, wrapErr(Error())
	}
	return &Device{handle}, nil
}

// OpenFirst opens the first HID device attached to the system with a matching
// vendor ID, and product ID. It returns an open device handle and an error,
// if any.
func OpenFirst(vid, pid uint16) (*Device, error) {
	handle := C.hid_open(C.uint16_t(vid), C.uint16_t(pid), nil)
	if handle == nil {
		return nil, wrapErr(Error())
	}
	return &Device{handle}, nil
}

// OpenPath opens the HID device attached to the system with the given path.
// It returns an open device handle and an error, if any.
func OpenPath(path string) (*Device, error) {
	cs := C.CString(path)
	defer C.free(unsafe.Pointer(cs))

	handle := C.hid_open_path(cs)
	if handle == nil {
		return nil, wrapErr(Error())
	}
	return &Device{handle}, nil
}

// Write sends an output report with len(p) bytes to the Device. It returns
// the number of bytes written and an error, if any.
//
// The first byte must contain the report ID; 0 should be used for devices
// which only support a single report. Data will be sent over the first OUT
// endpoint if it exists, otherwise the control endpoint will be used.
func (d *Device) Write(p []byte) (int, error) {
	data := (*C.uchar)(&p[0])
	length := C.size_t(len(p))

	res := C.hid_write(d.handle, data, length)
	if res == -1 {
		return int(res), wrapErr(d.Error())
	}
	return int(res), nil
}

// ReadWithTimeout receives an input report with len(p) bytes from the Device
// with the specified timeout. It returns the number of bytes read and an
// error, if any. ReadWithTimeout returns ErrTimeout if the operation timed
// out before completing.
//
// If the device supports multiple reports, the first byte will contain the
// report ID.
func (d *Device) ReadWithTimeout(p []byte, timeout time.Duration) (int, error) {
	data := (*C.uchar)(&p[0])
	length := C.size_t(len(p))
	milliseconds := C.int(timeout / time.Millisecond)

	res := C.hid_read_timeout(d.handle, data, length, milliseconds)
	switch res {
	case -1:
		return int(res), wrapErr(d.Error())
	case 0:
		return int(res), ErrTimeout
	}
	return int(res), nil
}

// Read receives an input report with len(p) bytes from the Device. It returns
// the number of bytes read and an error, if any. Read returns ErrTimeout if
// the operation timed out before completing.
//
// If the device supports multiple reports, the first byte will contain the
// report ID.
func (d *Device) Read(p []byte) (int, error) {
	data := (*C.uchar)(&p[0])
	length := C.size_t(len(p))

	res := C.hid_read(d.handle, data, length)
	switch res {
	case -1:
		return int(res), wrapErr(d.Error())
	case 0:
		return int(res), ErrTimeout
	}
	return int(res), nil
}

// SetNonblock changes the default behavior for Read. If nonblocking is true,
// Read will return immediately with ErrTimeout if data is not available to be
// read from the Device.
func (d *Device) SetNonblock(nonblocking bool) error {
	var nonblock C.int
	if nonblocking {
		nonblock = 1
	}

	res := C.hid_set_nonblocking(d.handle, nonblock)
	if res == -1 {
		return wrapErr(d.Error())
	}
	return nil
}

// SendFeatureReport sends a feature report with len(p) bytes to the Device.
// It returns the number of bytes written and an error, if any.
//
// The first byte must contain the report ID to send. Data will be sent over
// the control endpoint as a Set_Report transfer.
func (d *Device) SendFeatureReport(p []byte) (int, error) {
	data := (*C.uchar)(&p[0])
	length := C.size_t(len(p))

	res := C.hid_send_feature_report(d.handle, data, length)
	if res == -1 {
		return int(res), wrapErr(d.Error())
	}
	return int(res), nil
}

// GetFeatureReport receives a feature report with len(p) bytes from the
// Device. It returns the number of bytes read and an error, if any.
//
// The first byte must contain the report ID to receive.
func (d *Device) GetFeatureReport(p []byte) (int, error) {
	data := (*C.uchar)(&p[0])
	length := C.size_t(len(p))

	res := C.hid_get_feature_report(d.handle, data, length)
	if res == -1 {
		return int(res), wrapErr(d.Error())
	}
	return int(res), nil
}

// GetInputReport receives an input report with len(p) bytes from the Device.
// It returns the number of bytes read and an error, if any.
func (d *Device) GetInputReport(p []byte) (int, error) {
	data := (*C.uchar)(&p[0])
	length := C.size_t(len(p))

	res := C.hid_get_input_report(d.handle, data, length)
	if res == -1 {
		return int(res), wrapErr(d.Error())
	}
	return int(res), nil
}

// Close closes the Device.
func (d *Device) Close() error {
	C.hid_close(d.handle)
	return nil
}

// GetMfrStr returns the manufacturer string descriptor and an error, if any.
func (d *Device) GetMfrStr() (string, error) {
	wcs := (*C.wchar_t)(calloc(maxStrLen+1, C.sizeof_wchar_t))
	defer C.free(unsafe.Pointer(wcs))

	res := C.hid_get_manufacturer_string(d.handle, wcs, maxStrLen)
	if res == -1 {
		return "", wrapErr(d.Error())
	}
	return wcstogo(wcs), nil
}

// GetProductStr returns the product string descriptor and an error, if any.
func (d *Device) GetProductStr() (string, error) {
	wcs := (*C.wchar_t)(calloc(maxStrLen+1, C.sizeof_wchar_t))
	defer C.free(unsafe.Pointer(wcs))

	res := C.hid_get_product_string(d.handle, wcs, maxStrLen)
	if res == -1 {
		return "", wrapErr(d.Error())
	}
	return wcstogo(wcs), nil
}

// GetSerialNbr returns the serial number string descriptor and an error, if any.
func (d *Device) GetSerialNbr() (string, error) {
	wcs := (*C.wchar_t)(calloc(maxStrLen+1, C.sizeof_wchar_t))
	defer C.free(unsafe.Pointer(wcs))

	res := C.hid_get_serial_number_string(d.handle, wcs, maxStrLen)
	if res == -1 {
		return "", wrapErr(d.Error())
	}
	return wcstogo(wcs), nil
}

// GetDeviceInfo returns device information and an error, if any.
func (d *Device) GetDeviceInfo() (*DeviceInfo, error) {
	p := C.hid_get_device_info(d.handle)
	if p == nil {
		return nil, wrapErr(Error())
	}
	return newDeviceInfo(p), nil
}

// GetIndexedStr returns a string descriptor by index and an error, if any.
func (d *Device) GetIndexedStr(index int) (string, error) {
	wcs := (*C.wchar_t)(calloc(maxStrLen+1, C.sizeof_wchar_t))
	defer C.free(unsafe.Pointer(wcs))

	res := C.hid_get_indexed_string(d.handle, C.int(index), wcs, maxStrLen)
	if res == -1 {
		return "", wrapErr(d.Error())
	}
	return wcstogo(wcs), nil
}

// GetReportDescriptor receives a report descriptor with len(p) bytes from the
// Device. It returns the number of bytes read and an error, if any.
func (d *Device) GetReportDescriptor(p []byte) (int, error) {
	data := (*C.uchar)(&p[0])
	length := C.size_t(len(p))

	res := C.hid_get_report_descriptor(d.handle, data, length)
	if res == -1 {
		return int(res), wrapErr(d.Error())
	}
	return int(res), nil
}

// Error returns the last error that occurred on the Device. If no error
// occurred, nil is returned.
func (d *Device) Error() error {
	wcs := C.hid_error(d.handle)
	if wcs == nil {
		return nil // no error
	}
	return errors.New(wcstogo(wcs))
}

var _ io.ReadWriteCloser = (*Device)(nil)

// Error returns the last non-device-specific error that occurred. If no error
// occurred, nil is returned.
func Error() error {
	wcs := C.hid_error(nil)
	if wcs == nil {
		return nil // no error
	}
	return errors.New(wcstogo(wcs))
}

// APIVersion describes the HIDAPI version.
type APIVersion struct {
	Major int // Major version number
	Minor int // Minor version number
	Patch int // Patch version number
}

// GetVersion returns the HIDAPI version.
func GetVersion() APIVersion {
	v := C.hid_version()
	return APIVersion{
		Major: int(v.major),
		Minor: int(v.minor),
		Patch: int(v.patch),
	}
}

// GetVersion returns the HIDAPI version as a string.
func GetVersionStr() string {
	return C.GoString(C.hid_version_str())
}
