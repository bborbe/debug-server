// Copyright (c) 2018 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"

	"github.com/golang/glog"
)

type Handler struct {
	Writer io.Writer
}

func (h *Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	content, err := httputil.DumpRequest(req, true)
	if err != nil {
		glog.Warningf("dump request failed: %v", err)
	}
	fmt.Fprintln(h.Writer, string(content))
	resp.WriteHeader(200)
	fmt.Fprintln(resp, string(content))
}
