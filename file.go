package klog

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// nolint: gochecknoglobals
var _files = sync.Map{}

const (
	_defaultFileMode   os.FileMode = 0o664
	_defaultDirFileMod os.FileMode = 0o771
)

// CloseFile log file.
func CloseFile(file *os.File) error {
	_ = file.Sync()

	if info, err := file.Stat(); err == nil && info.Size() == 0 {
		if err := os.Remove(file.Name()); err != nil {
			return err
		}
	}

	return file.Close()
}

// Close all log file.
func Close() error {
	_files.Range(func(key, value any) bool {
		if file, ok := value.(*os.File); ok {
			if err := CloseFile(file); err != nil {
				return false
			}
		}

		return true
	})

	return nil
}

// LogSymlink log Symlink.
func LogSymlink(path, name string) (io.Writer, error) {
	_ = os.MkdirAll(path, _defaultDirFileMod)

	var (
		ext      = filepath.Ext(name)
		suffix   = time.Now().Format("06010215")
		log      = fmt.Sprintf("%s-%s%s", name[:len(name)-len(ext)], suffix, ext)
		fileName = filepath.Join(path, log)
	)

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, _defaultFileMode)
	if err != nil {
		return nil, err
	}

	link := filepath.Join(path, name)

	if old, has := _files.Load(link); has {
		if file, ok := old.(*os.File); ok {
			_ = CloseFile(file)
		}
	}

	_files.Store(link, file)
	_ = os.Remove(link)
	err = os.Symlink(log, link)

	return file, err
}
