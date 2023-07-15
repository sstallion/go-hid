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

//go:build !windows

package hid

/*
#include <stdbool.h>
#include <stdlib.h>
#include <wchar.h>

// Several functions declared in wchar.h return errors that are fundamentally
// incompatible with cgo's type system. The following function may be used to
// determine if an error occurred.
static bool
iswerr(size_t n)
{
	return n == (size_t)-1;
}
*/
import "C"

import (
	"unsafe"
)

// gotowcs converts a Go string to a C wide string. The returned string must
// be freed by the caller by calling C.free.
func gotowcs(s string) *C.wchar_t {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))

	n := C.size_t(len(s)) + 1
	wcs := (*C.wchar_t)(calloc(n, C.sizeof_wchar_t))

	if n, err := C.mbstowcs(wcs, cs, n); C.iswerr(n) {
		C.free(unsafe.Pointer(wcs))
		panic(err)
	}
	return wcs
}

// wcstogo converts a C wide string to a Go string.
func wcstogo(wcs *C.wchar_t) string {
	if wcs == nil {
		return ""
	}

	n := C.wcslen(wcs) + 1
	cs := (*C.char)(calloc(n, C.size_t(C.MB_CUR_MAX)))
	defer C.free(unsafe.Pointer(cs))

	if n, err := C.wcstombs(cs, wcs, n); C.iswerr(n) {
		panic(err)
	}
	return C.GoString(cs)
}
