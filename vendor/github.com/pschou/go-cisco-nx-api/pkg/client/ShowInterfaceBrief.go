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

type ShowInterfaceBriefResponse struct {
	TableInterface []struct {
		RowInterface []struct {
			Interface    string `json:"interface" xml:"interface"`
			State        string `json:"state,omitempty" xml:"state,omitempty"`
			IPAddr       string `json:"ip_addr,omitempty" xml:"ip_addr,omitempty"`
			Speed        string `json:"speed,omitempty" xml:"speed,omitempty"`
			MTU          int    `json:"mtu,omitempty" xml:"mtu,omitempty"`
			Vlan         string `json:"vlan,omitempty" xml:"vlan,omitempty"`
			Type         string `json:"type,omitempty" xml:"type,omitempty"`
			PortMode     string `json:"portmode,omitempty" xml:"portmode,omitempty"`
			StateRsnDesc string `json:"state_rsn_desc,omitempty" xml:"state_rsn_desc,omitempty"`
			RateMode     string `json:"ratemode,omitempty" xml:"ratemode,omitempty"`
			PortChan     int    `json:"portchan,omitempty" xml:"portchan,omitempty"`
			Proto        string `json:"proto,omitempty" xml:"proto,omitempty"`
			OperState    string `json:"oper_state,omitempty" xml:"oper_state,omitempty"`
			SviRsnDesc   string `json:"svi_rsn_desc,omitempty" xml:"svi_rsn_desc,omitempty"`
		} `json:"ROW_interface" xml:"ROW_interface"`
	} `json:"TABLE_interface" xml:"TABLE_interface"`
}

// NewShowInterfaceBriefFromString returns instance from an input string.
func NewShowInterfaceBriefFromString(s string) (*ShowInterfaceBriefResponse, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowInterfaceBriefFromReader(strings.NewReader(s))
}

// NewShowInterfaceBriefFromBytes returns instance from an input byte array.
func NewShowInterfaceBriefFromBytes(s []byte) (*ShowInterfaceBriefResponse, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowInterfaceBriefFromReader(bytes.NewReader(s))
}

// NewShowInterfaceBriefFromReader returns instance from an input reader.
func NewShowInterfaceBriefFromReader(s io.Reader) (*ShowInterfaceBriefResponse, error) {
	//si := &ShowInterfaceBrief{}
	ShowInterfaceBriefResponseDat := &ShowInterfaceBriefResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowInterfaceBriefResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowInterfaceBriefResponseDat, nil
}
