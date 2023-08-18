# klog

[![Action][action-svg]][action-url]
[![Report Card][goreport-svg]][goreport-url]
[![godoc][godoc-svg]][godoc-url]
[![License][license-svg]][license-url]

## Use

```golang
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
```

## License

© ender, 2023~time.Now

[MIT LICENSE](https://github.com/xuender/klog/blob/master/LICENSE)

[action-url]: https://github.com/xuender/klog/actions
[action-svg]: https://github.com/xuender/klog/workflows/Go/badge.svg

[goreport-url]: https://goreportcard.com/report/github.com/xuender/klog
[goreport-svg]: https://goreportcard.com/badge/github.com/xuender/klog

[godoc-url]: https://godoc.org/github.com/xuender/klog
[godoc-svg]: https://godoc.org/github.com/xuender/klog?status.svg

[license-url]: https://github.com/xuender/klog/blob/master/LICENSE
[license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
