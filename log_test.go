package main

import (
	"go.uber.org/zap"
	"testing"
)

func BenchmarkSugarLogger(b *testing.B) {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	msg := &Message{"message_text", 1}

	for i := 0; i < b.N; i++ {
		sugarLogger(sugar, msg)
	}
}
func BenchmarkStructuredLogger(b *testing.B) {
	logger, _ := zap.NewProduction()
	msg := &Message{"message_text", 1}

	for i := 0; i < b.N; i++ {
		structuredLogger(logger, msg)
	}
}
