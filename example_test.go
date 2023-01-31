// Copyright (c) 2019 Steven Stallion <sstallion@gmail.com>
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

package hid_test

import (
	"fmt"
	"log"

	"github.com/sstallion/go-hid"
)

// The following example was adapted from the HIDAPI documentation to
// demonstrate use of the hid package to communicate with a simple device.
func Example() {
	b := make([]byte, 65)

	// Initialize the hid package.
	if err := hid.Init(); err != nil {
		log.Fatal(err)
	}

	// Open the device using the VID and PID.
	d, err := hid.OpenFirst(0x4d8, 0x3f)
	if err != nil {
		log.Fatal(err)
	}

	// Read the Manufacturer String.
	s, err := d.GetMfrStr()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Manufacturer String: %s\n", s)

	// Read the Product String.
	s, err = d.GetProductStr()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Product String: %s\n", s)

	// Read the Serial Number String.
	s, err = d.GetSerialNbr()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Serial Number String: %s\n", s)

	// Read Indexed String 1.
	s, err = d.GetIndexedStr(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Indexed String 1: %s\n", s)

	// Toggle LED (cmd 0x80). The first byte is the report number (0x0).
	b[0] = 0x0
	b[1] = 0x80
	if _, err := d.Write(b); err != nil {
		log.Fatal(err)
	}

	// Request state (cmd 0x81). The first byte is the report number (0x0).
	b[0] = 0x0
	b[1] = 0x81
	if _, err := d.Write(b); err != nil {
		log.Fatal(err)
	}

	// Read requested state.
	if _, err := d.Read(b); err != nil {
		log.Fatal(err)
	}

	// Print out the returned buffer.
	for i, value := range b[0:3] {
		fmt.Printf("b[%d]: %d\n", i, value)
	}

	// Finalize the hid package.
	if err := hid.Exit(); err != nil {
		log.Fatal(err)
	}
}

// The following example demonstrates use of the Enumerate function to display
// device information for all HID devices attached to the system.
func ExampleEnumerate() {
	hid.Enumerate(hid.VendorIDAny, hid.ProductIDAny, func(info *hid.DeviceInfo) error {
		fmt.Printf("%s: ID %04x:%04x %s %s\n",
			info.Path,
			info.VendorID,
			info.ProductID,
			info.MfrStr,
			info.ProductStr)
		return nil
	})
}
