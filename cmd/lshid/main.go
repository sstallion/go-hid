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

//go:generate doxxer . -h
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/sstallion/go-hid"
	"github.com/sstallion/go-tools/util"
)

type versionFlag struct{}

func (versionFlag) IsBoolFlag() bool { return true }
func (versionFlag) String() string   { return "" }
func (versionFlag) Set(s string) error {
	fmt.Println("HIDAPI version", hid.GetVersionStr())
	os.Exit(0)
	return nil
}

var (
	verboseFlag bool
	vidFlag     uint
	pidFlag     uint
)

func fmtRelease(n uint16) string {
	if n == 0 {
		return "(empty)"
	}
	return fmt.Sprintf("%#04x (%x.%x)", n, n>>8, n&0xff)
}

func fmtString(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return "(empty)"
	}
	return s
}

func usage() {
	util.PrintGlobalUsage(`
Lshid lists HID devices attached to the system.

Usage:

  {{ .Program }} [-V] [-v] [-vid vendor] [-pid product]

Flags:

  {{ call .PrintDefaults }}

Report issues to https://github.com/sstallion/go-hid/issues.
`)
}

func main() {
	flag.Usage = usage
	flag.Var(versionFlag{}, "V", "Print HIDAPI version and exit")
	flag.BoolVar(&verboseFlag, "v", false, "Increase verbosity (show device information)")
	flag.UintVar(&vidFlag, "vid", hid.VendorIDAny, "Show devices with matching `vendor` ID")
	flag.UintVar(&pidFlag, "pid", hid.ProductIDAny, "Show devices with matching `product` ID")
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
			fmt.Printf("\tSerialNbr    %s\n", fmtString(info.SerialNbr))
			fmt.Printf("\tReleaseNbr   %s\n", fmtRelease(info.ReleaseNbr))
			fmt.Printf("\tMfrStr       %s\n", fmtString(info.MfrStr))
			fmt.Printf("\tProductStr   %s\n", fmtString(info.ProductStr))
			fmt.Printf("\tUsagePage    %#04x\n", info.UsagePage)
			fmt.Printf("\tUsage        %#04x\n", info.Usage)
			fmt.Printf("\tInterfaceNbr %d\n", info.InterfaceNbr)
			fmt.Printf("\tBusType      %s\n", info.BusType)
			fmt.Println()
		}
		return nil
	})
}
