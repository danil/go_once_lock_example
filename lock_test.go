// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oncelockexample

import (
	"testing"

	"github.com/danil/oncelockexample/atomic"
	"github.com/danil/oncelockexample/channel"
	"github.com/danil/oncelockexample/mutex"
	"github.com/danil/oncelockexample/rwmutex"
	"github.com/danil/oncelockexample/uberatomic"
)

func BenchmarkAtomic(b *testing.B) {
	b.ReportAllocs()

	b.Run("atomic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			atomic.Global()
		}
	})
}

func BenchmarkUberAtomic(b *testing.B) {
	b.ReportAllocs()

	b.Run("uber atomic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			uberatomic.Global()
		}
	})
}

func BenchmarkMutex(b *testing.B) {
	b.ReportAllocs()

	b.Run("mutex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mutex.Global()
		}
	})
}

func BenchmarkRWMutex(b *testing.B) {
	b.ReportAllocs()

	b.Run("read-write mutex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			rwmutex.Global()
		}
	})
}

func BenchmarkChannel(b *testing.B) {
	b.ReportAllocs()

	b.Run("channel", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			channel.Global()
		}
	})
}
