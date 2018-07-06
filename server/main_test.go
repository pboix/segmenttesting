package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestMockServerSavesPayload(t *testing.T) {
	jsonData := []byte(`
		{
			"integrations":{
			"Facebook Pixel":false,
			"Google Analytics":false
			},
			"context":{
			"page":{
			"path":"/form/r0yu4H/results",
			"referrer":"[https://admin.typeform.com/form/r0yu4H/create",
			"search":"",
			"title":"Results",
			"url":"https://admin.typeform.com/form/r0yu4H/results"
			},
			"userAgent":"Mozilla/5.0](https://admin.typeform.com/form/r0yu4H/create%22,%22search%22:%22%22,%22title%22:%22Results%22,%22url%22:%22https://admin.typeform.com/form/r0yu4H/results%22%7D,%22userAgent%22:%22Mozilla/5.0) (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36",
			"library":{
			"name":"analytics.js",
			"version":"3.6.0"
			}
			},
			"properties":{
			"category":"admin",
			"account_id":8994753,
			"ws_owner_account_id":8994753,
			"typeform_version":"v2",
			"typeform_property":"web_platform",
			"unique_pageview_id":"70a94dc5-24f5-4d64-a806-7541620b1dad",
			"unique_sectionview_id":"445eadcc-203c-4f25-8c60-0419fcd66ab1",
			"tracking_session_id":"4cf0d211-40c7-43dd-b76d-83df441c8abe",
			"page":"results",
			"section":"summary",
			"natero_feature_name":"results_summary",
			"was_displayed_by_default":1
			},
			"event":"view_page_section",
			"anonymousId":"e936c472-7bba-43b1-9c7f-c33c24193999",
			"timestamp":"2018-05-30T07:44:48.733Z",
			"type":"track",
			"writeKey":"5CHFipOVS715Gu3cR2JtJFPrMHc35h4k",
			"userId":8993819,
			"sentAt":"2018-05-30T07:44:48.735Z",
			"_metadata":{
			"bundled":[
			"AdWords",
			"Amplitude",
			"Appcues",
			"Bing Ads",
			"Facebook Pixel",
			"Google Analytics",
			"LinkedIn Insight Tag",
			"[Segment.io](http://segment.io/)"
			],
			"unbundled":[
			]
			},
			"messageId":"ajs-984d0ddfbb9bcc3c5f7061ab9d53de6d"
			}`)

	req, _ := http.NewRequest(http.MethodPost, "https://api.segment.io/v1/t", bytes.NewBuffer(jsonData))
	segmentMockHandler(&responseWriter{}, req)
	var j interface{}
	if err := json.Unmarshal(jsonData, &j); err != nil {
		t.Errorf("Could not unmarshall json fixture (%s)", err)
	}

	if !reflect.DeepEqual(actual.Body, string(jsonData)) {
		t.Errorf("Expected (%s), received (%s)", jsonData, actual.Body)
	}

}

type responseWriter struct{}

func (responseWriter) WriteHeader(int)           {}
func (responseWriter) Write([]byte) (int, error) { return 0, nil }
func (responseWriter) Header() http.Header       { return map[string][]string{} }
