package main

import (
	"github.com/rezakamalifard/BenchmarkZapLoggers/internal/ztest"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newZapLogger(lvl zapcore.Level) *zap.Logger {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeDuration = zapcore.NanosDurationEncoder
	ec.EncodeTime = zapcore.EpochNanosTimeEncoder
	enc := zapcore.NewJSONEncoder(ec)
	return zap.New(zapcore.NewCore(
		enc,
		&ztest.Discarder{},
		lvl,
	))
}

func BenchmarkSugarLogger(b *testing.B) {
	b.ReportAllocs()
	logger := newZapLogger(zap.DebugLevel)
	sugar := logger.Sugar()
	defer sugar.Sync()
	msg := &Message{"message_text", 1}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sugarLogger(sugar, msg)
	}
}
func BenchmarkStructuredLogger(b *testing.B) {
	logger := newZapLogger(zap.DebugLevel)
	defer logger.Sync()
	msg := &Message{"message_text", 1}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		structuredLogger(logger, msg)
	}
}
