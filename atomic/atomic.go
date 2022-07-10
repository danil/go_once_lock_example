// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic

import (
	"sync"
	"sync/atomic"
	"time"
)

var once sync.Once
var locked int32

func init() {
	atomic.StoreInt32((*int32)(&locked), 1)
}

func Global() {
	once.Do(func() {
		defer atomic.StoreInt32((*int32)(&locked), 0)
	})

	backoff := time.Nanosecond

	for {
		if atomic.LoadInt32((*int32)(&locked)) == 0 {
			break
		}
		time.Sleep(backoff)
		backoff *= 2
	}
}
