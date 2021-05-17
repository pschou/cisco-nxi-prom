// Copyright 2018 Paul Greenberg (greenpau@outlook.com)
//            and Paul Schou     (github.com/pschou)
//
// Licensed under the Apache License, ShowVersion 2.0 (the "License");
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

type ShowVersionResponse struct {
	InsAPI struct {
		Outputs struct {
			Output ShowVersionResponseResult `json:"output" xml:"output"`
		} `json:"outputs" xml:"outputs"`
		Sid         string `json:"sid" xml:"sid"`
		Type        string `json:"type" xml:"type"`
		ShowVersion string `json:"version" xml:"version"`
	} `json:"ins_api" xml:"ins_api"`
}

type ShowVersionResponseResult struct {
	Body  ShowVersionResultBody `json:"body" xml:"body"`
	Code  string                `json:"code" xml:"code"`
	Input string                `json:"input" xml:"input"`
	Msg   string                `json:"msg" xml:"msg"`
}

type ShowVersionResultBody struct {
	HeaderStr        string    `json:"header_str" xml:"header_str"`
	BiosVerStr       string    `json:"bios_ver_str" xml:"bios_ver_str"`
	KickstartVerStr  string    `json:"kickstart_ver_str" xml:"kickstart_ver_str"`
	NxosVerStr       string    `json:"nxos_ver_str" xml:"nxos_ver_str"`
	BiosCmplTime     TimeStamp `json:"bios_cmpl_time" xml:"bios_cmpl_time"`
	KickFileName     string    `json:"kick_file_name" xml:"kick_file_name"`
	NxosFileName     string    `json:"nxos_file_name" xml:"nxos_file_name"`
	KickCmplTime     TimeStamp `json:"kick_cmpl_time" xml:"kick_cmpl_time"`
	NxosCmplTime     TimeStamp `json:"nxos_cmpl_time" xml:"nxos_cmpl_time"`
	KickTmstmp       TimeStamp `json:"kick_tmstmp" xml:"kick_tmstmp"`
	NxosTmstmp       TimeStamp `json:"nxos_tmstmp" xml:"nxos_tmstmp"`
	ChassisID        string    `json:"chassis_id" xml:"chassis_id"`
	CPUName          string    `json:"cpu_name" xml:"cpu_name"`
	Memory           uint      `json:"memory" xml:"memory"`
	MemType          string    `json:"mem_type" xml:"mem_type"`
	ProcBoardID      string    `json:"proc_board_id" xml:"proc_board_id"`
	HostName         string    `json:"host_name" xml:"host_name"`
	BootflashSize    uint      `json:"bootflash_size" xml:"bootflash_size"`
	KernUptmDays     uint      `json:"kern_uptm_days" xml:"kern_uptm_days"`
	KernUptmHrs      uint      `json:"kern_uptm_hrs" xml:"kern_uptm_hrs"`
	KernUptmMins     uint      `json:"kern_uptm_mins" xml:"kern_uptm_mins"`
	KernUptmSecs     uint      `json:"kern_uptm_secs" xml:"kern_uptm_secs"`
	RrUsecs          uint      `json:"rr_usecs" xml:"rr_usecs"`
	RrCtime          TimeStamp `json:"rr_ctime" xml:"rr_ctime"`
	RrReason         string    `json:"rr_reason" xml:"rr_reason"`
	RrSysVer         string    `json:"rr_sys_ver" xml:"rr_sys_ver"`
	RrService        string    `json:"rr_service" xml:"rr_service"`
	Plugins          string    `json:"plugins" xml:"plugins"`
	Manufacturer     string    `json:"manufacturer" xml:"manufacturer"`
	TablePackageList []struct {
		RowPackageList []struct {
			PackageID string `json:"package_id" xml:"package_id"`
		} `json:"ROW_package_list" xml:"ROW_package_list"`
	} `json:"TABLE_package_list" xml:"TABLE_package_list"`
}

// NewShowVersionFromString returns instance from an input string.
func NewShowVersionFromString(s string) (*ShowVersionResponse, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowVersionFromReader(strings.NewReader(s))
}

// NewShowVersionFromBytes returns instance from an input byte array.
func NewShowVersionFromBytes(s []byte) (*ShowVersionResponse, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowVersionFromReader(bytes.NewReader(s))
}

// NewShowVersionFromReader returns instance from an input reader.
func NewShowVersionFromReader(s io.Reader) (*ShowVersionResponse, error) {
	//si := &ShowVersion{}
	ShowVersionResponseDat := &ShowVersionResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	jsonDec.IgnoreEmptyObject()
	err := jsonDec.Decode(ShowVersionResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowVersionResponseDat, nil
}

// NewShowVersionResultFromString returns instance from an input string.
func NewShowVersionResultFromString(s string) (*ShowVersionResponseResult, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowVersionResultFromReader(strings.NewReader(s))
}

// NewShowVersionResultFromBytes returns instance from an input byte array.
func NewShowVersionResultFromBytes(s []byte) (*ShowVersionResponseResult, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowVersionResultFromReader(bytes.NewReader(s))
}

// NewShowVersionResultFromReader returns instance from an input reader.
func NewShowVersionResultFromReader(s io.Reader) (*ShowVersionResponseResult, error) {
	//si := &ShowVersionResponseResult{}
	ShowVersionResponseResultDat := &ShowVersionResponseResult{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	jsonDec.IgnoreEmptyObject()
	err := jsonDec.Decode(ShowVersionResponseResultDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowVersionResponseResultDat, nil
}
