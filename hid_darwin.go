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

package hid

/*
#include <stdint.h>
#include "hidapi_darwin.h"
*/
import "C"

// GetLocationID returns the location ID and an error, if any.
func (d *Device) GetLocationID() (uint32, error) {
	var id C.uint32_t

	res := C.hid_darwin_get_location_id(d.handle, &id)
	if res == -1 {
		return uint32(res), wrapErr(d.Error())
	}
	return uint32(id), nil
}

// SetOpenExclusive changes the default behavior for Open. If exclusive is
// false, devices will be opened in non-exclusive mode.
func SetOpenExclusive(exclusive bool) {
	var open_exclusive C.int
	if exclusive {
		open_exclusive = 1
	}
	C.hid_darwin_set_open_exclusive(open_exclusive)
}

// GetOpenExclusive returns if exclusive mode is enabled.
func GetOpenExclusive() bool {
	open_exclusive := C.hid_darwin_get_open_exclusive()
	if open_exclusive == 0 {
		return false
	}
	return true
}

// IsOpenExclusive returns if the device is in exclusive mode and an error, if
// any.
func (d *Device) IsOpenExclusive() (bool, error) {
	res := C.hid_darwin_is_device_open_exclusive(d.handle)
	switch res {
	case -1:
		return false, wrapErr(d.Error())
	case 0:
		return false, nil
	}
	return true, nil
}
