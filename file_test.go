// nolint: paralleltest
package klog_test

import (
	"os"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/klog"
)

func TestLogSymlink(t *testing.T) {
	patches := gomonkey.ApplyFuncReturn(os.OpenFile, nil, os.ErrClosed)
	defer patches.Reset()

	ass := assert.New(t)
	_, err := klog.LogSymlink(os.TempDir(), "test")

	ass.NotNil(err)
}

func TestCloseFile(t *testing.T) {
	patches := gomonkey.ApplyFuncReturn(os.Remove, os.ErrClosed)
	defer patches.Reset()

	file, _ := os.CreateTemp(os.TempDir(), "test")

	assert.NotNil(t, klog.CloseFile(file))
}
