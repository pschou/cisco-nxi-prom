package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"
)

func UploadToCollector(url string, data []byte) (err error) {
	var HTTPClient = &http.Client{
		Timeout: time.Second * 30,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 10 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 10 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	var resp *http.Response
	var req *http.Request

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	req, err = http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(data))
	if err != nil {
		log.Println("  New request error:", err)
	}

	//for key, val := range Headers {
	//	if debug {
	//		fmt.Printf("Request Header: %s: %s\n", key, val)
	//	}
	//	req.Header.Set(key, val)
	//}

	resp, err = HTTPClient.Do(req)
	if err != nil || resp.StatusCode != 200 {
		log.Printf("response:\n %#v\nerror:\n %s\n", resp, err)
	} else {
		log.Println("...pushed")
	}

	return
}
