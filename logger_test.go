package gfslog

import (
	"fmt"
	"log/slog"
	"testing"
	"time"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
)

func TestLogger(t *testing.T) {
	ctx := gctx.New()
	logger := slog.New(NewGoFrameLogHandler(glog.New(), slog.LevelDebug))
	logger.Info("xxx")
	logger.LogAttrs(ctx, slog.LevelDebug, "xxxx", slog.Any("REQ", "req"))

	logger.InfoContext(ctx, "message 0")

	logger.Debug("message 1", "key", "value")

	logger.Info("message 2", "a", 1, "long", "has spaces")

	ts := time.Now()

	logger.Warn("message 3", "b", ts)

	logger.WarnContext(ctx, "message 4", "c", 1*time.Minute)

	logger = logger.
		With("environment", "dev").
		With("release", "v1.0.0")

	// log error
	logger.
		With("category", "sql").
		With("query.statement", "SELECT COUNT(*) FROM users;").
		With("query.duration", 1*time.Second).
		With("error", fmt.Errorf("could not count users")).
		Error("caramba!")

	// log user signup
	logger.
		With(
			slog.Group("user",
				slog.String("id", "user-123"),
				slog.Time("created_at", time.Now()),
			),
		).
		Info("user registration")
}
