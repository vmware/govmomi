/*
Copyright (c) 2014-2015 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package importx

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestRemoteOpen(t *testing.T) {
	cacheDir, _ = ioutil.TempDir("/tmp", "govc-remote-open")
	_ = os.MkdirAll(cacheDir, 0755)
	defer os.RemoveAll(cacheDir)

	ts := httptest.NewServer(http.FileServer(http.Dir(".")))
	defer ts.Close()

	d := Downloader{}

	var tests = []struct {
		path   string
		remote bool
		exist  bool
	}{
		{"archive.go", true, true},
		{"not-exist", true, false},
	}

	for _, test := range tests {
		path := test.path
		if test.remote {
			path = ts.URL + "/" + path
		}

		f, _, err := d.Ensure(path)
		if !test.exist {
			if err == nil {
				t.Fatal("expect error; got nil")
			}
			continue
		}
		if err != nil {
			t.Fatal(err)
		}

		expect, err := ioutil.ReadFile(test.path)
		if err != nil {
			t.Fatal(err)
		}

		got, err := ioutil.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(expect, got) {
			t.Fatalf("expect same contents; got %s, %s", expect, got)
		}

		if test.remote {
			_, err = os.Stat(cachePath(path))
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}
