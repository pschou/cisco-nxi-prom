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

type ShowHsrpResponse struct {
	InsAPI struct {
		Outputs struct {
			Output ShowHsrpResponseResult `json:"output" xml:"output"`
		} `json:"outputs" xml:"outputs"`
		Sid     string `json:"sid" xml:"sid"`
		Type    string `json:"type" xml:"type"`
		Version string `json:"version" xml:"version"`
	} `json:"ins_api" xml:"ins_api"`
}

type ShowHsrpResponseResult struct {
	Body  ShowHsrpResultBody `json:"body" xml:"body"`
	Code  string             `json:"code" xml:"code"`
	Input string             `json:"input" xml:"input"`
	Msg   string             `json:"msg" xml:"msg"`
}

type ShowHsrpResultBody struct {
	TableGrpDetail []struct {
		RowGrpDetail []struct {
			ShIfIndex                string `json:"sh_if_index" xml:"sh_if_index"`
			ShGroupNum               int    `json:"sh_group_num" xml:"sh_group_num"`
			ShGroupType              string `json:"sh_group_type" xml:"sh_group_type"`
			ShGroupVersion           string `json:"sh_group_version" xml:"sh_group_version"`
			ShGroupState             string `json:"sh_group_state" xml:"sh_group_state"`
			ShPrio                   int    `json:"sh_prio" xml:"sh_prio"`
			ShCfgPrio                int    `json:"sh_cfg_prio" xml:"sh_cfg_prio"`
			ShFwdLowerThreshold      int    `json:"sh_fwd_lower_threshold" xml:"sh_fwd_lower_threshold"`
			ShFwdUpperThreshold      int    `json:"sh_fwd_upper_threshold" xml:"sh_fwd_upper_threshold"`
			ShCanForward             string `json:"sh_can_forward" xml:"sh_can_forward"`
			ShPreempt                string `json:"sh_preempt" xml:"sh_preempt"`
			ShCurHello               int    `json:"sh_cur_hello" xml:"sh_cur_hello"`
			ShCurHelloAttr           string `json:"sh_cur_hello_attr" xml:"sh_cur_hello_attr"`
			ShCfgHello               int    `json:"sh_cfg_hello" xml:"sh_cfg_hello"`
			ShCfgHelloAttr           string `json:"sh_cfg_hello_attr" xml:"sh_cfg_hello_attr"`
			ShActiveHello            string `json:"sh_active_hello" xml:"sh_active_hello"`
			ShCurHold                int    `json:"sh_cur_hold" xml:"sh_cur_hold"`
			ShCurHoldAttr            string `json:"sh_cur_hold_attr" xml:"sh_cur_hold_attr"`
			ShCfgHold                int    `json:"sh_cfg_hold" xml:"sh_cfg_hold"`
			ShCfgHoldAttr            string `json:"sh_cfg_hold_attr" xml:"sh_cfg_hold_attr"`
			ShVip                    string `json:"sh_vip,omitempty" xml:"sh_vip,omitempty"`
			ShVipAttr                string `json:"sh_vip_attr" xml:"sh_vip_attr"`
			ShActiveRouterAddr       string `json:"sh_active_router_addr,omitempty" xml:"sh_active_router_addr,omitempty"`
			ShActiveRouterPrio       int    `json:"sh_active_router_prio" xml:"sh_active_router_prio"`
			ShActiveRouterTimer      string `json:"sh_active_router_timer" xml:"sh_active_router_timer"`
			ShStandbyRouterAddr      string `json:"sh_standby_router_addr,omitempty" xml:"sh_standby_router_addr,omitempty"`
			ShStandbyRouterPrio      int    `json:"sh_standby_router_prio" xml:"sh_standby_router_prio"`
			ShAuthenticationType     string `json:"sh_authentication_type" xml:"sh_authentication_type"`
			ShAuthenticationData     string `json:"sh_authentication_data" xml:"sh_authentication_data"`
			ShVmac                   string `json:"sh_vmac" xml:"sh_vmac"`
			ShVmacAttr               string `json:"sh_vmac_attr" xml:"sh_vmac_attr"`
			ShNumOfStateChanges      int    `json:"sh_num_of_state_changes" xml:"sh_num_of_state_changes"`
			ShLastStateChange        int    `json:"sh_last_state_change" xml:"sh_last_state_change"`
			ShNumOfTotalStateChanges int    `json:"sh_num_of_total_state_changes" xml:"sh_num_of_total_state_changes"`
			ShLastTotalStateChange   int    `json:"sh_last_total_state_change" xml:"sh_last_total_state_change"`
			ShNumTrackObj            int    `json:"sh_num_track_obj" xml:"sh_num_track_obj"`
			ShIPRedundName           string `json:"sh_ip_redund_name" xml:"sh_ip_redund_name"`
			ShIPRedundNameAttr       string `json:"sh_ip_redund_name_attr" xml:"sh_ip_redund_name_attr"`
			ShVipV6                  string `json:"sh_vip_v6,omitempty" xml:"sh_vip_v6,omitempty"`
			TableGrpVipSec           []struct {
				RowGrpVipSec []struct {
					ShVipSec string `json:"sh_vip_sec" xml:"sh_vip_sec"`
				} `json:"ROW_grp_vip_sec" xml:"ROW_grp_vip_sec"`
			} `json:"TABLE_grp_vip_sec,omitempty" xml:"TABLE_grp_vip_sec,omitempty"`
			ShActiveRouterAddrV6  string `json:"sh_active_router_addr_v6,omitempty" xml:"sh_active_router_addr_v6,omitempty"`
			ShStandbyRouterAddrV6 string `json:"sh_standby_router_addr_v6,omitempty" xml:"sh_standby_router_addr_v6,omitempty"`
		} `json:"ROW_grp_detail" xml:"ROW_grp_detail"`
	} `json:"TABLE_grp_detail" xml:"TABLE_grp_detail"`
}

// NewShowHsrpFromString returns instance from an input string.
func NewShowHsrpFromString(s string) (*ShowHsrpResponse, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowHsrpFromReader(strings.NewReader(s))
}

// NewShowHsrpFromBytes returns instance from an input byte array.
func NewShowHsrpFromBytes(s []byte) (*ShowHsrpResponse, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowHsrpFromReader(bytes.NewReader(s))
}

// NewShowHsrpFromReader returns instance from an input reader.
func NewShowHsrpFromReader(s io.Reader) (*ShowHsrpResponse, error) {
	//si := &ShowHsrp{}
	ShowHsrpResponseDat := &ShowHsrpResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	jsonDec.IgnoreEmptyObject()
	err := jsonDec.Decode(ShowHsrpResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowHsrpResponseDat, nil
}

// NewShowHsrpResultFromString returns instance from an input string.
func NewShowHsrpResultFromString(s string) (*ShowHsrpResponseResult, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowHsrpResultFromReader(strings.NewReader(s))
}

// NewShowHsrpResultFromBytes returns instance from an input byte array.
func NewShowHsrpResultFromBytes(s []byte) (*ShowHsrpResponseResult, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowHsrpResultFromReader(bytes.NewReader(s))
}

// NewShowHsrpResultFromReader returns instance from an input reader.
func NewShowHsrpResultFromReader(s io.Reader) (*ShowHsrpResponseResult, error) {
	//si := &ShowHsrpResponseResult{}
	ShowHsrpResponseResultDat := &ShowHsrpResponseResult{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowHsrpResponseResultDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowHsrpResponseResultDat, nil
}
