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

type ShowBgpSessionsResponse struct {
	InsAPI struct {
		Outputs struct {
			Output ShowBgpSessionsResponseResult `json:"output" xml:"output"`
		} `json:"outputs" xml:"outputs"`
		Sid     string `json:"sid" xml:"sid"`
		Type    string `json:"type" xml:"type"`
		Version string `json:"version" xml:"version"`
	} `json:"ins_api" xml:"ins_api"`
}

type ShowBgpSessionsResponseResult struct {
	Body  ShowBgpSessionsResultBody `json:"body" xml:"body"`
	Code  string                    `json:"code" xml:"code"`
	Input string                    `json:"input" xml:"input"`
	Msg   string                    `json:"msg" xml:"msg"`
}

type ShowBgpSessionsResultBody struct {
	TableVrf []struct {
		RowVrf []struct {
			TableNeighbor []struct {
				RowNeighbor []struct {
					ConnectionsDropped    int      `json:"connectionsdropped" xml:"connectionsdropped"`
					LastFlap              Duration `json:"lastflap" xml:"lastflap"`
					LastRead              Duration `json:"lastread,omitempty" xml:"lastread,omitempty"`
					LastWrite             Duration `json:"lastwrite,omitempty" xml:"lastwrite,omitempty"`
					LocalPort             int      `json:"localport" xml:"localport"`
					NeighborID            string   `json:"neighbor-id" xml:"neighbor-id"`
					NotificationsReceived int      `json:"notificationsreceived" xml:"notificationsreceived"`
					NotificationsSent     int      `json:"notificationssent" xml:"notificationssent"`
					RemoteAS              int      `json:"remoteas" xml:"remoteas"`
					RemotePort            int      `json:"remoteport" xml:"remoteport"`
					State                 string   `json:"state" xml:"state"`
				} `json:"ROW_neighbor" xml:"ROW_neighbor"`
			} `json:"TABLE_neighbor" xml:"TABLE_neighbor"`
			LocalAS             int    `json:"local-as" xml:"local-as"`
			RouterID            string `json:"router-id" xml:"router-id"`
			VrfNameOut          string `json:"vrf-name-out" xml:"vrf-name-out"`
			VrfEstablishedPeers int    `json:"vrfestablishedpeers" xml:"vrfestablishedpeers"`
			VrfPeers            int    `json:"vrfpeers" xml:"vrfpeers"`
		} `json:"ROW_vrf" xml:"ROW_vrf"`
	} `json:"TABLE_vrf" xml:"TABLE_vrf"`
	LocalAS               int `json:"localas" xml:"localas"`
	TotalEstablishedPeers int `json:"totalestablishedpeers" xml:"totalestablishedpeers"`
	TotalPeers            int `json:"totalpeers" xml:"totalpeers"`
}

type ShowBgpSessionsResultFlat struct {
	ConnectionsDropped    int      `json:"connectionsdropped" xml:"connectionsdropped"`
	LastFlap              Duration `json:"lastflap" xml:"lastflap"`
	LastRead              Duration `json:"lastread,omitempty" xml:"lastread,omitempty"`
	LastWrite             Duration `json:"lastwrite,omitempty" xml:"lastwrite,omitempty"`
	LocalPort             int      `json:"localport" xml:"localport"`
	NeighborID            string   `json:"neighbor-id" xml:"neighbor-id"`
	NotificationsReceived int      `json:"notificationsreceived" xml:"notificationsreceived"`
	NotificationsSent     int      `json:"notificationssent" xml:"notificationssent"`
	RemoteAS              int      `json:"remoteas" xml:"remoteas"`
	RemotePort            int      `json:"remoteport" xml:"remoteport"`
	State                 string   `json:"state" xml:"state"`
	LocalAS               int      `json:"local-as" xml:"local-as"`
	RouterID              string   `json:"router-id" xml:"router-id"`
	VrfNameOut            string   `json:"vrf-name-out" xml:"vrf-name-out"`
	VrfEstablishedPeers   int      `json:"vrfestablishedpeers" xml:"vrfestablishedpeers"`
	VrfPeers              int      `json:"vrfpeers" xml:"vrfpeers"`
}

func (d *ShowBgpSessionsResponseResult) Flat() (out []ShowBgpSessionsResultFlat) {
	for _, Tv := range d.Body.TableVrf {
		for _, Rv := range Tv.RowVrf {
			for _, Tn := range Rv.TableNeighbor {
				for _, Rn := range Tn.RowNeighbor {
					out = append(out, ShowBgpSessionsResultFlat{
						ConnectionsDropped:    Rn.ConnectionsDropped,
						LastFlap:              Rn.LastFlap,
						LastRead:              Rn.LastRead,
						LastWrite:             Rn.LastWrite,
						LocalPort:             Rn.LocalPort,
						NeighborID:            Rn.NeighborID,
						NotificationsReceived: Rn.NotificationsReceived,
						NotificationsSent:     Rn.NotificationsSent,
						RemoteAS:              Rn.RemoteAS,
						RemotePort:            Rn.RemotePort,
						State:                 Rn.State,
						LocalAS:               Rv.LocalAS,
						RouterID:              Rv.RouterID,
						VrfNameOut:            Rv.VrfNameOut,
						VrfEstablishedPeers:   Rv.VrfEstablishedPeers,
						VrfPeers:              Rv.VrfPeers,
					})
				}
			}
		}
	}
	return
}

func (d *ShowBgpSessionsResponse) Flat() (out []ShowBgpSessionsResultFlat) {
	return d.InsAPI.Outputs.Output.Flat()
}

// NewShowBgpSessionsFromString returns instance from an input string.
func NewShowBgpSessionsFromString(s string) (*ShowBgpSessionsResponse, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowBgpSessionsFromReader(strings.NewReader(s))
}

// NewShowBgpSessionsFromBytes returns instance from an input byte array.
func NewShowBgpSessionsFromBytes(s []byte) (*ShowBgpSessionsResponse, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowBgpSessionsFromReader(bytes.NewReader(s))
}

// NewShowBgpSessionsFromReader returns instance from an input reader.
func NewShowBgpSessionsFromReader(s io.Reader) (*ShowBgpSessionsResponse, error) {
	//si := &ShowBgpSessions{}
	ShowBgpSessionsResponseDat := &ShowBgpSessionsResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowBgpSessionsResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowBgpSessionsResponseDat, nil
}

// NewShowBgpSessionsResultFromString returns instance from an input string.
func NewShowBgpSessionsResultFromString(s string) (*ShowBgpSessionsResponseResult, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowBgpSessionsResultFromReader(strings.NewReader(s))
}

// NewShowBgpSessionsResultFromBytes returns instance from an input byte array.
func NewShowBgpSessionsResultFromBytes(s []byte) (*ShowBgpSessionsResponseResult, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("missing result")
	}
	return NewShowBgpSessionsResultFromReader(bytes.NewReader(s))
}

// NewShowBgpSessionsResultFromReader returns instance from an input reader.
func NewShowBgpSessionsResultFromReader(s io.Reader) (*ShowBgpSessionsResponseResult, error) {
	//si := &ShowBgpSessionsResponseResult{}
	ShowBgpSessionsResponseResultDat := &ShowBgpSessionsResponseResult{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(ShowBgpSessionsResponseResultDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return ShowBgpSessionsResponseResultDat, nil
}
