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

type ShowVpcResponse struct {
	InsAPI struct {
		Outputs struct {
			Output ShowVpcResponseResult `json:"output" xml:"output"`
		} `json:"outputs" xml:"outputs"`
		Sid     string `json:"sid" xml:"sid"`
		Type    string `json:"type" xml:"type"`
		Version string `json:"version" xml:"version"`
	} `json:"ins_api" xml:"ins_api"`
}

type ShowVpcResponseResult struct {
	Body  ShowVpcResultBody `json:"body" xml:"body"`
	Code  string            `json:"code" xml:"code"`
	Input string            `json:"input" xml:"input"`
	Msg   string            `json:"msg" xml:"msg"`
}

type ShowVpcResultBody struct {
	VpcDomainID                       string `json:"vpc-domain-id" xml:"vpc-domain-id"`
	VpcPeerStatus                     string `json:"vpc-peer-status" xml:"vpc-peer-status"`
	VpcPeerStatusReason               string `json:"vpc-peer-status-reason" xml:"vpc-peer-status-reason"`
	VpcPeerKeepaliveStatus            string `json:"vpc-peer-keepalive-status" xml:"vpc-peer-keepalive-status"`
	VpcPeerConsistency                string `json:"vpc-peer-consistency" xml:"vpc-peer-consistency"`
	VpcPerVlanPeerConsistency         string `json:"vpc-per-vlan-peer-consistency" xml:"vpc-per-vlan-peer-consistency"`
	VpcPeerConsistencyStatus          string `json:"vpc-peer-consistency-status" xml:"vpc-peer-consistency-status"`
	VpcType2Consistency               string `json:"vpc-type-2-consistency" xml:"vpc-type-2-consistency"`
	VpcType2ConsistencyStatus         string `json:"vpc-type-2-consistency-status" xml:"vpc-type-2-consistency-status"`
	VpcRole                           string `json:"vpc-role" xml:"vpc-role"`
	NumOfVpcs                         int    `json:"num-of-vpcs" xml:"num-of-vpcs"`
	PeerGateway                       string `json:"peer-gateway" xml:"peer-gateway"`
	DualActiveExcludedVlans           string `json:"dual-active-excluded-vlans" xml:"dual-active-excluded-vlans"`
	VpcGracefulConsistencyCheckStatus string `json:"vpc-graceful-consistency-check-status" xml:"vpc-graceful-consistency-check-status"`
	VpcAutoRecoveryStatus             string `json:"vpc-auto-recovery-status" xml:"vpc-auto-recovery-status"`
	VpcDelayRestoreStatus             string `json:"vpc-delay-restore-status" xml:"vpc-delay-restore-status"`
	VpcDelayRestoreSviStatus          string `json:"vpc-delay-restore-svi-status" xml:"vpc-delay-restore-svi-status"`
	OperationalL3Peer                 string `json:"operational-l3-peer" xml:"operational-l3-peer"`
	VpcPeerLinkHdr                    string `json:"vpc-peer-link-hdr" xml:"vpc-peer-link-hdr"`
	TablePeerlink                     []struct {
		RowPeerlink []struct {
			PeerLinkID        string `json:"peer-link-id" xml:"peer-link-id"`
			PeerlinkIfindex   string `json:"peerlink-ifindex" xml:"peerlink-ifindex"`
			PeerLinkPortState string `json:"peer-link-port-state" xml:"peer-link-port-state"`
			PeerUpVlanBitset  string `json:"peer-up-vlan-bitset" xml:"peer-up-vlan-bitset"`
		} `json:"ROW_peerlink" xml:"ROW_peerlink"`
	} `json:"TABLE_peerlink" xml:"TABLE_peerlink"`
	VpcEnd   []string `json:"vpc-end" xml:"vpc-end"`
	VpcHdr   string   `json:"vpc-hdr" xml:"vpc-hdr"`
	VpcNotEs string   `json:"vpc-not-es" xml:"vpc-not-es"`
	TableVpc []struct {
		RowVpc []struct {
			VpcID                int    `json:"vpc-id" xml:"vpc-id"`
			VpcIfindex           string `json:"vpc-ifindex" xml:"vpc-ifindex"`
			VpcPortState         string `json:"vpc-port-state" xml:"vpc-port-state"`
			PhyPortIfRemoved     string `json:"phy-port-if-removed" xml:"phy-port-if-removed"`
			VpcThruPeerlink      string `json:"vpc-thru-peerlink" xml:"vpc-thru-peerlink"`
			VpcConsistency       string `json:"vpc-consistency" xml:"vpc-consistency"`
			VpcConsistencyStatus string `json:"vpc-consistency-status" xml:"vpc-consistency-status"`
			UpVlanBitset         string `json:"up-vlan-bitset" xml:"up-vlan-bitset"`
			EsAttr               string `json:"es-attr" xml:"es-attr"`
		} `json:"ROW_vpc" xml:"ROW_vpc"`
	} `json:"TABLE_vpc" xml:"TABLE_vpc"`
}

// NewShowVpcFromString returns instance from an input string.
func NewShowVpcFromString(s string) (*ShowVpcResponse, error) { if len(s) == 0 { return nil, fmt.Errorf("missing result") }
	return NewShowVpcFromReader(strings.NewReader(s))
}

// NewShowVpcFromBytes returns instance from an input byte array.
func NewShowVpcFromBytes(s []byte) (*ShowVpcResponse, error) { if len(s) == 0 { return nil, fmt.Errorf("missing result") }
	return NewShowVpcFromReader(bytes.NewReader(s))
}

// NewShowVpcFromReader returns instance from an input reader.
func NewShowVpcFromReader(s io.Reader) (*ShowVpcResponse, error) {
	//si := &ShowVpc{}
	ShowVpcResponseDat := &ShowVpcResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowVpcResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowVpcResponseDat, nil
}

// NewShowVpcResultFromString returns instance from an input string.
func NewShowVpcResultFromString(s string) (*ShowVpcResponseResult, error) { if len(s) == 0 { return nil, fmt.Errorf("missing result") }
	return NewShowVpcResultFromReader(strings.NewReader(s))
}

// NewShowVpcResultFromBytes returns instance from an input byte array.
func NewShowVpcResultFromBytes(s []byte) (*ShowVpcResponseResult, error) { if len(s) == 0 { return nil, fmt.Errorf("missing result") }
	return NewShowVpcResultFromReader(bytes.NewReader(s))
}

// NewShowVpcResultFromReader returns instance from an input reader.
func NewShowVpcResultFromReader(s io.Reader) (*ShowVpcResponseResult, error) {
	//si := &ShowVpcResponseResult{}
	ShowVpcResponseResultDat := &ShowVpcResponseResult{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowVpcResponseResultDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowVpcResponseResultDat, nil
}
