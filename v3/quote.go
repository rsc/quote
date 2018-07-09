// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package quote collects pithy sayings.
package quote // import "rsc.io/quote"

import "rsc.io/sampler"

// Hello returns a greeting.
func HelloV3() string {
	return sampler.Hello()
