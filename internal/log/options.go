package log

import "go.uber.org/zap/zapcore"

type Options struct {
	DisableCaller     bool
	DisableStacktrace bool
	Level             string
	Format            string
	OutputPaths       []string
}

func getOptions() *Options {
	return &Options{
		DisableCaller:     false,
		DisableStacktrace: false,
		Level:             zapcore.InfoLevel.String(),
		Format:            "console",
		OutputPaths:       []string{"stdout"},
	}
}
