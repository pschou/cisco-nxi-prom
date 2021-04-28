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

type InterfaceTransceiverDetailsResponse struct {
	InsAPI struct {
		Outputs struct {
			Output InterfaceTransceiverDetailsResponseResult `json:"output" xml:"output"`
		} `json:"outputs" xml:"outputs"`
		Sid     string `json:"sid" xml:"sid"`
		Type    string `json:"type" xml:"type"`
		Version string `json:"version" xml:"version"`
	} `json:"ins_api" xml:"ins_api"`
}

type InterfaceTransceiverDetailsResponseResult struct {
	Body  InterfaceTransceiverDetailsResultBody `json:"body" xml:"body"`
	Code  string                                `json:"code" xml:"code"`
	Input string                                `json:"input" xml:"input"`
	Msg   string                                `json:"msg" xml:"msg"`
}

type InterfaceTransceiverDetailsResultBody struct {
	TableInterface []struct {
		RowInterface []struct {
			Interface       string `json:"interface" xml:"interface"`
			Sfp             string `json:"sfp" xml:"sfp"`
			Type            string `json:"type,omitempty" xml:"type,omitempty"`
			Name            string `json:"name,omitempty" xml:"name,omitempty"`
			PartNum         string `json:"partnum,omitempty" xml:"partnum,omitempty"`
			Rev             string `json:"rev,omitempty" xml:"rev,omitempty"`
			SerialNum       string `json:"serialnum,omitempty" xml:"serialnum,omitempty"`
			NomBitrate      int    `json:"nom_bitrate,omitempty" xml:"nom_bitrate,omitempty"`
			Len50OM3        int    `json:"len_50_OM3,omitempty" xml:"len_50_OM3,omitempty"`
			CiscoID         string `json:"ciscoid,omitempty" xml:"ciscoid,omitempty"`
			CiscoID1        int    `json:"ciscoid_1,omitempty" xml:"ciscoid_1,omitempty"`
			CiscoPartNumber string `json:"cisco_part_number,omitempty" xml:"cisco_part_number,omitempty"`
			CiscoProductID  string `json:"cisco_product_id,omitempty" xml:"cisco_product_id,omitempty"`
			TABLELane       []struct {
				ROWLane []struct {
					LaneNumber    int      `json:"lane_number" xml:"lane_number"`
					Temperature   int      `json:"temperature" xml:"temperature"`
					TempAlrmHi    int      `json:"temp_alrm_hi" xml:"temp_alrm_hi"`
					TempAlrmLo    int      `json:"temp_alrm_lo" xml:"temp_alrm_lo"`
					TempWarnHi    int      `json:"temp_warn_hi" xml:"temp_warn_hi"`
					TempWarnLo    int      `json:"temp_warn_lo" xml:"temp_warn_lo"`
					Voltage       int      `json:"voltage" xml:"voltage"`
					VoltAlrmHi    int      `json:"volt_alrm_hi" xml:"volt_alrm_hi"`
					VoltAlrmLo    int      `json:"volt_alrm_lo" xml:"volt_alrm_lo"`
					VoltWarnHi    int      `json:"volt_warn_hi" xml:"volt_warn_hi"`
					VoltWarnLo    int      `json:"volt_warn_lo" xml:"volt_warn_lo"`
					Current       int      `json:"current" xml:"current"`
					CurrentAlrmHi int      `json:"current_alrm_hi" xml:"current_alrm_hi"`
					CurrentAlrmLo int      `json:"current_alrm_lo" xml:"current_alrm_lo"`
					CurrentWarnHi int      `json:"current_warn_hi" xml:"current_warn_hi"`
					CurrentWarnLo int      `json:"current_warn_lo" xml:"current_warn_lo"`
					TxPwr         int      `json:"tx_pwr" xml:"tx_pwr"`
					TxPwrAlrmHi   int      `json:"tx_pwr_alrm_hi" xml:"tx_pwr_alrm_hi"`
					TxPwrAlrmLo   int      `json:"tx_pwr_alrm_lo" xml:"tx_pwr_alrm_lo"`
					TxPwrWarnHi   int      `json:"tx_pwr_warn_hi" xml:"tx_pwr_warn_hi"`
					TxPwrWarnLo   int      `json:"tx_pwr_warn_lo" xml:"tx_pwr_warn_lo"`
					RxPwr         int      `json:"rx_pwr" xml:"rx_pwr"`
					RxPwrAlrmHi   int      `json:"rx_pwr_alrm_hi" xml:"rx_pwr_alrm_hi"`
					RxPwrAlrmLo   int      `json:"rx_pwr_alrm_lo" xml:"rx_pwr_alrm_lo"`
					RxPwrWarnHi   int      `json:"rx_pwr_warn_hi" xml:"rx_pwr_warn_hi"`
					RxPwrWarnLo   int      `json:"rx_pwr_warn_lo" xml:"rx_pwr_warn_lo"`
					XmitFaults    int      `json:"xmit_faults" xml:"xmit_faults"`
					RxPwrFlag     []string `json:"rx_pwr_flag" xml:"rx_pwr_flag"`
					TxPwrFlag     []string `json:"tx_pwr_flag" xml:"tx_pwr_flag"`
					TempFlag      []string `json:"temp_flag" xml:"temp_flag"`
					VoltFlag      []string `json:"volt_flag" xml:"volt_flag"`
					CurrentFlag   []string `json:"current_flag" xml:"current_flag"`
				} `json:"ROW_lane" xml:"ROW_lane"`
			} `json:"TABLE_lane,omitempty" xml:"TABLE_lane,omitempty"`
			CiscoVendorID string `json:"cisco_vendor_id,omitempty" xml:"cisco_vendor_id,omitempty"`
		} `json:"ROW_interface" xml:"ROW_interface"`
	} `json:"TABLE_interface" xml:"TABLE_interface"`
}

type InterfaceTransceiverDetailsResultFlat struct {
	Interface       string   `json:"interface" xml:"interface"`
	Sfp             string   `json:"sfp" xml:"sfp"`
	Type            string   `json:"type,omitempty" xml:"type,omitempty"`
	Name            string   `json:"name,omitempty" xml:"name,omitempty"`
	PartNum         string   `json:"partnum,omitempty" xml:"partnum,omitempty"`
	Rev             string   `json:"rev,omitempty" xml:"rev,omitempty"`
	SerialNum       string   `json:"serialnum,omitempty" xml:"serialnum,omitempty"`
	NomBitrate      int      `json:"nom_bitrate,omitempty" xml:"nom_bitrate,omitempty"`
	Len50OM3        int      `json:"len_50_OM3,omitempty" xml:"len_50_OM3,omitempty"`
	CiscoID         string   `json:"ciscoid,omitempty" xml:"ciscoid,omitempty"`
	CiscoID1        int      `json:"ciscoid_1,omitempty" xml:"ciscoid_1,omitempty"`
	CiscoPartNumber string   `json:"cisco_part_number,omitempty" xml:"cisco_part_number,omitempty"`
	CiscoProductID  string   `json:"cisco_product_id,omitempty" xml:"cisco_product_id,omitempty"`
	LaneNumber      int      `json:"lane_number" xml:"lane_number"`
	Temperature     int      `json:"temperature" xml:"temperature"`
	TempAlrmHi      int      `json:"temp_alrm_hi" xml:"temp_alrm_hi"`
	TempAlrmLo      int      `json:"temp_alrm_lo" xml:"temp_alrm_lo"`
	TempWarnHi      int      `json:"temp_warn_hi" xml:"temp_warn_hi"`
	TempWarnLo      int      `json:"temp_warn_lo" xml:"temp_warn_lo"`
	Voltage         int      `json:"voltage" xml:"voltage"`
	VoltAlrmHi      int      `json:"volt_alrm_hi" xml:"volt_alrm_hi"`
	VoltAlrmLo      int      `json:"volt_alrm_lo" xml:"volt_alrm_lo"`
	VoltWarnHi      int      `json:"volt_warn_hi" xml:"volt_warn_hi"`
	VoltWarnLo      int      `json:"volt_warn_lo" xml:"volt_warn_lo"`
	Current         int      `json:"current" xml:"current"`
	CurrentAlrmHi   int      `json:"current_alrm_hi" xml:"current_alrm_hi"`
	CurrentAlrmLo   int      `json:"current_alrm_lo" xml:"current_alrm_lo"`
	CurrentWarnHi   int      `json:"current_warn_hi" xml:"current_warn_hi"`
	CurrentWarnLo   int      `json:"current_warn_lo" xml:"current_warn_lo"`
	TxPwr           int      `json:"tx_pwr" xml:"tx_pwr"`
	TxPwrAlrmHi     int      `json:"tx_pwr_alrm_hi" xml:"tx_pwr_alrm_hi"`
	TxPwrAlrmLo     int      `json:"tx_pwr_alrm_lo" xml:"tx_pwr_alrm_lo"`
	TxPwrWarnHi     int      `json:"tx_pwr_warn_hi" xml:"tx_pwr_warn_hi"`
	TxPwrWarnLo     int      `json:"tx_pwr_warn_lo" xml:"tx_pwr_warn_lo"`
	RxPwr           int      `json:"rx_pwr" xml:"rx_pwr"`
	RxPwrAlrmHi     int      `json:"rx_pwr_alrm_hi" xml:"rx_pwr_alrm_hi"`
	RxPwrAlrmLo     int      `json:"rx_pwr_alrm_lo" xml:"rx_pwr_alrm_lo"`
	RxPwrWarnHi     int      `json:"rx_pwr_warn_hi" xml:"rx_pwr_warn_hi"`
	RxPwrWarnLo     int      `json:"rx_pwr_warn_lo" xml:"rx_pwr_warn_lo"`
	XmitFaults      int      `json:"xmit_faults" xml:"xmit_faults"`
	RxPwrFlag       []string `json:"rx_pwr_flag" xml:"rx_pwr_flag"`
	TxPwrFlag       []string `json:"tx_pwr_flag" xml:"tx_pwr_flag"`
	TempFlag        []string `json:"temp_flag" xml:"temp_flag"`
	VoltFlag        []string `json:"volt_flag" xml:"volt_flag"`
	CurrentFlag     []string `json:"current_flag" xml:"current_flag"`
}

func (d *InterfaceTransceiverDetailsResponse) Flat() (out []InterfaceTransceiverDetailsResultFlat) {
	return d.InsAPI.Outputs.Output.Flat()
}
func (d *InterfaceTransceiverDetailsResponseResult) Flat() (out []InterfaceTransceiverDetailsResultFlat) {
	for _, Ti := range d.Body.TableInterface {
		for _, Ri := range Ti.RowInterface {
			if len(Ri.TABLELane) == 0 {
				out = append(out, InterfaceTransceiverDetailsResultFlat{
					Interface:       Ri.Interface,
					Sfp:             Ri.Sfp,
					Type:            Ri.Type,
					Name:            Ri.Name,
					PartNum:         Ri.PartNum,
					Rev:             Ri.Rev,
					SerialNum:       Ri.SerialNum,
					NomBitrate:      Ri.NomBitrate,
					Len50OM3:        Ri.Len50OM3,
					CiscoID:         Ri.CiscoID,
					CiscoID1:        Ri.CiscoID1,
					CiscoPartNumber: Ri.CiscoPartNumber,
					CiscoProductID:  Ri.CiscoProductID,
				})
			}
			for _, TL := range Ri.TABLELane {
				for _, RL := range TL.ROWLane {
					out = append(out, InterfaceTransceiverDetailsResultFlat{
						Interface:       Ri.Interface,
						Sfp:             Ri.Sfp,
						Type:            Ri.Type,
						Name:            Ri.Name,
						PartNum:         Ri.PartNum,
						Rev:             Ri.Rev,
						SerialNum:       Ri.SerialNum,
						NomBitrate:      Ri.NomBitrate,
						Len50OM3:        Ri.Len50OM3,
						CiscoID:         Ri.CiscoID,
						CiscoID1:        Ri.CiscoID1,
						CiscoPartNumber: Ri.CiscoPartNumber,
						CiscoProductID:  Ri.CiscoProductID,
						LaneNumber:      RL.LaneNumber,
						Temperature:     RL.Temperature,
						TempAlrmHi:      RL.TempAlrmHi,
						TempAlrmLo:      RL.TempAlrmLo,
						TempWarnHi:      RL.TempWarnHi,
						TempWarnLo:      RL.TempWarnLo,
						Voltage:         RL.Voltage,
						VoltAlrmHi:      RL.VoltAlrmHi,
						VoltAlrmLo:      RL.VoltAlrmLo,
						VoltWarnHi:      RL.VoltWarnHi,
						VoltWarnLo:      RL.VoltWarnLo,
						Current:         RL.Current,
						CurrentAlrmHi:   RL.CurrentAlrmHi,
						CurrentAlrmLo:   RL.CurrentAlrmLo,
						CurrentWarnHi:   RL.CurrentWarnHi,
						CurrentWarnLo:   RL.CurrentWarnLo,
						TxPwr:           RL.TxPwr,
						TxPwrAlrmHi:     RL.TxPwrAlrmHi,
						TxPwrAlrmLo:     RL.TxPwrAlrmLo,
						TxPwrWarnHi:     RL.TxPwrWarnHi,
						TxPwrWarnLo:     RL.TxPwrWarnLo,
						RxPwr:           RL.RxPwr,
						RxPwrAlrmHi:     RL.RxPwrAlrmHi,
						RxPwrAlrmLo:     RL.RxPwrAlrmLo,
						RxPwrWarnHi:     RL.RxPwrWarnHi,
						RxPwrWarnLo:     RL.RxPwrWarnLo,
						XmitFaults:      RL.XmitFaults,
						RxPwrFlag:       RL.RxPwrFlag,
						TxPwrFlag:       RL.TxPwrFlag,
						TempFlag:        RL.TempFlag,
						VoltFlag:        RL.VoltFlag,
						CurrentFlag:     RL.CurrentFlag,
					})
				}
			}
		}
	}
	return
}

// NewInterfaceTransceiverDetailsFromString returns instance from an input string.
func NewInterfaceTransceiverDetailsFromString(s string) (*InterfaceTransceiverDetailsResponse, error) {
	return NewInterfaceTransceiverDetailsFromReader(strings.NewReader(s))
}

// NewInterfaceTransceiverDetailsFromBytes returns instance from an input byte array.
func NewInterfaceTransceiverDetailsFromBytes(s []byte) (*InterfaceTransceiverDetailsResponse, error) {
	return NewInterfaceTransceiverDetailsFromReader(bytes.NewReader(s))
}

// NewInterfaceTransceiverDetailsFromReader returns instance from an input reader.
func NewInterfaceTransceiverDetailsFromReader(s io.Reader) (*InterfaceTransceiverDetailsResponse, error) {
	//si := &InterfaceTransceiverDetails{}
	InterfaceTransceiverDetailsResponseDat := &InterfaceTransceiverDetailsResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(InterfaceTransceiverDetailsResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return InterfaceTransceiverDetailsResponseDat, nil
}

// NewInterfaceTransceiverDetailsResultFromString returns instance from an input string.
func NewInterfaceTransceiverDetailsResultFromString(s string) (*InterfaceTransceiverDetailsResponseResult, error) {
	return NewInterfaceTransceiverDetailsResultFromReader(strings.NewReader(s))
}

// NewInterfaceTransceiverDetailsResultFromBytes returns instance from an input byte array.
func NewInterfaceTransceiverDetailsResultFromBytes(s []byte) (*InterfaceTransceiverDetailsResponseResult, error) {
	return NewInterfaceTransceiverDetailsResultFromReader(bytes.NewReader(s))
}

// NewInterfaceTransceiverDetailsResultFromReader returns instance from an input reader.
func NewInterfaceTransceiverDetailsResultFromReader(s io.Reader) (*InterfaceTransceiverDetailsResponseResult, error) {
	//si := &InterfaceTransceiverDetailsResponseResult{}
	InterfaceTransceiverDetailsResponseResultDat := &InterfaceTransceiverDetailsResponseResult{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(InterfaceTransceiverDetailsResponseResultDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return InterfaceTransceiverDetailsResponseResultDat, nil
}
