package logfmt

import (
	"cloud.google.com/go/logging"
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func FromCtx(ctx context.Context) *zerolog.Logger {
	instance := zerolog.Ctx(ctx)
	if instance.GetLevel() == zerolog.Disabled {
		instance = &log.Logger
	}

	return instance
}
func WithSeverity(ctx context.Context, severity logging.Severity) *zerolog.Event {
	instance := FromCtx(ctx)

	var ev *zerolog.Event
	switch severity {
	case logging.Info, logging.Notice:
		ev = instance.Info()
	case logging.Warning, logging.Alert:
		ev = instance.Warn()
	case logging.Error, logging.Critical, logging.Emergency:
		ev = instance.Error()
	default:
		ev = instance.Debug()
	}

	return ev.Str("severity", severity.String())
}

// Default will return chainable zerolog event with Default severity.
func Default(ctx context.Context) *zerolog.Event {
	return WithSeverity(ctx, logging.Default)
}

// Debug will return chainable zerolog event with Debug severity.
func Debug(ctx context.Context) *zerolog.Event {
	return WithSeverity(ctx, logging.Debug)
}

// Info will return chainable zerolog event with Info severity.
func Info(ctx context.Context) *zerolog.Event {
	return WithSeverity(ctx, logging.Info)
}

// Notice will return chainable zerolog event with Notice severity.
func Notice(ctx context.Context) *zerolog.Event {
	return WithSeverity(ctx, logging.Notice)
}

// Warning will return chainable zerolog event with Warning severity.
func Warning(ctx context.Context) *zerolog.Event {
	return WithSeverity(ctx, logging.Warning)
}

// Error will return chainable zerolog event with Error severity.
func Error(ctx context.Context) *zerolog.Event {
	return WithSeverity(ctx, logging.Error)
}

// Critical will return chainable zerolog event with Critical severity.
func Critical(ctx context.Context) *zerolog.Event {
	return WithSeverity(ctx, logging.Critical)
}

// Alert will return chainable zerolog event with Alert severity.
func Alert(ctx context.Context) *zerolog.Event {
	return WithSeverity(ctx, logging.Alert)
}

// Emergency will return chainable zerolog event with Emergency severity.
func Emergency(ctx context.Context) *zerolog.Event {
	return WithSeverity(ctx, logging.Emergency)
}

// PrintErr is a convenient method to easily logfmt a structured error.
func PrintErr(ctx context.Context, err error) {
	Error(ctx).Err(err).Msg(err.Error())
}
