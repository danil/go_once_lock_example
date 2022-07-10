// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rwmutex

import "sync"

var once sync.Once
var locker sync.RWMutex

func init() {
	locker.Lock()
}

func Global() {
	once.Do(func() {
		defer locker.Unlock()
	})

	locker.RLock()
	defer locker.RUnlock()
}
