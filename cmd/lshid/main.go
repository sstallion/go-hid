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

// lshid lists HID devices
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sstallion/go-hid"
)

type versionFlag struct{}

func (versionFlag) IsBoolFlag() bool { return true }
func (versionFlag) String() string   { return "" }
func (versionFlag) Set(s string) error {
	fmt.Printf("HIDAPI version %s\n", hid.GetVersionStr())
	os.Exit(0)
	return nil
}

var (
	verboseFlag bool
	vidFlag     uint
	pidFlag     uint
)

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: lshid [options]...")
	fmt.Fprintln(os.Stderr, "List HID devices")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Var(versionFlag{}, "V",
		"Print HIDAPI version and exit")
	flag.BoolVar(&verboseFlag, "v", false,
		"Increase verbosity (show device information)")
	flag.UintVar(&vidFlag, "vid", hid.VendorIDAny,
		"Show only devices with the specified `vendor` ID")
	flag.UintVar(&pidFlag, "pid", hid.ProductIDAny,
		"Show only devices with the specified `product` ID")
	flag.Parse()

	vid, pid := uint16(vidFlag), uint16(pidFlag)
	hid.Enumerate(vid, pid, func(info *hid.DeviceInfo) error {
		fmt.Printf("%s: ID %04x:%04x %s %s\n",
			info.Path, info.VendorID, info.ProductID, info.MfrStr, info.ProductStr)
		if verboseFlag {
			fmt.Println("Device Information:")
			fmt.Printf("\tPath         %s\n", info.Path)
			fmt.Printf("\tVendorID     %#04x\n", info.VendorID)
			fmt.Printf("\tProductID    %#04x\n", info.ProductID)
			fmt.Printf("\tSerialNbr    %s\n", info.SerialNbr)
			fmt.Printf("\tReleaseNbr   %x.%x\n", info.ReleaseNbr>>8, info.ReleaseNbr&0xff)
			fmt.Printf("\tMfrStr       %s\n", info.MfrStr)
			fmt.Printf("\tProductStr   %s\n", info.ProductStr)
			fmt.Printf("\tUsagePage    %#x\n", info.UsagePage)
			fmt.Printf("\tUsage        %#x\n", info.Usage)
			fmt.Printf("\tInterfaceNbr %d\n", info.InterfaceNbr)
			fmt.Printf("\tBusType      %s\n", info.BusType)
			fmt.Println()
		}
		return nil
	})
}
