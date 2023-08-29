package klog

import "golang.org/x/exp/slog"

// nolint: gochecknoglobals
var _programLevel = new(slog.LevelVar)

func SetLogFile(path, file string) error {
	writer, err := LogSymlink(path, file)
	if err != nil {
		return err
	}

	handler := slog.NewTextHandler(writer, &slog.HandlerOptions{
		Level: _programLevel,
	})

	slog.SetDefault(slog.New(handler))

	return nil
}

func SetJSONFile(path, file string) error {
	writer, err := LogSymlink(path, file)
	if err != nil {
		return err
	}

	handler := slog.NewJSONHandler(writer, &slog.HandlerOptions{
		Level: _programLevel,
	})

	slog.SetDefault(slog.New(handler))

	return nil
}

func SetLevel(level slog.Level) {
	_programLevel.Set(level)
}
