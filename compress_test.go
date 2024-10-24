// Copyright (c) 2012-2016 The Revel Framework Authors, All rights reserved.
// Revel Framework source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.
package revel

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBenchmarkCompressed(t *testing.T) {
	startFakeBookingApp()
	resp := httptest.NewRecorder()
	showRequest := httptest.NewRequest("GET", "/hotels/show/3", nil)
	resp = httptest.NewRecorder()
	c := NewController(NewRequest(showRequest), NewResponse(resp))

	if err := c.SetAction("Hotels", "Show"); err != nil {
		t.Errorf("SetAction failed: %s", err)
	}

	Config.SetOption("results.compressed", "true")

	result := Hotels{c}.Show(3)
	result.Apply(c.Request, c.Response)

	if !strings.Contains(resp.Body.String(), "300 Main St.") {
		t.Errorf("Failed to find hotel address in action response:\n%s", resp.Body.String())
	}
}

func BenchmarkRenderCompressed(b *testing.B) {
	startFakeBookingApp()

	for i := 0; i < b.N; i++ {

		showRequest := httptest.NewRequest("GET", "/hotels/show/3", nil)
		resp := httptest.NewRecorder()
		c := NewController(NewRequest(showRequest), NewResponse(resp))

		if err := c.SetAction("Hotels", "Show"); err != nil {
			b.Errorf("SetAction failed: %s", err)
		}

		Config.SetOption("results.compressed", "true")

		hotels := Hotels{c}
		hotels.Show(3).Apply(c.Request, c.Response)
	}
}

func BenchmarkRenderUnCompressed(b *testing.B) {
	startFakeBookingApp()

	for i := 0; i < b.N; i++ {

		resp := httptest.NewRecorder()
		c := NewController(NewRequest(showRequest), NewResponse(resp))

		if err := c.SetAction("Hotels", "Show"); err != nil {
			b.Errorf("SetAction failed: %s", err)
		}

		Config.SetOption("results.compressed", "false")

		hotels := Hotels{c}
		hotels.Show(3).Apply(c.Request, c.Response)
	}
}
