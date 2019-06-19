package zaplog

import (
	"os"
	"path"
	"strings"
	"time"

	"github.com/abaole/gframe/logger/conf"
	"github.com/abaole/gframe/logger/fileout"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Log struct {
	logger *zap.Logger
}

//var logger *zap.loghook //全局日志

func parseLevel(lvl string) zapcore.Level {
	switch strings.ToLower(lvl) {
	case "panic", "dpanic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	case "error":
		return zapcore.ErrorLevel
	case "warn", "warning":
		return zapcore.WarnLevel
	case "info":
		return zapcore.InfoLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.DebugLevel
	}
}

//创建日志
func New(opts ...conf.Option) *Log {
	o := &conf.Options{
		LogPath:     conf.LogPath,
		SvcName:     conf.SvcName,
		LogLevel:    conf.LogLevel,
		MaxSize:     conf.MaxSize,
		MaxAge:      conf.MaxAge,
		Stacktrace:  conf.Stacktrace,
		IsStdOut:    conf.IsStdOut,
		ProjectName: conf.ProjectName,
	}
	for _, opt := range opts {
		opt(o)
	}

	directory := path.Join(o.LogPath, o.SvcName)

	writers := []zapcore.WriteSyncer{fileout.NewRollingFile(directory, o.MaxSize, o.MaxAge)}
	if o.IsStdOut == "yes" {
		writers = append(writers, os.Stdout)
	}
	logger := newZapLogger(o.IsProduction, parseLevel(o.Stacktrace), zapcore.NewMultiWriteSyncer(writers...))
	zap.RedirectStdLog(logger)
	logger = logger.With(zap.String("project", o.ProjectName)) //加上项目名称
	return &Log{logger: logger}
}

func newZapLogger(isProduction bool, stacktrace zapcore.Level, output zapcore.WriteSyncer) *zap.Logger {
	encCfg := zapcore.EncoderConfig{
		TimeKey:        "@timestamp",
		LevelKey:       "level",
		NameKey:        "app",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
		},
	}
	var encoder zapcore.Encoder
	dyn := zap.NewAtomicLevel()

	if isProduction {
		dyn.SetLevel(zap.InfoLevel)
		encCfg.EncodeLevel = zapcore.LowercaseLevelEncoder
		encoder = zapcore.NewJSONEncoder(encCfg)
	} else {
		dyn.SetLevel(zap.DebugLevel)
		encCfg.EncodeLevel = zapcore.LowercaseColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encCfg)
	}

	return zap.New(zapcore.NewCore(encoder, output, dyn), zap.AddCaller(), zap.AddStacktrace(stacktrace), zap.AddCallerSkip(2))
}
