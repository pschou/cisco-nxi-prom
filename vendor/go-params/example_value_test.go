// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package params_test

import (
	"fmt"
	"github.com/pschou/go-params"
	"net/url"
)

type URLValue struct {
	URL *url.URL
}

func (v URLValue) String() string {
	if v.URL != nil {
		return v.URL.String()
	}
	return ""
}

func (v URLValue) Set(s []string) error {
	if u, err := url.Parse(s[0]); err != nil {
		return err
	} else {
		*v.URL = *u
	}
	return nil
}

var u = &url.URL{}

func ExampleValue() {
	fs := params.NewFlagSet("ExampleValue", params.ExitOnError)
	fs.Var(&URLValue{u}, "url", "The URL to parse", "ADDRESS", 1)

	fs.Parse([]string{"--url", "https://golang.org/pkg/flag/"})
	fmt.Printf(`{scheme: %q, host: %q, path: %q}`, u.Scheme, u.Host, u.Path)

	// Output:
	// {scheme: "https", host: "golang.org", path: "/pkg/flag/"}
}
