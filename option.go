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

// OptionFunc 是选项函数的类型，用于修改 Option 结构体
type OptionFunc func(*Option)

// NewOption 使用默认值初始化 Option，并应用所有选项函数
func NewOption(opts ...OptionFunc) *Option {
	o := &Option{
		Level:  slog.LevelDebug, // 默认日志级别为 Debug
		Logger: glog.New(),      // 默认日志记录器
		// 其他默认值（如 Converter、AttrFromContext 等可留空或设置默认值）
	}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// SetOption 设置选项
func SetOption(o *Option, opts ...OptionFunc) *Option {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// WithLevel 设置日志级别
func WithLevel(level slog.Level) OptionFunc {
	return func(o *Option) {
		o.Level = level
	}
}

// WithLogger 设置自定义日志记录器
func WithLogger(logger *glog.Logger) OptionFunc {
	return func(o *Option) {
		o.Logger = logger
	}
}

// WithConverter 设置自定义的 JSON 转换器
func WithConverter(converter Converter) OptionFunc {
	return func(o *Option) {
		o.Converter = converter
	}
}

// WithAttrFromContext 添加从 Context 中提取属性的函数
func WithAttrFromContext(f func(context.Context) []slog.Attr) OptionFunc {
	return func(o *Option) {
		o.AttrFromContext = append(o.AttrFromContext, f)
	}
}

// WithAddSource 配置 slog.Handler 的 AddSource
func WithAddSource(addSource bool) OptionFunc {
	return func(o *Option) {
		o.AddSource = addSource
	}
}

// WithReplaceAttr 配置 slog.Handler 的 ReplaceAttr
func WithReplaceAttr(replaceAttr func(groups []string, a slog.Attr) slog.Attr) OptionFunc {
	return func(o *Option) {
		o.ReplaceAttr = replaceAttr
	}
}
