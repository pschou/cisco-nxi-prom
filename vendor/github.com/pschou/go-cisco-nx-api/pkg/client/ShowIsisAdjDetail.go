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
	//"time"
)

type ShowIsisAdjDetailResponse struct {
	InsAPI struct {
		Outputs struct {
			Output ShowIsisAdjDetailResponseResult `json:"output" xml:"output"`
		} `json:"outputs" xml:"outputs"`
		Sid     string `json:"sid" xml:"sid"`
		Type    string `json:"type" xml:"type"`
		Version string `json:"version" xml:"version"`
	} `json:"ins_api" xml:"ins_api"`
}

type ShowIsisAdjDetailResponseResult struct {
	Body  ShowIsisAdjDetailResultBody `json:"body" xml:"body"`
	Code  string                      `json:"code" xml:"code"`
	Input string                      `json:"input" xml:"input"`
	Msg   string                      `json:"msg" xml:"msg"`
}

type ShowIsisAdjDetailResultBody struct {
	TableProcessTag []struct {
		RowProcessTag []struct {
			ProcessTagOut string `json:"process-tag-out" xml:"process-tag-out"`
			TableVrf      []struct {
				RowVrf []struct {
					VrfNameOut      string `json:"vrf-name-out" xml:"vrf-name-out"`
					AdjSummaryOut   bool   `json:"adj-summary-out" xml:"adj-summary-out"`
					AdjInterfaceOut bool   `json:"adj-interface-out" xml:"adj-interface-out"`
					TableProcessAdj []struct {
						RowProcessAdj []struct {
							AdjSysNameOut              string   `json:"adj-sys-name-out" xml:"adj-sys-name-out"`
							AdjSysIDOut                string   `json:"adj-sys-id-out" xml:"adj-sys-id-out"`
							AdjUsageOut                string   `json:"adj-usage-out" xml:"adj-usage-out"`
							AdjStateOut                string   `json:"adj-state-out" xml:"adj-state-out"`
							AdjHoldTimeOut             Duration `json:"adj-hold-time-out" xml:"adj-hold-time-out"`
							AdjIntfNameOut             string   `json:"adj-intf-name-out" xml:"adj-intf-name-out"`
							AdjDetailSetOut            bool     `json:"adj-detail-set-out" xml:"adj-detail-set-out"`
							AdjTransitionsOut          int      `json:"adj-transitions-out" xml:"adj-transitions-out"`
							AdjFlapOut                 bool     `json:"adj-flap-out" xml:"adj-flap-out"`
							AdjFlapTimeOut             Duration `json:"adj-flap-time-out" xml:"adj-flap-time-out"`
							AdjCktTypeOut              string   `json:"adj-ckt-type-out" xml:"adj-ckt-type-out"`
							AdjIpv4AddrOut             string   `json:"adj-ipv4-addr-out" xml:"adj-ipv4-addr-out"`
							AdjIpv6AddrOut             string   `json:"adj-ipv6-addr-out" xml:"adj-ipv6-addr-out"`
							AdjBcastOut                bool     `json:"adj-bcast-out" xml:"adj-bcast-out"`
							AdjBfdIpv4EstablishOut     bool     `json:"adj-bfd-ipv4-establish-out" xml:"adj-bfd-ipv4-establish-out"`
							AdjBfdIpv6EstablishOut     bool     `json:"adj-bfd-ipv6-establish-out" xml:"adj-bfd-ipv6-establish-out"`
							AdjResurrectOut            bool     `json:"adj-resurrect-out" xml:"adj-resurrect-out"`
							AdjRestartCapableOut       bool     `json:"adj-restart-capable-out" xml:"adj-restart-capable-out"`
							AdjRestartAckOut           bool     `json:"adj-restart-ack-out" xml:"adj-restart-ack-out"`
							AdjRestartModeOut          bool     `json:"adj-restart-mode-out" xml:"adj-restart-mode-out"`
							AdjRestartAdjSeenRaOut     bool     `json:"adj-restart-adj-seen-ra-out" xml:"adj-restart-adj-seen-ra-out"`
							AdjRestartAdjSeenCsnpOut   bool     `json:"adj-restart-adj-seen-csnp-out" xml:"adj-restart-adj-seen-csnp-out"`
							AdjRestartAdjSeenL1CsnpOut bool     `json:"adj-restart-adj-seen-l1-csnp-out" xml:"adj-restart-adj-seen-l1-csnp-out"`
							AdjRestartAdjSeenL2CsnpOut bool     `json:"adj-restart-adj-seen-l2-csnp-out" xml:"adj-restart-adj-seen-l2-csnp-out"`
							AdjRestartSuppressAdjOut   bool     `json:"adj-restart-suppress-adj-out" xml:"adj-restart-suppress-adj-out"`
							TableAdjSid                []struct {
								RowAdjSid []struct {
									AdjSidValue  int  `json:"adj-sid-value" xml:"adj-sid-value"`
									AdjSidFFlag  bool `json:"adj-sid-f-flag" xml:"adj-sid-f-flag"`
									AdjSidBFlag  bool `json:"adj-sid-b-flag" xml:"adj-sid-b-flag"`
									AdjSidVFlag  bool `json:"adj-sid-v-flag" xml:"adj-sid-v-flag"`
									AdjSidLFlag  bool `json:"adj-sid-l-flag" xml:"adj-sid-l-flag"`
									AdjSidSFlag  bool `json:"adj-sid-s-flag" xml:"adj-sid-s-flag"`
									AdjSidPFlag  bool `json:"adj-sid-p-flag" xml:"adj-sid-p-flag"`
									AdjSidWeight int  `json:"adj-sid-weight" xml:"adj-sid-weight"`
								} `json:"ROW_adj_sid" xml:"ROW_adj_sid"`
							} `json:"TABLE_adj_sid" xml:"TABLE_adj_sid"`
						} `json:"ROW_process_adj" xml:"ROW_process_adj"`
					} `json:"TABLE_process_adj" xml:"TABLE_process_adj"`
				} `json:"ROW_vrf" xml:"ROW_vrf"`
			} `json:"TABLE_vrf" xml:"TABLE_vrf"`
		} `json:"ROW_process_tag" xml:"ROW_process_tag"`
	} `json:"TABLE_process_tag" xml:"TABLE_process_tag"`
}

type ShowIsisAdjDetailResultFlat struct {
	//Rv
	VrfNameOut      string `json:"vrf-name-out" xml:"vrf-name-out"`
	AdjSummaryOut   bool   `json:"adj-summary-out" xml:"adj-summary-out"`
	AdjInterfaceOut bool   `json:"adj-interface-out" xml:"adj-interface-out"`
	//Rp
	AdjSysNameOut              string   `json:"adj-sys-name-out" xml:"adj-sys-name-out"`
	AdjSysIDOut                string   `json:"adj-sys-id-out" xml:"adj-sys-id-out"`
	AdjUsageOut                string   `json:"adj-usage-out" xml:"adj-usage-out"`
	AdjStateOut                string   `json:"adj-state-out" xml:"adj-state-out"`
	AdjHoldTimeOut             Duration `json:"adj-hold-time-out" xml:"adj-hold-time-out"`
	AdjIntfNameOut             string   `json:"adj-intf-name-out" xml:"adj-intf-name-out"`
	AdjDetailSetOut            bool     `json:"adj-detail-set-out" xml:"adj-detail-set-out"`
	AdjTransitionsOut          int      `json:"adj-transitions-out" xml:"adj-transitions-out"`
	AdjFlapOut                 bool     `json:"adj-flap-out" xml:"adj-flap-out"`
	AdjFlapTimeOut             Duration `json:"adj-flap-time-out" xml:"adj-flap-time-out"`
	AdjCktTypeOut              string   `json:"adj-ckt-type-out" xml:"adj-ckt-type-out"`
	AdjIpv4AddrOut             string   `json:"adj-ipv4-addr-out" xml:"adj-ipv4-addr-out"`
	AdjIpv6AddrOut             string   `json:"adj-ipv6-addr-out" xml:"adj-ipv6-addr-out"`
	AdjBcastOut                bool     `json:"adj-bcast-out" xml:"adj-bcast-out"`
	AdjBfdIpv4EstablishOut     bool     `json:"adj-bfd-ipv4-establish-out" xml:"adj-bfd-ipv4-establish-out"`
	AdjBfdIpv6EstablishOut     bool     `json:"adj-bfd-ipv6-establish-out" xml:"adj-bfd-ipv6-establish-out"`
	AdjResurrectOut            bool     `json:"adj-resurrect-out" xml:"adj-resurrect-out"`
	AdjRestartCapableOut       bool     `json:"adj-restart-capable-out" xml:"adj-restart-capable-out"`
	AdjRestartAckOut           bool     `json:"adj-restart-ack-out" xml:"adj-restart-ack-out"`
	AdjRestartModeOut          bool     `json:"adj-restart-mode-out" xml:"adj-restart-mode-out"`
	AdjRestartAdjSeenRaOut     bool     `json:"adj-restart-adj-seen-ra-out" xml:"adj-restart-adj-seen-ra-out"`
	AdjRestartAdjSeenCsnpOut   bool     `json:"adj-restart-adj-seen-csnp-out" xml:"adj-restart-adj-seen-csnp-out"`
	AdjRestartAdjSeenL1CsnpOut bool     `json:"adj-restart-adj-seen-l1-csnp-out" xml:"adj-restart-adj-seen-l1-csnp-out"`
	AdjRestartAdjSeenL2CsnpOut bool     `json:"adj-restart-adj-seen-l2-csnp-out" xml:"adj-restart-adj-seen-l2-csnp-out"`
	AdjRestartSuppressAdjOut   bool     `json:"adj-restart-suppress-adj-out" xml:"adj-restart-suppress-adj-out"`
}

func (d *ShowIsisAdjDetailResponse) Flat() (out []ShowIsisAdjDetailResultFlat) {
	return d.InsAPI.Outputs.Output.Flat()
}
func (d *ShowIsisAdjDetailResponseResult) Flat() (out []ShowIsisAdjDetailResultFlat) {
	for _, Tpt := range d.Body.TableProcessTag {
		for _, Rpt := range Tpt.RowProcessTag {
			for _, Tv := range Rpt.TableVrf {
				for _, Rv := range Tv.RowVrf {
					for _, Tp := range Rv.TableProcessAdj {
						for _, Rp := range Tp.RowProcessAdj {
							out = append(out, ShowIsisAdjDetailResultFlat{
								//Rv
								VrfNameOut:      Rv.VrfNameOut,
								AdjSummaryOut:   Rv.AdjSummaryOut,
								AdjInterfaceOut: Rv.AdjInterfaceOut,
								//Rp
								AdjSysNameOut:              Rp.AdjSysNameOut,
								AdjSysIDOut:                Rp.AdjSysIDOut,
								AdjUsageOut:                Rp.AdjUsageOut,
								AdjStateOut:                Rp.AdjStateOut,
								AdjHoldTimeOut:             Rp.AdjHoldTimeOut,
								AdjIntfNameOut:             Rp.AdjIntfNameOut,
								AdjDetailSetOut:            Rp.AdjDetailSetOut,
								AdjTransitionsOut:          Rp.AdjTransitionsOut,
								AdjFlapOut:                 Rp.AdjFlapOut,
								AdjFlapTimeOut:             Rp.AdjFlapTimeOut,
								AdjCktTypeOut:              Rp.AdjCktTypeOut,
								AdjIpv4AddrOut:             Rp.AdjIpv4AddrOut,
								AdjIpv6AddrOut:             Rp.AdjIpv6AddrOut,
								AdjBcastOut:                Rp.AdjBcastOut,
								AdjBfdIpv4EstablishOut:     Rp.AdjBfdIpv4EstablishOut,
								AdjBfdIpv6EstablishOut:     Rp.AdjBfdIpv6EstablishOut,
								AdjResurrectOut:            Rp.AdjResurrectOut,
								AdjRestartCapableOut:       Rp.AdjRestartCapableOut,
								AdjRestartAckOut:           Rp.AdjRestartAckOut,
								AdjRestartModeOut:          Rp.AdjRestartModeOut,
								AdjRestartAdjSeenRaOut:     Rp.AdjRestartAdjSeenRaOut,
								AdjRestartAdjSeenCsnpOut:   Rp.AdjRestartAdjSeenCsnpOut,
								AdjRestartAdjSeenL1CsnpOut: Rp.AdjRestartAdjSeenL1CsnpOut,
								AdjRestartAdjSeenL2CsnpOut: Rp.AdjRestartAdjSeenL2CsnpOut,
								AdjRestartSuppressAdjOut:   Rp.AdjRestartSuppressAdjOut,
							})
						}
					}
				}
			}
		}
	}
	return
}

// NewShowIsisAdjDetailFromString returns instance from an input string.
func NewShowIsisAdjDetailFromString(s string) (*ShowIsisAdjDetailResponse, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowIsisAdjDetailFromReader(strings.NewReader(s))
}

// NewShowIsisAdjDetailFromBytes returns instance from an input byte array.
func NewShowIsisAdjDetailFromBytes(s []byte) (*ShowIsisAdjDetailResponse, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowIsisAdjDetailFromReader(bytes.NewReader(s))
}

// NewShowIsisAdjDetailFromReader returns instance from an input reader.
func NewShowIsisAdjDetailFromReader(s io.Reader) (*ShowIsisAdjDetailResponse, error) {
	//si := &ShowIsisAdjDetail{}
	ShowIsisAdjDetailResponseDat := &ShowIsisAdjDetailResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowIsisAdjDetailResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowIsisAdjDetailResponseDat, nil
}

// NewShowIsisAdjDetailResultFromString returns instance from an input string.
func NewShowIsisAdjDetailResultFromString(s string) (*ShowIsisAdjDetailResponseResult, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowIsisAdjDetailResultFromReader(strings.NewReader(s))
}

// NewShowIsisAdjDetailResultFromBytes returns instance from an input byte array.
func NewShowIsisAdjDetailResultFromBytes(s []byte) (*ShowIsisAdjDetailResponseResult, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowIsisAdjDetailResultFromReader(bytes.NewReader(s))
}

// NewShowIsisAdjDetailResultFromReader returns instance from an input reader.
func NewShowIsisAdjDetailResultFromReader(s io.Reader) (*ShowIsisAdjDetailResponseResult, error) {
	//si := &ShowIsisAdjDetail{}
	ShowIsisAdjDetailDat := &ShowIsisAdjDetailResponseResult{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowIsisAdjDetailDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowIsisAdjDetailDat, nil
}
