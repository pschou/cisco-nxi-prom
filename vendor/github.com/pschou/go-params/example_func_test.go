// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package params_test

import (
	"errors"
	"fmt"
	"github.com/pschou/go-params"
	"net"
	"os"
)

func ExampleFunc() {
	fs := params.NewFlagSet("ExampleFunc", params.ContinueOnError)
	fs.SetOutput(os.Stdout)
	var ip net.IP
	// Allow both flags -h and --ip by using "h ip" instead of "ip"
	fs.FlagFunc("ip", "`IP address` to parse", "ADDR", 1, func(s []string) error {
		ip = net.ParseIP(s[0])
		if ip == nil {
			return errors.New("could not parse IP")
		}
		return nil
	})
	fs.Parse([]string{"--ip", "127.0.0.1"})
	fmt.Printf("{ip: %v, loopback: %t}\n\n", ip, ip.IsLoopback())

	// 256 is not a valid IPv4 component
	fs.Parse([]string{"--ip", "256.0.0.1"})
	fmt.Printf("{ip: %v, loopback: %t}\n\n", ip, ip.IsLoopback())

	// Output:
	// {ip: 127.0.0.1, loopback: true}
	//
	// invalid value "256.0.0.1" for parameter --ip: could not parse IP
	// Option:
	// --ip ADDR  `IP address` to parse  (Default: "")
	// {ip: <nil>, loopback: false}
}
