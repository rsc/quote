// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package quote

import "testing"

func TestHello(t *testing.T) {
	hello := "Hello, world."
	if out := Hello(); out != hello {
		t.Errorf("Hello() = %q, want %q", out, hello)
	}
}
