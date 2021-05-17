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

type ShowNtpPeerStatusResponse struct {
	InsAPI struct {
		Outputs struct {
			Output ShowNtpPeerStatusResponseResult `json:"output" xml:"output"`
		} `json:"outputs" xml:"outputs"`
		Sid     string `json:"sid" xml:"sid"`
		Type    string `json:"type" xml:"type"`
		Version string `json:"version" xml:"version"`
	} `json:"ins_api" xml:"ins_api"`
}

type ShowNtpPeerStatusResponseResult struct {
	Body  ShowNtpPeerStatusResultBody `json:"body" xml:"body"`
	Code  string                      `json:"code" xml:"code"`
	Input string                      `json:"input" xml:"input"`
	Msg   string                      `json:"msg" xml:"msg"`
}

type ShowNtpPeerStatusResultBody struct {
	Totalpeers       string `json:"totalpeers" xml:"totalpeers"`
	TablePeersstatus []struct {
		RowPeersstatus []struct {
			Syncmode string  `json:"syncmode" xml:"syncmode"`
			Remote   string  `json:"remote" xml:"remote"`
			Local    string  `json:"local" xml:"local"`
			St       int     `json:"st" xml:"st"`
			Poll     int     `json:"poll" xml:"poll"`
			Reach    string  `json:"reach" xml:"reach"`
			Delay    float32 `json:"delay" xml:"delay"`
			Vrf      string  `json:"vrf,omitempty" xml:"vrf,omitempty"`
		} `json:"ROW_peersstatus" xml:"ROW_peersstatus"`
	} `json:"TABLE_peersstatus" xml:"TABLE_peersstatus"`
}

type ShowNtpPeerStatusResultFlat struct {
	Syncmode string  `json:"syncmode" xml:"syncmode"`
	Remote   string  `json:"remote" xml:"remote"`
	Local    string  `json:"local" xml:"local"`
	St       int     `json:"st" xml:"st"`
	Poll     int     `json:"poll" xml:"poll"`
	Reach    string  `json:"reach" xml:"reach"`
	Delay    float32 `json:"delay" xml:"delay"`
	Vrf      string  `json:"vrf,omitempty" xml:"vrf,omitempty"`
}

func (d *ShowNtpPeerStatusResponseResult) Flat() (out []ShowNtpPeerStatusResultFlat) {
	for _, Tv := range d.Body.TablePeersstatus {
		for _, Rv := range Tv.RowPeersstatus {
			out = append(out, ShowNtpPeerStatusResultFlat{
				Syncmode: Rv.Syncmode,
				Remote:   Rv.Remote,
				Local:    Rv.Local,
				St:       Rv.St,
				Poll:     Rv.Poll,
				Reach:    Rv.Reach,
				Delay:    Rv.Delay,
				Vrf:      Rv.Vrf,
			})
		}
	}
	return
}

func (d *ShowNtpPeerStatusResponse) Flat() (out []ShowNtpPeerStatusResultFlat) {
	return d.InsAPI.Outputs.Output.Flat()
}

// NewShowNtpPeerStatusFromString returns instance from an input string.
func NewShowNtpPeerStatusFromString(s string) (*ShowNtpPeerStatusResponse, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowNtpPeerStatusFromReader(strings.NewReader(s))
}

// NewShowNtpPeerStatusFromBytes returns instance from an input byte array.
func NewShowNtpPeerStatusFromBytes(s []byte) (*ShowNtpPeerStatusResponse, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowNtpPeerStatusFromReader(bytes.NewReader(s))
}

// NewShowNtpPeerStatusFromReader returns instance from an input reader.
func NewShowNtpPeerStatusFromReader(s io.Reader) (*ShowNtpPeerStatusResponse, error) {
	//si := &ShowNtpPeerStatus{}
	ShowNtpPeerStatusResponseDat := &ShowNtpPeerStatusResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowNtpPeerStatusResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowNtpPeerStatusResponseDat, nil
}

// NewShowNtpPeerStatusResultFromString returns instance from an input string.
func NewShowNtpPeerStatusResultFromString(s string) (*ShowNtpPeerStatusResponseResult, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowNtpPeerStatusResultFromReader(strings.NewReader(s))
}

// NewShowNtpPeerStatusResultFromBytes returns instance from an input byte array.
func NewShowNtpPeerStatusResultFromBytes(s []byte) (*ShowNtpPeerStatusResponseResult, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowNtpPeerStatusResultFromReader(bytes.NewReader(s))
}

// NewShowNtpPeerStatusResultFromReader returns instance from an input reader.
func NewShowNtpPeerStatusResultFromReader(s io.Reader) (*ShowNtpPeerStatusResponseResult, error) {
	//si := &ShowNtpPeerStatusResponseResult{}
	ShowNtpPeerStatusResponseResultDat := &ShowNtpPeerStatusResponseResult{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowNtpPeerStatusResponseResultDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowNtpPeerStatusResponseResultDat, nil
}
