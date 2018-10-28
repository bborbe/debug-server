// Copyright (c) 2018 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/bborbe/debug-server/debug"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Debug Handler", func() {
	It("return 200 and ok", func() {
		buf := &bytes.Buffer{}
		handler := &debug.Handler{
			Writer: buf,
		}
		recorder := &httptest.ResponseRecorder{}
		handler.ServeHTTP(recorder, &http.Request{
			Body:   ioutil.NopCloser(bytes.NewBufferString("hello world")),
			Header: map[string][]string{"foo": {"bar"}},
		})
		response := recorder.Result()
		Expect(response.StatusCode).To(Equal(200))
		content, err := ioutil.ReadAll(response.Body)
		Expect(err).To(BeNil())
		Expect(string(content)).To(Equal("ok"))
	})
})
