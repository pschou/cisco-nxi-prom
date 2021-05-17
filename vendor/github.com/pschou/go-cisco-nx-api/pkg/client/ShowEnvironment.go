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

type ShowEnvironmentResponse struct {
	InsAPI struct {
		Outputs struct {
			Output ShowEnvironmentResponseResult `json:"output" xml:"output"`
		} `json:"outputs" xml:"outputs"`
		Sid     string `json:"sid" xml:"sid"`
		Type    string `json:"type" xml:"type"`
		Version string `json:"version" xml:"version"`
	} `json:"ins_api" xml:"ins_api"`
}

type ShowEnvironmentResponseResult struct {
	Body  ShowEnvironmentResultBody `json:"body" xml:"body"`
	Code  string                    `json:"code" xml:"code"`
	Input string                    `json:"input" xml:"input"`
	Msg   string                    `json:"msg" xml:"msg"`
}

type ShowEnvironmentResultBody struct {
	TableTempinfo []struct {
		RowTempinfo []struct {
			AlarmStatus string  `json:"alarmstatus" xml:"alarmstatus"`
			CurTemp     float32 `json:"curtemp" xml:"curtemp"`
			MajThres    float32 `json:"majthres" xml:"majthres"`
			MinThres    float32 `json:"minthres" xml:"minthres"`
			Sensor      string  `json:"sensor" xml:"sensor"`
			Tempmod     int     `json:"tempmod" xml:"tempmod"`
		} `json:"ROW_tempinfo" xml:"ROW_tempinfo"`
	} `json:"TABLE_tempinfo" xml:"TABLE_tempinfo"`
	FanDetails struct {
		TableFanZoneSpeed []struct {
			RowFanZoneSpeed []struct {
				Zone      int    `json:"zone" xml:"zone"`
				ZoneSpeed string `json:"zonespeed" xml:"zonespeed"`
			} `json:"ROW_fan_zone_speed" xml:"ROW_fan_zone_speed"`
		} `json:"TABLE_fan_zone_speed" xml:"TABLE_fan_zone_speed"`
		TableFaninfo []struct {
			RowFaninfo []struct {
				FanDir    string `json:"fandir" xml:"fandir"`
				FanHwVer  string `json:"fanhwver" xml:"fanhwver"`
				FanModel  string `json:"fanmodel" xml:"fanmodel"`
				FanName   string `json:"fanname" xml:"fanname"`
				FanStatus string `json:"fanstatus" xml:"fanstatus"`
			} `json:"ROW_faninfo" xml:"ROW_faninfo"`
		} `json:"TABLE_faninfo" xml:"TABLE_faninfo"`
		FanFilterStatus string `json:"fan_filter_status" xml:"fan_filter_status"`
	} `json:"fandetails" xml:"fandetails"`
	Powersup struct {
		TableModPowInfo []struct {
			RowModPowInfo []struct {
				ActualDraw Watts  `json:"actual_draw" xml:"actual_draw"`
				Allocated  Watts  `json:"allocated" xml:"allocated"`
				ModModel   string `json:"mod_model" xml:"mod_model"`
				Modnum     string `json:"modnum" xml:"modnum"`
				Modstatus  string `json:"modstatus" xml:"modstatus"`
			} `json:"ROW_mod_pow_info" xml:"ROW_mod_pow_info"`
		} `json:"TABLE_mod_pow_info" xml:"TABLE_mod_pow_info"`
		TablePsinfo []struct {
			RowPsinfo []struct {
				ActualInput Watts  `json:"actual_input" xml:"actual_input"`
				ActualOut   Watts  `json:"actual_out" xml:"actual_out"`
				PsStatus    string `json:"ps_status" xml:"ps_status"`
				PsModel     string `json:"psmodel" xml:"psmodel"`
				PsNum       int    `json:"psnum" xml:"psnum"`
				TotCapa     Watts  `json:"tot_capa" xml:"tot_capa"`
			} `json:"ROW_psinfo" xml:"ROW_psinfo"`
		} `json:"TABLE_psinfo" xml:"TABLE_psinfo"`
		PowerSummary struct {
			AvailablePow          Watts  `json:"available_pow" xml:"available_pow"`
			CumulativePower       Watts  `json:"cumulative_power" xml:"cumulative_power"`
			PsOperMode            string `json:"ps_oper_mode" xml:"ps_oper_mode"`
			PsRedunMode           string `json:"ps_redun_mode" xml:"ps_redun_mode"`
			TotGridaCapacity      Watts  `json:"tot_gridA_capacity" xml:"tot_gridA_capacity"`
			TotGridbCapacity      Watts  `json:"tot_gridB_capacity" xml:"tot_gridB_capacity"`
			TotPowAllocBudgeted   Watts  `json:"tot_pow_alloc_budgeted" xml:"tot_pow_alloc_budgeted"`
			TotPowCapacity        Watts  `json:"tot_pow_capacity" xml:"tot_pow_capacity"`
			TotPowInputActualDraw Watts  `json:"tot_pow_input_actual_draw" xml:"tot_pow_input_actual_draw"`
			TotPowOutActualDraw   Watts  `json:"tot_pow_out_actual_draw" xml:"tot_pow_out_actual_draw"`
		} `json:"power_summary" xml:"power_summary"`
		VoltageLevel int `json:"voltage_level" xml:"voltage_level"`
	} `json:"powersup" xml:"powersup"`
}

// NewShowEnvironmentFromString returns instance from an input string.
func NewShowEnvironmentFromString(s string) (*ShowEnvironmentResponse, error) { if len(s) == 0 { return nil, fmt.Errorf("missing result") }
	return NewShowEnvironmentFromReader(strings.NewReader(s))
}

// NewShowEnvironmentFromBytes returns instance from an input byte array.
func NewShowEnvironmentFromBytes(s []byte) (*ShowEnvironmentResponse, error) { if len(s) == 0 { return nil, fmt.Errorf("missing result") }
	return NewShowEnvironmentFromReader(bytes.NewReader(s))
}

// NewShowEnvironmentFromReader returns instance from an input reader.
func NewShowEnvironmentFromReader(s io.Reader) (*ShowEnvironmentResponse, error) {
	//si := &ShowEnvironment{}
	ShowEnvironmentResponseDat := &ShowEnvironmentResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvertWithTrimSpace()
	jsonDec.UseAutoTrimSpace()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowEnvironmentResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowEnvironmentResponseDat, nil
}

// NewShowEnvironmentResultFromString returns instance from an input string.
func NewShowEnvironmentResultFromString(s string) (*ShowEnvironmentResponseResult, error) { if len(s) == 0 { return nil, fmt.Errorf("missing result") }
	return NewShowEnvironmentResultFromReader(strings.NewReader(s))
}

// NewShowEnvironmentResultFromBytes returns instance from an input byte array.
func NewShowEnvironmentResultFromBytes(s []byte) (*ShowEnvironmentResponseResult, error) { if len(s) == 0 { return nil, fmt.Errorf("missing result") }
	return NewShowEnvironmentResultFromReader(bytes.NewReader(s))
}

// NewShowEnvironmentResultFromReader returns instance from an input reader.
func NewShowEnvironmentResultFromReader(s io.Reader) (*ShowEnvironmentResponseResult, error) {
	//si := &ShowEnvironmentResponseResult{}
	ShowEnvironmentResponseResultDat := &ShowEnvironmentResponseResult{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvertWithTrimSpace()
	jsonDec.UseAutoTrimSpace()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowEnvironmentResponseResultDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowEnvironmentResponseResultDat, nil
}
