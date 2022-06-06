package logs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/mertdogan12/cn/internal/conf"
)

type loki struct {
	Streams []struct {
		Stream struct {
			Job string `json:"job"`
		} `json:"stream"`
		Values [][]string `json:"values"`
	} `json:"streams"`
}

func lokiLog(message string) {
	if !*conf.LOKI {
		return
	}

	timeNano := strconv.Itoa(int(time.Now().UnixNano()))
	obj := loki{
		Streams: []struct {
			Stream struct {
				Job string "json:\"job\""
			} "json:\"stream\""
			Values [][]string "json:\"values\""
		}{
			{
				Stream: struct {
					Job string "json:\"job\""
				}{
					Job: "cn",
				},
				Values: [][]string{
					{
						timeNano,
						message,
					},
				},
			},
		},
	}

	body, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, conf.LOKI_URI+"loki/api/v1/push", bytes.NewReader(body))
	if err != nil {
		panic(err)
	}

	req.Header.Add("X-Scope-OrgID", "tenant1")
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		var respBody []byte
		respBody, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		fmt.Printf("not 200: %s\n", string(respBody))
	}
}

func Log(log string) {
	fmt.Println(log)
	message := "log: " + log
	go lokiLog(message)
}

func LogErr(err error) {
	fmt.Println(err)
	message := "error: " + err.Error()
	go lokiLog(message)
}
