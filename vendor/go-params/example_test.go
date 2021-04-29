// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// These examples demonstrate more intricate uses of the params package.
package params_test

import (
	"errors"
	"fmt"
	"github.com/pschou/go-params"
	"strings"
	"time"
)

// Example 1: A single string params called "species" with default value "gopher".
var species = params.String("species", "gopher", "the species we are studying", "TEXT")

// Example 2: Two paramss sharing a variable, so we can have a shorthand.
// The order of initialization is undefined, so make sure both use the
// same default value. They must be set up with an init function.
var gopherType string

func init() {
	const (
		defaultGopher = "pocket"
		usage         = "the variety of gopher"
	)
	params.StringVar(&gopherType, "gopher_type", defaultGopher, usage, "TYPE")
	params.StringVar(&gopherType, "g", defaultGopher, usage+" (shorthand)", "")
}

// Example 3: A user-defined params type, a slice of durations.
type interval []time.Duration

// String is the method to format the params's value, part of the params.Value interface.
// The String method's output will be used in diagnostics.
func (i *interval) String() string {
	return fmt.Sprint(*i)
}

// Set is the method to set the params value, part of the params.Value interface.
// Set's argument is a string to be parsed to set the params.
// It's a comma-separated list, so we split it.
func (i *interval) Set(value []string) error {
	// If we wanted to allow the params to be set multiple times,
	// accumulating values, we would delete this if statement.
	// That would permit usages such as
	//	-deltaT 10s -deltaT 15s
	// and other combinations.
	if len(*i) > 0 {
		return errors.New("interval params already set")
	}
	for _, dt := range strings.Split(value[0], ",") {
		duration, err := time.ParseDuration(dt)
		if err != nil {
			return err
		}
		*i = append(*i, duration)
	}
	return nil
}

// Define a params to accumulate durations. Because it has a special type,
// we need to use the Var function and therefore create the params during
// init.

var intervalFlag interval

func init() {
	// Tie the command-line params to the intervalFlag variable and
	// set a usage message.
	params.Var(&intervalFlag, "deltaT", "comma-separated list of intervals to use between events", "VALS", 1)
}

func Example() {
	// All the interesting pieces are with the variables declared above, but
	// to enable the params package to see the paramss defined there, one must
	// execute, typically at the start of main (not init!):
	//	params.Parse()
	// We don't run it here because this is not a main function and
	// the testing suite has already parsed the paramss.
}
