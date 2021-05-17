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

type ShowModuleResponse struct {
	InsAPI struct {
		Outputs struct {
			Output ShowModuleResponseResult `json:"output" xml:"output"`
		} `json:"outputs" xml:"outputs"`
		Sid     string `json:"sid" xml:"sid"`
		Type    string `json:"type" xml:"type"`
		Version string `json:"version" xml:"version"`
	} `json:"ins_api" xml:"ins_api"`
}

type ShowModuleResponseResult struct {
	Body  ShowModuleResultBody `json:"body" xml:"body"`
	Code  string               `json:"code" xml:"code"`
	Input string               `json:"input" xml:"input"`
	Msg   string               `json:"msg" xml:"msg"`
}

type ShowModuleResultBody struct {
	TableModdiaginfo []struct {
		RowModdiaginfo []struct {
			Diagstatus string `json:"diagstatus" xml:"diagstatus"`
			Mod        int    `json:"mod" xml:"mod"`
		} `json:"ROW_moddiaginfo" xml:"ROW_moddiaginfo"`
	} `json:"TABLE_moddiaginfo" xml:"TABLE_moddiaginfo"`
	TableModinfo []struct {
		RowModinfo []struct {
			Model   string `json:"model" xml:"model"`
			Modinf  int    `json:"modinf" xml:"modinf"`
			Modtype string `json:"modtype" xml:"modtype"`
			Ports   int    `json:"ports" xml:"ports"`
			Status  string `json:"status" xml:"status"`
		} `json:"ROW_modinfo" xml:"ROW_modinfo"`
	} `json:"TABLE_modinfo" xml:"TABLE_modinfo"`
	TableModmacinfo []struct {
		RowModmacinfo []struct {
			Mac       string `json:"mac" xml:"mac"`
			Modmac    int    `json:"modmac" xml:"modmac"`
			Serialnum string `json:"serialnum" xml:"serialnum"`
		} `json:"ROW_modmacinfo" xml:"ROW_modmacinfo"`
	} `json:"TABLE_modmacinfo" xml:"TABLE_modmacinfo"`
	TableModpwrinfo []struct {
		RowModpwrinfo []struct {
			Modpwr  int    `json:"modpwr" xml:"modpwr"`
			Pwrstat string `json:"pwrstat" xml:"pwrstat"`
			Reason  string `json:"reason" xml:"reason"`
		} `json:"ROW_modpwrinfo" xml:"ROW_modpwrinfo"`
	} `json:"TABLE_modpwrinfo" xml:"TABLE_modpwrinfo"`
	TableModwwninfo []struct {
		RowModwwninfo []struct {
			Hw       string `json:"hw" xml:"hw"`
			Modwwn   int    `json:"modwwn" xml:"modwwn"`
			Slottype string `json:"slottype" xml:"slottype"`
			Sw       string `json:"sw" xml:"sw"`
		} `json:"ROW_modwwninfo" xml:"ROW_modwwninfo"`
	} `json:"TABLE_modwwninfo" xml:"TABLE_modwwninfo"`
}

// NewShowModuleFromString returns instance from an input string.
func NewShowModuleFromString(s string) (*ShowModuleResponse, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowModuleFromReader(strings.NewReader(s))
}

// NewShowModuleFromBytes returns instance from an input byte array.
func NewShowModuleFromBytes(s []byte) (*ShowModuleResponse, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowModuleFromReader(bytes.NewReader(s))
}

// NewShowModuleFromReader returns instance from an input reader.
func NewShowModuleFromReader(s io.Reader) (*ShowModuleResponse, error) {
	//si := &ShowModule{}
	ShowModuleResponseDat := &ShowModuleResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowModuleResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowModuleResponseDat, nil
}

// NewShowModuleResultFromString returns instance from an input string.
func NewShowModuleResultFromString(s string) (*ShowModuleResponseResult, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowModuleResultFromReader(strings.NewReader(s))
}

// NewShowModuleResultFromBytes returns instance from an input byte array.
func NewShowModuleResultFromBytes(s []byte) (*ShowModuleResponseResult, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowModuleResultFromReader(bytes.NewReader(s))
}

// NewShowModuleResultFromReader returns instance from an input reader.
func NewShowModuleResultFromReader(s io.Reader) (*ShowModuleResponseResult, error) {
	//si := &ShowModuleResponseResult{}
	ShowModuleResponseResultDat := &ShowModuleResponseResult{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowModuleResponseResultDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowModuleResponseResultDat, nil
}
