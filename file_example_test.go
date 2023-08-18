package klog_test

import (
	"github.com/xuender/klog"
)

func ExampleLogSymlink() {
	writer, _ := klog.LogSymlink("/tmp/log", "link.log")
	defer klog.Close()

	_, _ = writer.Write([]byte("xxx"))
	// Output:
}
