// Copyright (c) 2022 Steven Stallion <sstallion@gmail.com>
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

//go:build freebsd || (linux && libusb)

package hid

/*
#include <stdint.h>
#include "hidapi_libusb.h"
*/
import "C"

// InterfaceNbrAny can be passed to the OpenSysDevice function to match any
// USB interface number.
const InterfaceNbrAny = -1

// OpenSysDevice opens the HID device attached to the system using
// libusb_wrap_sys_device. This function wraps a platform-specific file
// descriptor known to libusb and the USB interface number specified by fd and
// ifnum, respectively.
func OpenSysDevice(fd uintptr, ifnum int) (*Device, error) {
	handle := C.hid_libusb_wrap_sys_device(C.intptr_t(fd), C.int(ifnum))
	if handle == nil {
		return nil, wrapErr(Error())
	}
	return &Device{handle}, nil
}
