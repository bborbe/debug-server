// Copyright (c) 2018 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/bborbe/debug-server/debug"
	flag "github.com/bborbe/flagenv"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/golang/glog"
)

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	runtime.GOMAXPROCS(runtime.NumCPU())

	port := flag.Int("port", 8080, "Port")
	_ = flag.Set("logtostderr", "true")
	flag.Parse()

	glog.V(0).Infof("Parameter Port: %d", *port)

	handler := &debug.Handler{
		Writer: os.Stdout,
	}
	if err := gracehttp.Serve(&http.Server{Addr: fmt.Sprintf(":%d", *port), Handler: handler}); err != nil {
		glog.Exitf("serve http failed: %+v", err)
	}
}
