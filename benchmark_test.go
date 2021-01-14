package wparams_test

import (
	"context"
	"testing"

	wparams "github.com/palantir/witchcraft-go-params"
)

func BenchmarkContextWithSafeParam(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		ctx = wparams.ContextWithSafeParam(ctx, "val1", 1)
		ctx = wparams.ContextWithSafeParam(ctx, "val2", 2)
		ctx = wparams.ContextWithSafeParam(ctx, "val3", 3)
		_, _ = wparams.SafeAndUnsafeParamsFromContext(ctx)
	}
}

func BenchmarkContextWithSafeParams(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		ctx = wparams.ContextWithSafeParams(ctx, map[string]interface{}{
			"val1": 1,
			"val2": 2,
			"val3": 3,
		})
		_, _ = wparams.SafeAndUnsafeParamsFromContext(ctx)
	}
}

func BenchmarkContextWithSafeAndUnsafeParam(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		ctx = wparams.ContextWithSafeParam(ctx, "safe1", 1)
		ctx = wparams.ContextWithSafeParam(ctx, "safe2", 2)
		ctx = wparams.ContextWithSafeParam(ctx, "safe3", 3)
		ctx = wparams.ContextWithUnsafeParam(ctx, "unsafe1", 1)
		ctx = wparams.ContextWithUnsafeParam(ctx, "unsafe2", 2)
		ctx = wparams.ContextWithUnsafeParam(ctx, "unsafe3", 3)
		_, _ = wparams.SafeAndUnsafeParamsFromContext(ctx)
	}
}

func BenchmarkContextWithSafeAndUnsafeParams(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		ctx = wparams.ContextWithSafeParams(ctx, map[string]interface{}{
			"safe1": 1,
			"safe2": 2,
			"safe3": 3,
		})
		ctx = wparams.ContextWithUnsafeParams(ctx, map[string]interface{}{
			"unsafe1": 1,
			"unsafe2": 2,
			"unsafe3": 3,
		})
		_, _ = wparams.SafeAndUnsafeParamsFromContext(ctx)
	}
}
