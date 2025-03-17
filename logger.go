package gfslog

import (
	"context"
	"log/slog"

	"github.com/gogf/gf/v2/os/glog"
	slogcommon "github.com/samber/slog-common"
)

// GoFrameLogHandler
type GoFrameLogHandler struct {
	Option
	attrs  []slog.Attr
	groups []string
}

// NewGoFrameLogHandler
func NewGoFrameLogHandler(logger *glog.Logger, level slog.Leveler, opts ...OptionFunc) slog.Handler {
	if logger == nil {
		logger = glog.New()
	}
	if level == nil {
		level = &slog.LevelVar{}
	}
	option := Option{
		Level:           level,
		Logger:          logger,
		AttrFromContext: []func(ctx context.Context) []slog.Attr{},
	}

	SetOption(&option, opts...)

	return &GoFrameLogHandler{
		Option: option,
		attrs:  []slog.Attr{},
		groups: []string{},
	}

}

var _ slog.Handler = &GoFrameLogHandler{}

// Enabled
func (h *GoFrameLogHandler) Enabled(_ context.Context, level slog.Level) bool {
	if h.Level == nil {
		h.Level = &slog.LevelVar{} // Info level by default.
	}
	return level >= h.Level.Level()
}

// Handle
func (h *GoFrameLogHandler) Handle(ctx context.Context, record slog.Record) error {
	converter := DefaultConverter
	if h.Option.Converter != nil {
		converter = h.Option.Converter
	}

	fromContext := slogcommon.ContextExtractor(ctx, h.Option.AttrFromContext)
	args := converter(h.Option.AddSource, h.Option.ReplaceAttr, append(h.attrs, fromContext...), h.groups, &record)

	return log(ctx, *h.Logger, goframeLevelFunc(record.Level), args, record.Message)
}

// WithAttrs
func (h *GoFrameLogHandler) WithAttrs(slogAttrs []slog.Attr) slog.Handler {
	return &GoFrameLogHandler{
		Option: h.Option,
		attrs:  slogcommon.AppendAttrsToGroup(h.groups, h.attrs, slogAttrs...),
		groups: h.groups,
	}
}

// WithGroup
func (h *GoFrameLogHandler) WithGroup(name string) slog.Handler {
	// https://cs.opensource.google/go/x/exp/+/46b07846:slog/handler.go;l=247
	if name == "" {
		return h
	}

	return &GoFrameLogHandler{
		Option: h.Option,
		attrs:  h.attrs,
		groups: append(h.groups, name),
	}
}
