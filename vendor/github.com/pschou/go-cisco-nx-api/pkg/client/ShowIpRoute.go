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

type ShowIpRouteResponse struct {
	InsAPI struct {
		Outputs struct {
			Output ShowIpRouteResponseResult `json:"output" xml:"output"`
		} `json:"outputs" xml:"outputs"`
		Sid     string `json:"sid" xml:"sid"`
		Type    string `json:"type" xml:"type"`
		Version string `json:"version" xml:"version"`
	} `json:"ins_api" xml:"ins_api"`
}

type ShowIpRouteResponseResult struct {
	Body  ShowIpRouteResultBody `json:"body" xml:"body"`
	Code  string                `json:"code" xml:"code"`
	Input string                `json:"input" xml:"input"`
	Msg   string                `json:"msg" xml:"msg"`
}

type ShowIpRouteResultBody struct {
	TableVrf []struct {
		RowVrf []struct {
			TableAddrf []struct {
				RowAddrf []struct {
					TablePrefix []struct {
						RowPrefix []struct {
							TablePath []struct {
								RowPath []struct {
									ClientName string   `json:"clientname" xml:"clientname"`
									IfName     string   `json:"ifname" xml:"ifname"`
									Metric     int      `json:"metric" xml:"metric"`
									Pref       int      `json:"pref" xml:"pref"`
									UBest      bool     `json:"ubest" xml:"ubest"`
									UpTime     Duration `json:"uptime" xml:"uptime"`
								} `json:"ROW_path" xml:"ROW_path"`
							} `json:"TABLE_path" xml:"TABLE_path"`
							Attached   bool   `json:"attached" xml:"attached"`
							IPPrefix   string `json:"ipprefix" xml:"ipprefix"`
							MCastNHops int    `json:"mcast-nhops" xml:"mcast-nhops"`
							UCastNHops int    `json:"ucast-nhops" xml:"ucast-nhops"`
						} `json:"ROW_prefix" xml:"ROW_prefix"`
					} `json:"TABLE_prefix" xml:"TABLE_prefix"`
					AddRf string `json:"addrf" xml:"addrf"`
				} `json:"ROW_addrf" xml:"ROW_addrf"`
			} `json:"TABLE_addrf" xml:"TABLE_addrf"`
			VrfNameOut string `json:"vrf-name-out" xml:"vrf-name-out"`
		} `json:"ROW_vrf" xml:"ROW_vrf"`
	} `json:"TABLE_vrf" xml:"TABLE_vrf"`
}

type ShowIpRouteResultFlat struct {
	ClientName string   `json:"clientname" xml:"clientname"`
	IfName     string   `json:"ifname" xml:"ifname"`
	Metric     int      `json:"metric" xml:"metric"`
	Pref       int      `json:"pref" xml:"pref"`
	UBest      bool     `json:"ubest" xml:"ubest"`
	UpTime     Duration `json:"uptime" xml:"uptime"`
	Attached   bool     `json:"attached" xml:"attached"`
	IPPrefix   string   `json:"ipprefix" xml:"ipprefix"`
	MCastNHops int      `json:"mcast-nhops" xml:"mcast-nhops"`
	UCastNHops int      `json:"ucast-nhops" xml:"ucast-nhops"`
	AddRf      string   `json:"addrf" xml:"addrf"`
	VrfNameOut string   `json:"vrf-name-out" xml:"vrf-name-out"`
}

func (d *ShowIpRouteResponse) Flat() (out []ShowIpRouteResultFlat) {
	return d.InsAPI.Outputs.Output.Flat()
}
func (d *ShowIpRouteResponseResult) Flat() (out []ShowIpRouteResultFlat) {
	for _, Tv := range d.Body.TableVrf {
		for _, Rv := range Tv.RowVrf {
			for _, Ta := range Rv.TableAddrf {
				for _, Ra := range Ta.RowAddrf {
					for _, Tpre := range Ra.TablePrefix {
						for _, Rpre := range Tpre.RowPrefix {
							for _, Tp := range Rpre.TablePath {
								for _, Rp := range Tp.RowPath {
									out = append(out, ShowIpRouteResultFlat{
										ClientName: Rp.ClientName,
										IfName:     Rp.IfName,
										Metric:     Rp.Metric,
										Pref:       Rp.Pref,
										UBest:      Rp.UBest,
										UpTime:     Rp.UpTime,
										Attached:   Rpre.Attached,
										IPPrefix:   Rpre.IPPrefix,
										MCastNHops: Rpre.MCastNHops,
										UCastNHops: Rpre.UCastNHops,
										AddRf:      Ra.AddRf,
										VrfNameOut: Rv.VrfNameOut,
									})
								}
							}
						}
					}
				}
			}
		}
	}
	return
}

// NewShowIpRouteFromString returns instance from an input string.
func NewShowIpRouteFromString(s string) (*ShowIpRouteResponse, error) { if len(s) == 0 { return nil, fmt.Errorf("missing result") }
	return NewShowIpRouteFromReader(strings.NewReader(s))
}

// NewShowIpRouteFromBytes returns instance from an input byte array.
func NewShowIpRouteFromBytes(s []byte) (*ShowIpRouteResponse, error) { if len(s) == 0 { return nil, fmt.Errorf("missing result") }
	return NewShowIpRouteFromReader(bytes.NewReader(s))
}

// NewShowIpRouteFromReader returns instance from an input reader.
func NewShowIpRouteFromReader(s io.Reader) (*ShowIpRouteResponse, error) {
	//si := &ShowIpRoute{}
	ShowIpRouteResponseDat := &ShowIpRouteResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowIpRouteResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowIpRouteResponseDat, nil
}

// NewShowIpRouteResultFromString returns instance from an input string.
func NewShowIpRouteResultFromString(s string) (*ShowIpRouteResponseResult, error) { if len(s) == 0 { return nil, fmt.Errorf("missing result") }
	return NewShowIpRouteResultFromReader(strings.NewReader(s))
}

// NewShowIpRouteResultFromBytes returns instance from an input byte array.
func NewShowIpRouteResultFromBytes(s []byte) (*ShowIpRouteResponseResult, error) { if len(s) == 0 { return nil, fmt.Errorf("missing result") }
	return NewShowIpRouteResultFromReader(bytes.NewReader(s))
}

// NewShowIpRouteResultFromReader returns instance from an input reader.
func NewShowIpRouteResultFromReader(s io.Reader) (*ShowIpRouteResponseResult, error) {
	//si := &ShowIpRouteResponseResult{}
	ShowIpRouteResponseResultDat := &ShowIpRouteResponseResult{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowIpRouteResponseResultDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowIpRouteResponseResultDat, nil
}
