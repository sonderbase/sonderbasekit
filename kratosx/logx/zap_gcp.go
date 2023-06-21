package logx

import (
	kzap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ToGCP converts logger adhering to GCP format
// refer: https://github.com/uber-go/zap/discussions/1110
func ToGCP(
	id string,
	name string,
	version string,
) (log.Logger, error) {
	zCfg := &zap.Config{
		Level:    zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:       "time",
			LevelKey:      "severity",
			NameKey:       "logger",
			CallerKey:     "caller",
			MessageKey:    "",
			StacktraceKey: "stacktrace",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel: func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
				switch l {
				case zapcore.DebugLevel:
					enc.AppendString("DEBUG")
				case zapcore.InfoLevel:
					enc.AppendString("INFO")
				case zapcore.WarnLevel:
					enc.AppendString("WARNING")
				case zapcore.ErrorLevel:
					enc.AppendString("ERROR")
				case zapcore.DPanicLevel:
					enc.AppendString("CRITICAL")
				case zapcore.PanicLevel:
					enc.AppendString("ALERT")
				case zapcore.FatalLevel:
					enc.AppendString("EMERGENCY")
				}
			},
			EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
			EncodeDuration: zapcore.MillisDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	z, err := zCfg.Build(
		zap.AddCallerSkip(3),
	)
	if err != nil {
		return nil, err
	}

	logger := log.With(
		kzap.NewLogger(z),
		// these are already taken care by zap
		// "ts", log.DefaultTimestamp,
		// "caller", log.DefaultCaller,
		"service.id", id,
		"service.name", name,
		"service.version", version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	log.SetLogger(logger)
	log.DefaultLogger = logger
	return logger, nil
}
