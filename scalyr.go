package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

const _SCALYR_BASE_ENDPOINT = "https://www.scalyr.com/"
const _SCALYR_ADDEVENTS_ENDPOINT = _SCALYR_BASE_ENDPOINT + "addEvents"

var _SCALYR_KEY = os.Getenv("SCALYR_KEY")
var _APPLICATION_PROCESS_ID = UUID()
var _APPLICATION_HOSTNAME = HostName()

type Severity int

type scalyrEvent struct {
	TS    string      `json:"ts"`
	Type  int         `json:"type"`
	Sev   int         `json:"sev"`
	Attrs interface{} `json:"attrs"`
}

type scalyrSessionInfo struct {
	ServerType string `json:"serverType"`
	ServerId   string `json:"serverId"`
}

type scalyrEvents struct {
	Token       string            `json:"token"`
	Session     string            `json:"session"`
	SessionInfo scalyrSessionInfo `json:"sessionInfo"`
	Events      []scalyrEvent     `json:"events"`
}

const (
	Debug Severity = 2 + iota
	Info
	Warning
	Error
	Fatal
)

var severityLevels = [...]string{
	"Debug",
	"Info",
	"Warning",
	"Error",
	"Fatal",
}

// String returns the English name of the month ("January", "February", ...).
func (sev Severity) String() string { return severityLevels[sev-2] }

func testSeverityLevel(sev Severity) {

	if sev < 2 || sev > 6 {
		panic("Severity value was out of range!")
	}

}

// TODO
//
// event should accept Severity ( 0 - 6 ??, 3 being info:  and a message :  and 1 or more additional attributes)
// The "sev" (severity) field should range from 0 to 6, and identifies the importance of this event, using the
//  classic scale "finest, finer, fine, info, warning, error, fatal". This field is optional (defaults to 3 / info).

func Event(sev Severity, attributes interface{}) {
	//Print(message)
	Print(string(sev))

	//fmt.Println(sev)

	//fmt.Println("scalyr key:", os.Getenv("SCALYR_KEY"))

	// TODO:  test for scalyr api key being present...

	si := &scalyrSessionInfo{
		ServerType: "server type...",
		ServerId:   _APPLICATION_HOSTNAME,
	}

	/*

		var attrs interface{}
		err := json.Unmarshal([]byte("{}"), &attrs)
		m := attrs.(map[string]interface{})
		m["message"] = message

	*/

	//m["completelynewattribute"] = "NEW ATTR"
	//m["this will be a number"] = 42

	for k := range attributes.(map[string]interface{}) {
		Print(k)
	}

	//usethisone := attributes.(map[string]interface{})
	//usethisone["message"] = message

	/*

		b := []byte(`{"message":"","attroneasnumber":6,"attrtwoasstring":"my custom attribute..."}`)

		var attrs interface{}
		err := json.Unmarshal(b, &attrs)

		m := attrs.(map[string]interface{})
		m["message"] = "attribute adding on fly !!"
		m["completelynewattribute"] = "NEW ATTR"
		m["this will be a number"] = 42

		se := &scalyrEvent{
			TS:    strconv.FormatInt(time.Now().UnixNano(), 10),
			Type:  0,
			Sev:   int(sev),
			Attrs: attrs,
		}

	*/

	se := &scalyrEvent{
		TS:    strconv.FormatInt(time.Now().UnixNano(), 10),
		Type:  0,
		Sev:   int(sev),
		Attrs: attributes,
	}

	events := make([]scalyrEvent, 1)
	events[0] = *se

	ses := &scalyrEvents{
		Token:       _SCALYR_KEY,
		Session:     _APPLICATION_PROCESS_ID,
		SessionInfo: *si,
		Events:      events,
	}

	json, _ := json.Marshal(ses)
	fmt.Println(string(json))

	url := _SCALYR_ADDEVENTS_ENDPOINT
	fmt.Println("URL:>", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}
