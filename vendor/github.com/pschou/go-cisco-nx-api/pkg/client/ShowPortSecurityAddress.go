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

type ShowPortSecurityAddressResponse struct {
	InsAPI struct {
		Outputs struct {
			Output ShowPortSecurityAddressResponseResult `json:"output" xml:"output"`
		} `json:"outputs" xml:"outputs"`
		Sid     string `json:"sid" xml:"sid"`
		Type    string `json:"type" xml:"type"`
		Version string `json:"version" xml:"version"`
	} `json:"ins_api" xml:"ins_api"`
}

type ShowPortSecurityAddressResponseResult struct {
	Body  ShowPortSecurityAddressResultBody `json:"body" xml:"body"`
	Code  string                            `json:"code" xml:"code"`
	Input string                            `json:"input" xml:"input"`
	Msg   string                            `json:"msg" xml:"msg"`
}

type ShowPortSecurityAddressResultBody struct {
	TotalAddr               int `json:"total_addr" xml:"total_addr"`
	MaxSysLimit             int `json:"max_sys_limit" xml:"max_sys_limit"`
	TableEthPortSecMacAddrs []struct {
		RowEthPortSecMacAddrs []struct {
			IfIndex      string `json:"if_index" xml:"if_index"`
			VlanID       int    `json:"vlan_id" xml:"vlan_id"`
			Type         string `json:"type" xml:"type"`
			MacAddr      string `json:"mac_addr" xml:"mac_addr"`
			RemainAge    int    `json:"remain_age" xml:"remain_age"`
			RemoteLearnt int    `json:"remote_learnt" xml:"remote_learnt"`
			RemoteAged   int    `json:"remote_aged" xml:"remote_aged"`
			NumElems     []int  `json:"num_elems" xml:"num_elems"`
			CmdAddrIndex []int  `json:"cmd_addr_index" xml:"cmd_addr_index"`
		} `json:"ROW_eth_port_sec_mac_addrs" xml:"ROW_eth_port_sec_mac_addrs"`
	} `json:"TABLE_eth_port_sec_mac_addrs" xml:"TABLE_eth_port_sec_mac_addrs"`
}

func (d *ShowPortSecurityAddressResponse) Flat() (out []ShowPortSecurityAddressResultFlat) {
	return d.InsAPI.Outputs.Output.Flat()
}
func (d *ShowPortSecurityAddressResponseResult) Flat() (out []ShowPortSecurityAddressResultFlat) {
	for _, Tv := range d.Body.TableEthPortSecMacAddrs {
		for _, Rv := range Tv.RowEthPortSecMacAddrs {
			out = append(out, ShowPortSecurityAddressResultFlat{
				IfIndex:      Rv.IfIndex,
				VlanID:       Rv.VlanID,
				Type:         Rv.Type,
				MacAddr:      Rv.MacAddr,
				RemainAge:    Rv.RemainAge,
				RemoteLearnt: Rv.RemoteLearnt,
				RemoteAged:   Rv.RemoteAged,
				NumElems:     Rv.NumElems,
				CmdAddrIndex: Rv.CmdAddrIndex,
			})
		}
	}
	return
}

type ShowPortSecurityAddressResultFlat struct {
	IfIndex      string `json:"if_index" xml:"if_index"`
	VlanID       int    `json:"vlan_id" xml:"vlan_id"`
	Type         string `json:"type" xml:"type"`
	MacAddr      string `json:"mac_addr" xml:"mac_addr"`
	RemainAge    int    `json:"remain_age" xml:"remain_age"`
	RemoteLearnt int    `json:"remote_learnt" xml:"remote_learnt"`
	RemoteAged   int    `json:"remote_aged" xml:"remote_aged"`
	NumElems     []int  `json:"num_elems" xml:"num_elems"`
	CmdAddrIndex []int  `json:"cmd_addr_index" xml:"cmd_addr_index"`
}

// NewShowPortSecurityAddressFromString returns instance from an input string.
func NewShowPortSecurityAddressFromString(s string) (*ShowPortSecurityAddressResponse, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowPortSecurityAddressFromReader(strings.NewReader(s))
}

// NewShowPortSecurityAddressFromBytes returns instance from an input byte array.
func NewShowPortSecurityAddressFromBytes(s []byte) (*ShowPortSecurityAddressResponse, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowPortSecurityAddressFromReader(bytes.NewReader(s))
}

// NewShowPortSecurityAddressFromReader returns instance from an input reader.
func NewShowPortSecurityAddressFromReader(s io.Reader) (*ShowPortSecurityAddressResponse, error) {
	//si := &ShowPortSecurityAddress{}
	ShowPortSecurityAddressResponseDat := &ShowPortSecurityAddressResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowPortSecurityAddressResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowPortSecurityAddressResponseDat, nil
}

// NewShowPortSecurityAddressResultFromString returns instance from an input string.
func NewShowPortSecurityAddressResultFromString(s string) (*ShowPortSecurityAddressResponseResult, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowPortSecurityAddressResultFromReader(strings.NewReader(s))
}

// NewShowPortSecurityAddressResultFromBytes returns instance from an input byte array.
func NewShowPortSecurityAddressResultFromBytes(s []byte) (*ShowPortSecurityAddressResponseResult, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowPortSecurityAddressResultFromReader(bytes.NewReader(s))
}

// NewShowPortSecurityAddressResultFromReader returns instance from an input reader.
func NewShowPortSecurityAddressResultFromReader(s io.Reader) (*ShowPortSecurityAddressResponseResult, error) {
	//si := &ShowPortSecurityAddressResponseResult{}
	ShowPortSecurityAddressResponseResultDat := &ShowPortSecurityAddressResponseResult{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowPortSecurityAddressResponseResultDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowPortSecurityAddressResponseResultDat, nil
}
