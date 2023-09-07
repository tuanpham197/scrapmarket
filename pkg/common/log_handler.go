package common

import (
	"fmt"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SugarLog() *zap.SugaredLogger {
	writerSyncer := getLogWriter()
	encoder := getEncoder()

	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)

	logger := zap.New(core)
	defer logger.Sync()

	return logger.Sugar()
}

func NewProductionEncoderConfigCustom() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:          "ts",
		LevelKey:         "level",
		NameKey:          "logger",
		CallerKey:        "caller",
		FunctionKey:      zapcore.OmitKey,
		MessageKey:       "msg",
		StacktraceKey:    "stacktrace",
		LineEnding:       zapcore.DefaultLineEnding,
		ConsoleSeparator: "  |  ",
		EncodeLevel:      zapcore.CapitalLevelEncoder,
		EncodeTime:       zapcore.TimeEncoderOfLayout(time.RFC3339),
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
	}
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(NewProductionEncoderConfigCustom())
}

func getLogWriter() zapcore.WriteSyncer {
	layout := "01-02-2006"
	t := time.Now()
	logFolder := "logs"
	if _, err := os.Stat(logFolder); os.IsNotExist(err) {
		if err := os.Mkdir(logFolder, 0755); err != nil {
			log.Fatalf("Failed to create log folder: %s", err.Error())
		}
	}

	logFile := "log-" + t.Format(layout) + ".txt"
	path := fmt.Sprintf("./%s/%s", logFolder, logFile)
	file, _ := os.Create(path)

	return zapcore.AddSync(file)
}
