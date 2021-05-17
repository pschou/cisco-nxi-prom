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

type ShowIpArpDetailVrfAllResponse struct {
	InsAPI struct {
		Outputs struct {
			Output ShowIpArpDetailVrfAllResponseResult `json:"output" xml:"output"`
		} `json:"outputs" xml:"outputs"`
		Sid     string `json:"sid" xml:"sid"`
		Type    string `json:"type" xml:"type"`
		Version string `json:"version" xml:"version"`
	} `json:"ins_api" xml:"ins_api"`
}

type ShowIpArpDetailVrfAllResponseResult struct {
	Body  ShowIpArpDetailVrfAllResultBody `json:"body" xml:"body"`
	Code  string                          `json:"code" xml:"code"`
	Input string                          `json:"input" xml:"input"`
	Msg   string                          `json:"msg" xml:"msg"`
}

type ShowIpArpDetailVrfAllResultBody struct {
	TableVrf []struct {
		RowVrf []struct {
			VrfNameOut string `json:"vrf-name-out"`
			CntTotal   int    `json:"cnt-total"`
			TableAdj   []struct {
				RowAdj []struct {
					IntfOut   string    `json:"intf-out"`
					IPAddrOut string    `json:"ip-addr-out"`
					TimeStamp TimeStamp `json:"time-stamp"`
					Mac       string    `json:"mac"`
					PhyIntf   string    `json:"phy-intf"`
				} `json:"ROW_adj"`
			} `json:"TABLE_adj"`
		} `json:"ROW_vrf"`
	} `json:"TABLE_vrf"`
}

// NewShowIpArpDetailVrfAllFromString returns instance from an input string.
func NewShowIpArpDetailVrfAllFromString(s string) (*ShowIpArpDetailVrfAllResponse, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowIpArpDetailVrfAllFromReader(strings.NewReader(s))
}

// NewShowIpArpDetailVrfAllFromBytes returns instance from an input byte array.
func NewShowIpArpDetailVrfAllFromBytes(s []byte) (*ShowIpArpDetailVrfAllResponse, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowIpArpDetailVrfAllFromReader(bytes.NewReader(s))
}

// NewShowIpArpDetailVrfAllFromReader returns instance from an input reader.
func NewShowIpArpDetailVrfAllFromReader(s io.Reader) (*ShowIpArpDetailVrfAllResponse, error) {
	//si := &ShowIpArpDetailVrfAll{}
	ShowIpArpDetailVrfAllResponseDat := &ShowIpArpDetailVrfAllResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowIpArpDetailVrfAllResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowIpArpDetailVrfAllResponseDat, nil
}

// NewShowIpArpDetailVrfAllResultFromString returns instance from an input string.
func NewShowIpArpDetailVrfAllResultFromString(s string) (*ShowIpArpDetailVrfAllResponseResult, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowIpArpDetailVrfAllResultFromReader(strings.NewReader(s))
}

// NewShowIpArpDetailVrfAllResultFromBytes returns instance from an input byte array.
func NewShowIpArpDetailVrfAllResultFromBytes(s []byte) (*ShowIpArpDetailVrfAllResponseResult, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowIpArpDetailVrfAllResultFromReader(bytes.NewReader(s))
}

// NewShowIpArpDetailVrfAllResultFromReader returns instance from an input reader.
func NewShowIpArpDetailVrfAllResultFromReader(s io.Reader) (*ShowIpArpDetailVrfAllResponseResult, error) {
	//si := &ShowIpArpDetailVrfAllResponseResult{}
	ShowIpArpDetailVrfAllResponseResultDat := &ShowIpArpDetailVrfAllResponseResult{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowIpArpDetailVrfAllResponseResultDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowIpArpDetailVrfAllResponseResultDat, nil
}
