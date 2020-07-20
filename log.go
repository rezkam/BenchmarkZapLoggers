package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Message is a general user defined type that we can use.
type Message struct {
	s string
	i int64
}

// Add MarshalLogObject to implement ObjectEncoder interface to Marshal Message type for logging.
func (m *Message) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("s", m.s)
	enc.AddInt64("index", m.i)
	return nil
}

func main() {
	// we create a new logger and set the environment for it (production)
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	// initialize sugar logger
	sugar := logger.Sugar()

	// create new message
	msg := &Message{"message_text", 1}
	sugarLogger(sugar, msg)
	structuredLogger(logger, msg)
}

func sugarLogger(sugar *zap.SugaredLogger, msg *Message) {
	sugar.Info("Sugar logger",
		"type", "sugar",
		"status", "OK",
		"error", 0,
		"message", msg,
	)
}

func structuredLogger(logger *zap.Logger, msg *Message) {
	logger.Info("Sugar logger",
		zap.String("type", "sugar"),
		zap.String("status", "OK"),
		zap.Int("error", 0),
		zap.Object("message", msg))
}
