package main

import (
	"log/slog"
	"os"

	"github.com/xuender/klog"
)

func main() {
	klog.SetLogFile(os.TempDir(), "test.log")
	klog.SetLevel(slog.LevelDebug)

	slog.Debug("debug")
	slog.Info("info")
}
