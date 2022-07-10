// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uberatomic

import (
	"sync"
	"time"

	"go.uber.org/atomic"
)

var once sync.Once
var locked = atomic.NewBool(true)

func Global() {
	once.Do(func() {
		defer locked.Store(false)
	})

	backoff := time.Nanosecond

	for {
		if !locked.Load() {
			break
		}
		time.Sleep(backoff)
		backoff *= 2
	}
}
