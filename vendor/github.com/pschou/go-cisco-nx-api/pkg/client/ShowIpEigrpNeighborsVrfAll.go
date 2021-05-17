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

type ShowIpEigrpNeighborsVrfAllResponse struct {
	InsAPI struct {
		Outputs struct {
			Output ShowIpEigrpNeighborsVrfAllResponseResult `json:"output" xml:"output"`
		} `json:"outputs" xml:"outputs"`
		Sid     string `json:"sid" xml:"sid"`
		Type    string `json:"type" xml:"type"`
		Version string `json:"version" xml:"version"`
	} `json:"ins_api" xml:"ins_api"`
}

type ShowIpEigrpNeighborsVrfAllResponseResult struct {
	Body  ShowIpEigrpNeighborsVrfAllResultBody `json:"body" xml:"body"`
	Code  string                               `json:"code" xml:"code"`
	Input string                               `json:"input" xml:"input"`
	Msg   string                               `json:"msg" xml:"msg"`
}

type ShowIpEigrpNeighborsVrfAllResultBody struct {
	TableAsn []struct {
		RowAsn []struct {
			Asn      string `json:"asn" xml:"asn"`
			TableVrf []struct {
				RowVrf []struct {
					Vrf       string `json:"vrf" xml:"vrf"`
					TablePeer []struct {
						RowPeer []struct {
							PeerHandle     string   `json:"peer_handle" xml:"peer_handle"`
							PeerIpaddr     string   `json:"peer_ipaddr" xml:"peer_ipaddr"`
							PeerIfname     string   `json:"peer_ifname" xml:"peer_ifname"`
							PeerHoldtime   uint     `json:"peer_holdtime" xml:"peer_holdtime"`
							PeerSrtt       uint     `json:"peer_srtt" xml:"peer_srtt"`
							PeerRto        uint     `json:"peer_rto" xml:"peer_rto"`
							PeerXmitqCount uint     `json:"peer_xmitq_count" xml:"peer_xmitq_count"`
							PeerLastSeqno  uint     `json:"peer_last_seqno" xml:"peer_last_seqno"`
							PeerUptime     Duration `json:"peer_uptime" xml:"peer_uptime"`
						} `json:"ROW_peer" xml:"ROW_peer"`
					} `json:"TABLE_peer" xml:"TABLE_peer"`
				} `json:"ROW_vrf" xml:"ROW_vrf"`
			} `json:"TABLE_vrf" xml:"TABLE_vrf"`
		} `json:"ROW_asn" xml:"ROW_asn"`
	} `json:"TABLE_asn" xml:"TABLE_asn"`
}

// NewShowIpEigrpNeighborsVrfAllFromString returns instance from an input string.
func NewShowIpEigrpNeighborsVrfAllFromString(s string) (*ShowIpEigrpNeighborsVrfAllResponse, error) { if len(s) == 0 { return nil, fmt.Errorf("missing result") }
	return NewShowIpEigrpNeighborsVrfAllFromReader(strings.NewReader(s))
}

// NewShowIpEigrpNeighborsVrfAllFromBytes returns instance from an input byte array.
func NewShowIpEigrpNeighborsVrfAllFromBytes(s []byte) (*ShowIpEigrpNeighborsVrfAllResponse, error) { if len(s) == 0 { return nil, fmt.Errorf("missing result") }
	return NewShowIpEigrpNeighborsVrfAllFromReader(bytes.NewReader(s))
}

// NewShowIpEigrpNeighborsVrfAllFromReader returns instance from an input reader.
func NewShowIpEigrpNeighborsVrfAllFromReader(s io.Reader) (*ShowIpEigrpNeighborsVrfAllResponse, error) {
	//si := &ShowIpEigrpNeighborsVrfAll{}
	ShowIpEigrpNeighborsVrfAllResponseDat := &ShowIpEigrpNeighborsVrfAllResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowIpEigrpNeighborsVrfAllResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowIpEigrpNeighborsVrfAllResponseDat, nil
}

// NewShowIpEigrpNeighborsVrfAllResultFromString returns instance from an input string.
func NewShowIpEigrpNeighborsVrfAllResultFromString(s string) (*ShowIpEigrpNeighborsVrfAllResponseResult, error) { if len(s) == 0 { return nil, fmt.Errorf("missing result") }
	return NewShowIpEigrpNeighborsVrfAllResultFromReader(strings.NewReader(s))
}

// NewShowIpEigrpNeighborsVrfAllResultFromBytes returns instance from an input byte array.
func NewShowIpEigrpNeighborsVrfAllResultFromBytes(s []byte) (*ShowIpEigrpNeighborsVrfAllResponseResult, error) { if len(s) == 0 { return nil, fmt.Errorf("missing result") }
	return NewShowIpEigrpNeighborsVrfAllResultFromReader(bytes.NewReader(s))
}

// NewShowIpEigrpNeighborsVrfAllResultFromReader returns instance from an input reader.
func NewShowIpEigrpNeighborsVrfAllResultFromReader(s io.Reader) (*ShowIpEigrpNeighborsVrfAllResponseResult, error) {
	//si := &ShowIpEigrpNeighborsVrfAllResponseResult{}
	ShowIpEigrpNeighborsVrfAllResponseResultDat := &ShowIpEigrpNeighborsVrfAllResponseResult{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowIpEigrpNeighborsVrfAllResponseResultDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowIpEigrpNeighborsVrfAllResponseResultDat, nil
}
