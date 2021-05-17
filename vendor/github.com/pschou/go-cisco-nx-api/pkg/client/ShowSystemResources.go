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

type ShowSystemResourcesResponse struct {
	InsAPI struct {
		Outputs struct {
			Output ShowSystemResourcesResponseResult `json:"output" xml:"output"`
		} `json:"outputs" xml:"outputs"`
		Sid     string `json:"sid" xml:"sid"`
		Type    string `json:"type" xml:"type"`
		Version string `json:"version" xml:"version"`
	} `json:"ins_api" xml:"ins_api"`
}

type ShowSystemResourcesResponseResult struct {
	Body  ShowSystemResourcesResultBody `json:"body" xml:"body"`
	Code  string                        `json:"code" xml:"code"`
	Input string                        `json:"input" xml:"input"`
	Msg   string                        `json:"msg" xml:"msg"`
}

type ShowSystemResourcesResultBody struct {
	LoadAvg1Min         float32 `json:"load_avg_1min" xml:"load_avg_1min"`
	LoadAvg5Min         float32 `json:"load_avg_5min" xml:"load_avg_5min"`
	LoadAvg15Min        float32 `json:"load_avg_15min" xml:"load_avg_15min"`
	ProcessesTotal      int     `json:"processes_total" xml:"processes_total"`
	ProcessesRunning    int     `json:"processes_running" xml:"processes_running"`
	CPUStateUser        float32 `json:"cpu_state_user" xml:"cpu_state_user"`
	CPUStateKernel      float32 `json:"cpu_state_kernel" xml:"cpu_state_kernel"`
	CPUStateIdle        float32 `json:"cpu_state_idle" xml:"cpu_state_idle"`
	MemoryUsageTotal    int     `json:"memory_usage_total" xml:"memory_usage_total"`
	MemoryUsageUsed     int     `json:"memory_usage_used" xml:"memory_usage_used"`
	MemoryUsageFree     int     `json:"memory_usage_free" xml:"memory_usage_free"`
	CurrentMemoryStatus string  `json:"current_memory_status" xml:"current_memory_status"`
	TableCPUUsage       []struct {
		RowCPUUsage []struct {
			Cpuid  int     `json:"cpuid" xml:"cpuid"`
			User   float32 `json:"user" xml:"user"`
			Kernel float32 `json:"kernel" xml:"kernel"`
			Idle   float32 `json:"idle" xml:"idle"`
		} `json:"ROW_cpu_usage" xml:"ROW_cpu_usage"`
	} `json:"TABLE_cpu_usage" xml:"TABLE_cpu_usage"`
}

// NewShowSystemResourcesFromString returns instance from an input string.
func NewShowSystemResourcesFromString(s string) (*ShowSystemResourcesResponse, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowSystemResourcesFromReader(strings.NewReader(s))
}

// NewShowSystemResourcesFromBytes returns instance from an input byte array.
func NewShowSystemResourcesFromBytes(s []byte) (*ShowSystemResourcesResponse, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowSystemResourcesFromReader(bytes.NewReader(s))
}

// NewShowSystemResourcesFromReader returns instance from an input reader.
func NewShowSystemResourcesFromReader(s io.Reader) (*ShowSystemResourcesResponse, error) {
	//si := &ShowSystemResources{}
	ShowSystemResourcesResponseDat := &ShowSystemResourcesResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowSystemResourcesResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowSystemResourcesResponseDat, nil
}

// NewShowSystemResourcesResultFromString returns instance from an input string.
func NewShowSystemResourcesResultFromString(s string) (*ShowSystemResourcesResponseResult, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowSystemResourcesResultFromReader(strings.NewReader(s))
}

// NewShowSystemResourcesResultFromBytes returns instance from an input byte array.
func NewShowSystemResourcesResultFromBytes(s []byte) (*ShowSystemResourcesResponseResult, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowSystemResourcesResultFromReader(bytes.NewReader(s))
}

// NewShowSystemResourcesResultFromReader returns instance from an input reader.
func NewShowSystemResourcesResultFromReader(s io.Reader) (*ShowSystemResourcesResponseResult, error) {
	//si := &ShowSystemResourcesResponseResult{}
	ShowSystemResourcesResponseResultDat := &ShowSystemResourcesResponseResult{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowSystemResourcesResponseResultDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowSystemResourcesResponseResultDat, nil
}
