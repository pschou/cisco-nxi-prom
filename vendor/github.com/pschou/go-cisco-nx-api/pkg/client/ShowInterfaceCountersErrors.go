// Copyright 2018 Paul Greenberg (greenpau@outlook.com)
//            and Paul Schou     (github.com/pschou)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"bytes"
	"fmt"
	"github.com/pschou/go-json"
	"io"
	"strings"
)

type ShowInterfaceCountersErrorsResponse struct {
	InsAPI struct {
		Outputs struct {
			Output ShowInterfaceCountersErrorsResponseResult `json:"output" xml:"output"`
		} `json:"outputs" xml:"outputs"`
		Sid     string `json:"sid" xml:"sid"`
		Type    string `json:"type" xml:"type"`
		Version string `json:"version" xml:"version"`
	} `json:"ins_api" xml:"ins_api"`
}

type ShowInterfaceCountersErrorsResponseResult struct {
	Body  ShowInterfaceCountersErrorsResultBody `json:"body" xml:"body"`
	Code  string                                `json:"code" xml:"code"`
	Input string                                `json:"input" xml:"input"`
	Msg   string                                `json:"msg" xml:"msg"`
}

type ShowInterfaceCountersErrorsResultBody struct {
	TableInterface []struct {
		RowInterface []struct {
			Interface    string `json:"interface" xml:"interface"`
			EthAlignErr  int    `json:"eth_align_err" xml:"eth_align_err"`
			EthFcsErr    int    `json:"eth_fcs_err" xml:"eth_fcs_err"`
			EthXmitErr   int    `json:"eth_xmit_err" xml:"eth_xmit_err"`
			EthRcvErr    int    `json:"eth_rcv_err" xml:"eth_rcv_err"`
			EthUndersize int    `json:"eth_undersize" xml:"eth_undersize"`
			EthOutdisc   int    `json:"eth_outdisc" xml:"eth_outdisc"`
		} `json:"ROW_interface" xml:"ROW_interface"`
	} `json:"TABLE_interface" xml:"TABLE_interface"`
}

// NewShowInterfaceCountersErrorsFromString returns instance from an input string.
func NewShowInterfaceCountersErrorsFromString(s string) (*ShowInterfaceCountersErrorsResponse, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowInterfaceCountersErrorsFromReader(strings.NewReader(s))
}

// NewShowInterfaceCountersErrorsFromBytes returns instance from an input byte array.
func NewShowInterfaceCountersErrorsFromBytes(s []byte) (*ShowInterfaceCountersErrorsResponse, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowInterfaceCountersErrorsFromReader(bytes.NewReader(s))
}

// NewShowInterfaceCountersErrorsFromReader returns instance from an input reader.
func NewShowInterfaceCountersErrorsFromReader(s io.Reader) (*ShowInterfaceCountersErrorsResponse, error) {
	//si := &ShowInterfaceCountersErrors{}
	ShowInterfaceCountersErrorsResponseDat := &ShowInterfaceCountersErrorsResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowInterfaceCountersErrorsResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowInterfaceCountersErrorsResponseDat, nil
}

// NewShowInterfaceCountersErrorsResultFromString returns instance from an input string.
func NewShowInterfaceCountersErrorsResultFromString(s string) (*ShowInterfaceCountersErrorsResponseResult, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowInterfaceCountersErrorsResultFromReader(strings.NewReader(s))
}

// NewShowInterfaceCountersErrorsResultFromBytes returns instance from an input byte array.
func NewShowInterfaceCountersErrorsResultFromBytes(s []byte) (*ShowInterfaceCountersErrorsResponseResult, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowInterfaceCountersErrorsResultFromReader(bytes.NewReader(s))
}

// NewShowInterfaceCountersErrorsResultFromReader returns instance from an input reader.
func NewShowInterfaceCountersErrorsResultFromReader(s io.Reader) (*ShowInterfaceCountersErrorsResponseResult, error) {
	//si := &ShowInterfaceCountersErrorsResponseResult{}
	ShowInterfaceCountersErrorsResponseResultDat := &ShowInterfaceCountersErrorsResponseResult{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowInterfaceCountersErrorsResponseResultDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowInterfaceCountersErrorsResponseResultDat, nil
}
