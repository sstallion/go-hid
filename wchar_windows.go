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

// C.mbstowcs() and C.wcstombs() behave poorly on Windows when used with MBCS
// strings (eg. Chinese, Korean, and Japanese). This is due to several related
// issues working with code pages, encoding, the C runtime, and the compiler.
// See #6 and #7 for details.

/*
#include <wchar.h>
*/
import "C"

import (
	"golang.org/x/sys/windows"
)

// gotowcs converts a Go string to a C wide string. The returned string must
// be freed by the caller by calling C.free.
func gotowcs(s string) *C.wchar_t {
	u16s, err := windows.UTF16FromString(s)
	if err != nil {
		panic(err)
	}
	n := C.size_t(len(u16s))
	wcs := (*C.wchar_t)(calloc(n+1, C.sizeof_wchar_t))
	return C.wmemcpy(wcs, (*C.wchar_t)(&u16s[0]), n)
}

// wcstogo converts a C wide string to a Go string.
func wcstogo(wcs *C.wchar_t) string {
	return windows.UTF16PtrToString((*uint16)(wcs))
}
