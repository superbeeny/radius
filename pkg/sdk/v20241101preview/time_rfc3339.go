//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package v20241101preview



import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"regexp"
	"strings"
	"time"
)



const (
	utcLayoutJSON = `"2006-01-02T15:04:05.999999999"`
	utcLayout     = "2006-01-02T15:04:05.999999999"
	rfc3339JSON   = `"` + time.RFC3339Nano + `"`
)

// Azure reports time in UTC but it doesn't include the 'Z' time zone suffix in some cases.
var tzOffsetRegex = regexp.MustCompile(`(Z|z|\+|-)(\d+:\d+)*"*$`)

type timeRFC3339 time.Time

func (t timeRFC3339) MarshalJSON() (json []byte, err error) {
	tt := time.Time(t)
	return tt.MarshalJSON()
}

func (t timeRFC3339) MarshalText() (text []byte, err error) {
	tt := time.Time(t)
	return tt.MarshalText()
}

func (t *timeRFC3339) UnmarshalJSON(data []byte) error {
	layout := utcLayoutJSON
	if tzOffsetRegex.Match(data) {
		layout = rfc3339JSON
	}
	return t.Parse(layout, string(data))
}

func (t *timeRFC3339) UnmarshalText(data []byte) (err error) {
	layout := utcLayout
	if tzOffsetRegex.Match(data) {
		layout = time.RFC3339Nano
	}
	return t.Parse(layout, string(data))
}

func (t *timeRFC3339) Parse(layout, value string) error {
	p, err := time.Parse(layout, strings.ToUpper(value))
	*t = timeRFC3339(p)
	return err
}


func populateTimeRFC3339(m map[string]interface{}, k string, t *time.Time) {
	if t == nil {
		return
	} else if azcore.IsNullValue(t) {
		m[k] = nil
		return
	} else if reflect.ValueOf(t).IsNil() {
		return
	}
	m[k] = (*timeRFC3339)(t)
}

func unpopulateTimeRFC3339(data json.RawMessage, fn string, t **time.Time) error {
	if data == nil || strings.EqualFold(string(data), "null") {
		return nil
	}
	var aux timeRFC3339
	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	*t = (*time.Time)(&aux)
	return nil
}
