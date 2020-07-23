package main

import (
	"io/ioutil"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// A Syncer is a spy for the Sync portion of zapcore.WriteSyncer.
type Syncer struct {
	err    error
	called bool
}

// SetError sets the error that the Sync method will return.
func (s *Syncer) SetError(err error) {
	s.err = err
}

// Sync records that it was called, then returns the user-supplied error (if
// any).
func (s *Syncer) Sync() error {
	s.called = true
	return s.err
}

// Called reports whether the Sync method was called.
func (s *Syncer) Called() bool {
	return s.called
}

// A Discarder sends all writes to ioutil.Discard.
type Discarder struct{ Syncer }

// Write implements io.Writer.
func (d *Discarder) Write(b []byte) (int, error) {
	return ioutil.Discard.Write(b)
}

func newZapLogger(lvl zapcore.Level) *zap.Logger {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeDuration = zapcore.NanosDurationEncoder
	ec.EncodeTime = zapcore.EpochNanosTimeEncoder
	enc := zapcore.NewJSONEncoder(ec)
	return zap.New(zapcore.NewCore(
		enc,
		&Discarder{},
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
