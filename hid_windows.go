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
#include <stdlib.h>
#include "hidapi_winapi.h"
*/
import "C"

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

// GetContainerID returns the container ID pointed to by guid and an error, if
// any.
func (d *Device) GetContainerID(guid *windows.GUID) error {
	container_id := (*C.GUID)(unsafe.Pointer(guid))
	if res := C.hid_winapi_get_container_id(d.handle, container_id); res == -1 {
		return wrapErr(d.Error())
	}
	return nil
}

// ReconstructDescriptorData reconstructs a HID Report Descriptor from a Win32
// HIDP_PREPARSED_DATA structure pointed to by data. It returns the number of
// bytes reconstructed and an error, if any.
func ReconstructDescriptorData(data interface{}, p []byte) (int, error) {
	hidp_preparsed_data := unsafe.Pointer(&data)
	buf := (*C.uchar)(&p[0])
	buf_size := C.size_t(len(p))

	res := C.hid_winapi_descriptor_reconstruct_pp_data(hidp_preparsed_data, buf, buf_size)
	if res == -1 {
		return int(res), wrapErr(Error())
	}
	return int(res), nil
}
