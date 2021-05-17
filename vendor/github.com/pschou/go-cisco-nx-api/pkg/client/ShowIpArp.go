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

type ShowIpArpResponse struct {
	InsAPI struct {
		Outputs struct {
			Output ShowIpArpResponseResult `json:"output" xml:"output"`
		} `json:"outputs" xml:"outputs"`
		Sid     string `json:"sid" xml:"sid"`
		Type    string `json:"type" xml:"type"`
		Version string `json:"version" xml:"version"`
	} `json:"ins_api" xml:"ins_api"`
}

type ShowIpArpResponseResult struct {
	Body  ShowIpArpResultBody `json:"body" xml:"body"`
	Code  string              `json:"code" xml:"code"`
	Input string              `json:"input" xml:"input"`
	Msg   string              `json:"msg" xml:"msg"`
}

type ShowIpArpResultBody struct {
	TableVrf []struct {
		RowVrf []struct {
			TableAdj []struct {
				RowAdj []struct {
					Flags      string   `json:"flags" xml:"flags"`
					IntfOut    string   `json:"intf-out" xml:"intf-out"`
					IPAddrOut  string   `json:"ip-addr-out" xml:"ip-addr-out"`
					MAC        string   `json:"mac,omitempty" xml:"mac,omitempty"`
					TimeStamp  Duration `json:"time-stamp" xml:"time-stamp"`
					Incomplete bool     `json:"incomplete,omitempty" xml:"incomplete,omitempty"`
				} `json:"ROW_adj" xml:"ROW_adj"`
			} `json:"TABLE_adj" xml:"TABLE_adj"`
			CntTotal   int    `json:"cnt-total" xml:"cnt-total"`
			VrfNameOut string `json:"vrf-name-out" xml:"vrf-name-out"`
		} `json:"ROW_vrf" xml:"ROW_vrf"`
	} `json:"TABLE_vrf" xml:"TABLE_vrf"`
}

func (d *ShowIpArpResponse) Flat() (out []ShowIpArpResultFlat) {
	return d.InsAPI.Outputs.Output.Flat()
}
func (d *ShowIpArpResponseResult) Flat() (out []ShowIpArpResultFlat) {
	for _, Tv := range d.Body.TableVrf {
		for _, Rv := range Tv.RowVrf {
			for _, Ta := range Rv.TableAdj {
				for _, Ra := range Ta.RowAdj {
					out = append(out, ShowIpArpResultFlat{
						Flags:      Ra.Flags,
						IntfOut:    Ra.IntfOut,
						IPAddrOut:  Ra.IPAddrOut,
						MAC:        Ra.MAC,
						TimeStamp:  Ra.TimeStamp,
						Incomplete: Ra.Incomplete,
						//CntTotal:   Rv.CntTotal,
						VrfNameOut: Rv.VrfNameOut,
					})
				}
			}
		}
	}
	return
}

type ShowIpArpResultFlat struct {
	Flags      string   `json:"flags" xml:"flags"`
	IntfOut    string   `json:"intf-out" xml:"intf-out"`
	IPAddrOut  string   `json:"ip-addr-out" xml:"ip-addr-out"`
	MAC        string   `json:"mac,omitempty" xml:"mac,omitempty"`
	TimeStamp  Duration `json:"time-stamp" xml:"time-stamp"`
	Incomplete bool     `json:"incomplete,omitempty" xml:"incomplete,omitempty"`
	//CntTotal   int           `json:"cnt-total" xml:"cnt-total"`
	VrfNameOut string `json:"vrf-name-out" xml:"vrf-name-out"`
}

// NewShowIpArpFromString returns instance from an input string.
func NewShowIpArpFromString(s string) (*ShowIpArpResponse, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowIpArpFromReader(strings.NewReader(s))
}

// NewShowIpArpFromBytes returns instance from an input byte array.
func NewShowIpArpFromBytes(s []byte) (*ShowIpArpResponse, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowIpArpFromReader(bytes.NewReader(s))
}

// NewShowIpArpFromReader returns instance from an input reader.
func NewShowIpArpFromReader(s io.Reader) (*ShowIpArpResponse, error) {
	//si := &ShowIpArp{}
	ShowIpArpResponseDat := &ShowIpArpResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowIpArpResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowIpArpResponseDat, nil
}

// NewShowIpArpResultFromString returns instance from an input string.
func NewShowIpArpResultFromString(s string) (*ShowIpArpResponseResult, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowIpArpResultFromReader(strings.NewReader(s))
}

// NewShowIpArpResultFromBytes returns instance from an input byte array.
func NewShowIpArpResultFromBytes(s []byte) (*ShowIpArpResponseResult, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowIpArpResultFromReader(bytes.NewReader(s))
}

// NewShowIpArpResultFromReader returns instance from an input reader.
func NewShowIpArpResultFromReader(s io.Reader) (*ShowIpArpResponseResult, error) {
	//si := &ShowIpArpResponseResult{}
	ShowIpArpResponseResultDat := &ShowIpArpResponseResult{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowIpArpResponseResultDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowIpArpResponseResultDat, nil
}
