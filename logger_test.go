package main

import (
	"github.com/rezakamalifard/BenchmarkZapLoggers/internal/ztest"
	"testing"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newWriterTestConfig() zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      false,
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func NewWriterSamplerConfig() zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}


func NewWriterLogger(options ...zap.Option) (*zap.Logger, error) {
	return newWriterTestConfig().Build(options...)
}

func NewWriterSamplerLogger(options ...zap.Option) (*zap.Logger, error) {
	return NewWriterSamplerConfig().Build(options...)
}

func newDiscardedLogger(lvl zapcore.Level) *zap.Logger {
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

func newSampledLogger(lvl zapcore.Level) *zap.Logger {
	return zap.New(zapcore.NewSamplerWithOptions(
		newDiscardedLogger(zap.DebugLevel).Core(),
		100*time.Millisecond,
		10, // first
		10, // thereafter
	))
}

func BenchmarkDiscardedSugarLogger(b *testing.B) {
	b.ReportAllocs()
	logger := newDiscardedLogger(zap.DebugLevel)
	sugar := logger.Sugar()
	defer sugar.Sync()
	msg := &Message{"message_text", 1}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sugarLogger(sugar, msg)
	}
}
func BenchmarkDiscardedStructuredLogger(b *testing.B) {
	logger := newDiscardedLogger(zap.DebugLevel)
	defer logger.Sync()
	msg := &Message{"message_text", 1}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		structuredLogger(logger, msg)
	}
}

func BenchmarkSampledSugarLogger(b *testing.B) {
	b.ReportAllocs()
	logger := newSampledLogger(zap.DebugLevel)
	sugar := logger.Sugar()
	defer sugar.Sync()
	msg := &Message{"message_text", 1}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sugarLogger(sugar, msg)
	}
}
func BenchmarkSampledStructuredLogger(b *testing.B) {
	logger := newSampledLogger(zap.DebugLevel)
	defer logger.Sync()
	msg := &Message{"message_text", 1}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		structuredLogger(logger, msg)
	}
}
func BenchmarkWriterSugarLogger(b *testing.B) {
	b.ReportAllocs()
	logger, _ := NewWriterLogger()
	sugar := logger.Sugar()
	defer sugar.Sync()
	msg := &Message{"message_text", 1}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sugarLogger(sugar, msg)
	}
}
func BenchmarkWriterStructuredLogger(b *testing.B) {
	logger, _ := NewWriterLogger()
	defer logger.Sync()
	msg := &Message{"message_text", 1}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		structuredLogger(logger, msg)
	}
}
func BenchmarkWriterSamplerSugarLogger(b *testing.B) {
	b.ReportAllocs()
	logger, _ := NewWriterSamplerLogger()
	sugar := logger.Sugar()
	defer sugar.Sync()
	msg := &Message{"message_text", 1}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sugarLogger(sugar, msg)
	}
}
func BenchmarkWriterSamplerStructuredLogger(b *testing.B) {
	logger, _ := NewWriterSamplerLogger()
	defer logger.Sync()
	msg := &Message{"message_text", 1}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		structuredLogger(logger, msg)
	}
}
