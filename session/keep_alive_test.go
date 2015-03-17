/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

package session

import (
	"fmt"
	"net/url"
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/vmware/govmomi/test"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	"golang.org/x/net/context"
)

type testKeepAlive int

func (t *testKeepAlive) Func(soap.RoundTripper) {
	*t++
}

func newManager(t *testing.T) (*Manager, *url.URL) {
	u := test.URL()
	if u == nil {
		t.SkipNow()
	}

	soapClient := soap.NewClient(u, true)
	vimClient, err := vim25.NewClient(context.Background(), soapClient)
	if err != nil {
		t.Fatal(err)
	}

	return NewManager(vimClient), u
}

func TestKeepAlive(t *testing.T) {
	var i testKeepAlive
	var j int

	m, u := newManager(t)
	k := KeepAlive(m.client.RoundTripper, time.Millisecond)
	k.(*keepAlive).keepAlive = i.Func
	m.client.RoundTripper = k

	// Expect keep alive to not have triggered yet
	if i != 0 {
		t.Errorf("Expected i == 0, got i: %d", i)
	}

	// Logging in starts keep alive
	err := m.Login(context.Background(), u.User)
	if err != nil {
		t.Error(err)
	}

	time.Sleep(2 * time.Millisecond)

	// Expect keep alive to triggered at least once
	if i == 0 {
		t.Errorf("Expected i != 0, got i: %d", i)
	}

	j = int(i)
	time.Sleep(2 * time.Millisecond)

	// Expect keep alive to triggered at least once more
	if int(i) <= j {
		t.Errorf("Expected i > j, got i: %d, j: %d", i, j)
	}

	// Logging out stops keep alive
	err = m.Logout(context.Background())
	if err != nil {
		t.Error(err)
	}

	j = int(i)
	time.Sleep(2 * time.Millisecond)

	// Expect keep alive to have stopped
	if int(i) != j {
		t.Errorf("Expected i == j, got i: %d, j: %d", i, j)
	}
}

func testSessionOK(t *testing.T, m *Manager, ok bool) {
	s, err := m.UserSession(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	_, file, line, _ := runtime.Caller(1)
	prefix := fmt.Sprintf("%s:%d", file, line)

	if ok && s == nil {
		t.Fatalf("%s: Expected session to be OK, but is invalid", prefix)
	}

	if !ok && s != nil {
		t.Fatalf("%s: Expected session to be invalid, but is OK", prefix)
	}
}

// Run with:
//
//   env GOVMOMI_KEEPALIVE_TEST=1 go test -timeout=60m -run TestRealKeepAlive
//
func TestRealKeepAlive(t *testing.T) {
	if os.Getenv("GOVMOMI_KEEPALIVE_TEST") != "1" {
		t.SkipNow()
	}

	m1, u1 := newManager(t)
	m2, u2 := newManager(t)

	// Enable keepalive on m2
	k := KeepAlive(m2.client.RoundTripper, 10*time.Minute)
	m2.client.RoundTripper = k

	// Expect both sessions to be invalid
	testSessionOK(t, m1, false)
	testSessionOK(t, m2, false)

	// Logging in starts keep alive
	if err := m1.Login(context.Background(), u1.User); err != nil {
		t.Error(err)
	}
	if err := m2.Login(context.Background(), u2.User); err != nil {
		t.Error(err)
	}

	// Expect both sessions to be valid
	testSessionOK(t, m1, true)
	testSessionOK(t, m2, true)

	// Wait for m1 to time out
	delay := 31 * time.Minute
	fmt.Printf("%s: Waiting %d minutes for session to time out...\n", time.Now(), int(delay.Minutes()))
	time.Sleep(delay)

	// Expect m1's session to be invalid, m2's session to be valid
	testSessionOK(t, m1, false)
	testSessionOK(t, m2, true)
}
