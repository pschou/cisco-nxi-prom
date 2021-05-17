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

type ShowCdpNeighborsResponse struct {
	InsAPI struct {
		Outputs struct {
			Output ShowCdpNeighborsResponseResult `json:"output" xml:"output"`
		} `json:"outputs" xml:"outputs"`
		Sid     string `json:"sid" xml:"sid"`
		Type    string `json:"type" xml:"type"`
		Version string `json:"version" xml:"version"`
	} `json:"ins_api" xml:"ins_api"`
}

type ShowCdpNeighborsResponseResult struct {
	Body  ShowCdpNeighborsResultBody `json:"body" xml:"body"`
	Code  string                     `json:"code" xml:"code"`
	Input string                     `json:"input" xml:"input"`
	Msg   string                     `json:"msg" xml:"msg"`
}

type ShowCdpNeighborsResultBody struct {
	NeighCount                int `json:"neigh_count" xml:"neigh_count"`
	TableCdpNeighborBriefInfo []struct {
		RowCdpNeighborBriefInfo []struct {
			Ifindex    int      `json:"ifindex" xml:"ifindex"`
			DeviceID   string   `json:"device_id" xml:"device_id"`
			IntfID     string   `json:"intf_id" xml:"intf_id"`
			TTL        int      `json:"ttl" xml:"ttl"`
			Capability []string `json:"capability" xml:"capability"`
			PlatformID string   `json:"platform_id" xml:"platform_id"`
			PortID     string   `json:"port_id" xml:"port_id"`
		} `json:"ROW_cdp_neighbor_brief_info" xml:"ROW_cdp_neighbor_brief_info"`
	} `json:"TABLE_cdp_neighbor_brief_info" xml:"TABLE_cdp_neighbor_brief_info"`
}

// NewShowCdpNeighborsFromString returns instance from an input string.
func NewShowCdpNeighborsFromString(s string) (*ShowCdpNeighborsResponse, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowCdpNeighborsFromReader(strings.NewReader(s))
}

// NewShowCdpNeighborsFromBytes returns instance from an input byte array.
func NewShowCdpNeighborsFromBytes(s []byte) (*ShowCdpNeighborsResponse, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowCdpNeighborsFromReader(bytes.NewReader(s))
}

// NewShowCdpNeighborsFromReader returns instance from an input reader.
func NewShowCdpNeighborsFromReader(s io.Reader) (*ShowCdpNeighborsResponse, error) {
	//si := &ShowCdpNeighbors{}
	ShowCdpNeighborsResponseDat := &ShowCdpNeighborsResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowCdpNeighborsResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowCdpNeighborsResponseDat, nil
}

// NewShowCdpNeighborsResultFromString returns instance from an input string.
func NewShowCdpNeighborsResultFromString(s string) (*ShowCdpNeighborsResponseResult, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowCdpNeighborsResultFromReader(strings.NewReader(s))
}

// NewShowCdpNeighborsResultFromBytes returns instance from an input byte array.
func NewShowCdpNeighborsResultFromBytes(s []byte) (*ShowCdpNeighborsResponseResult, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowCdpNeighborsResultFromReader(bytes.NewReader(s))
}

// NewShowCdpNeighborsResultFromReader returns instance from an input reader.
func NewShowCdpNeighborsResultFromReader(s io.Reader) (*ShowCdpNeighborsResponseResult, error) {
	//si := &ShowCdpNeighborsResponseResult{}
	ShowCdpNeighborsResponseResultDat := &ShowCdpNeighborsResponseResult{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowCdpNeighborsResponseResultDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowCdpNeighborsResponseResultDat, nil
}
