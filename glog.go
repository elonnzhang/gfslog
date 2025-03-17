package gfslog

import (
	"context"
	"log/slog"

	"github.com/gogf/gf/v2/os/glog"
)

func goframeLevelFunc(lvl slog.Level) int {
	var glevel int
	switch lvl {
	case slog.LevelDebug:
		glevel = glog.LEVEL_DEBU
	case slog.LevelInfo:
		glevel = glog.LEVEL_INFO
	case slog.LevelWarn:
		glevel = glog.LEVEL_WARN
	case slog.LevelError:
		glevel = glog.LEVEL_ERRO
	default:
		glevel = glog.LEVEL_DEBU
	}

	return glevel
}

func log(ctx context.Context, logger glog.Logger, level int, msg ...any) error {
	switch level {
	case glog.LEVEL_DEBU:
		logger.Debug(ctx, msg...)
	case glog.LEVEL_INFO:
		logger.Info(ctx, msg...)
	case glog.LEVEL_WARN:
		logger.Warning(ctx, msg...)
	case glog.LEVEL_ERRO:
		logger.SetStackSkip(2)
		logger.Error(ctx, msg...)
	default:
		logger.Info(ctx, msg...)
	}
	return nil
}
