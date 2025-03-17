package gfslog

import (
	"context"
	"log/slog"

	"github.com/gogf/gf/v2/os/glog"
)

// Option
type Option struct {
	// log level (default: debug)
	Level slog.Leveler

	// optional: glog logger (default: glog.New())
	Logger *glog.Logger

	// optional: customize json payload builder
	Converter Converter
	// optional: fetch attributes from context
	AttrFromContext []func(ctx context.Context) []slog.Attr

	// optional: see slog.HandlerOptions
	AddSource   bool
	ReplaceAttr func(groups []string, a slog.Attr) slog.Attr
}

// OptionFunc
type OptionFunc func(*Option)

// NewOption New with Option
func NewOption(opts ...OptionFunc) *Option {
	o := &Option{
		Level:  slog.LevelDebug, // default level Debug
		Logger: glog.New(),      // default glog
	}
	return SetOption(o, opts...)
}

// SetOption set option
func SetOption(o *Option, opts ...OptionFunc) *Option {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// WithLevel set level
func WithLevel(level slog.Level) OptionFunc {
	return func(o *Option) {
		o.Level = level
	}
}

// WithLogger set logger
func WithLogger(logger *glog.Logger) OptionFunc {
	return func(o *Option) {
		o.Logger = logger
	}
}

// WithConverter set converter
func WithConverter(converter Converter) OptionFunc {
	return func(o *Option) {
		o.Converter = converter
	}
}

// WithAttrFromContext set Context attr
func WithAttrFromContext(f func(context.Context) []slog.Attr) OptionFunc {
	return func(o *Option) {
		o.AttrFromContext = append(o.AttrFromContext, f)
	}
}

// WithAddSource set add source
func WithAddSource(addSource bool) OptionFunc {
	return func(o *Option) {
		o.AddSource = addSource
	}
}

// WithReplaceAttr set slog.Handler ReplaceAttr
func WithReplaceAttr(replaceAttr func(groups []string, a slog.Attr) slog.Attr) OptionFunc {
	return func(o *Option) {
		o.ReplaceAttr = replaceAttr
	}
}
