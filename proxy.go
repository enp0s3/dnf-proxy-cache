/*
 * This file is part of the DNF proxy cache project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2023 enp0s3 <enp0s8@proton.me>
 *
 */

package main

import (
	"fmt"
	"html"
	"net/http"
	"time"
)

type Proxy interface {
	Start() error
}

type DnfProxy struct {
	port   uint
	server *http.Server
}

func New(port uint) Proxy {
	return DnfProxy{port: port}
}

func (d DnfProxy) mainHandler() {

}

func (d DnfProxy) Start() error {
	d.server = &http.Server{
		Addr:           fmt.Sprintf(":%d", d.port),
		Handler:        http.DefaultServeMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	return d.server.ListenAndServe()
}
