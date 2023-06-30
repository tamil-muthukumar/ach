// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestFileWrite(t *testing.T) {
	dir, err := os.MkdirTemp("", "ach-writeACH-test")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer os.RemoveAll(dir)

	path := filepath.Join(dir, fmt.Sprintf("%s.ach", time.Now().UTC().Format("200601021504")))
	write(path)

	s, err := os.Stat(path)
	if err != nil {
		t.Fatal(err.Error())
	}
	if s.Size() <= 0 {
		t.Fatal("expected non-empty file")
	}
}
