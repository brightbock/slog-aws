# AWS Lambda advanced logging controls for structured logs in Go with native "log/slog" module

AWS Lambda supports advanced logging controls[1] that help you manage how your function's logs are captured, processed, and consumed.

1: [https://docs.aws.amazon.com/lambda/latest/dg/monitoring-cloudwatchlogs.html#monitoring-cloudwatchlogs-advanced](https://docs.aws.amazon.com/lambda/latest/dg/monitoring-cloudwatchlogs.html#monitoring-cloudwatchlogs-advanced)

---

### This module:

1. includes a shortcut `LambdaLogger()` for setting JSON log format in slog.
1. overrides the name of the time key to `timestamp` (note: not necessary for Lambda).
2. set the log-level in your application to match the setting in Lambda (note: Lambda will filter logs according using the `level` field in each log and the log-level you set in Lambda function config any way)
3. provides a `FatalError(...)` function that will `slog.Error(...)` and then `os.Exit(1)`

### How to use it:

```
package main

import (
    "github.com/brightbock/slogcloud"
    . "github.com/brightbock/slogcloud/helpers"
    "log/slog"
)

func init() {
    slog.SetDefault(slogcloud.LambdaLogger())
}

func main() {
    slog.Warn("This is a warning!")
    FatalError("log and exit with status 1")
}
```

Output:
```
{"timestamp":"2023-01-01T01:01:01.000000+00:00","level":"WARN","msg":"This is a warning!"}
{"timestamp":"2023-01-01T01:01:01.000000+00:00","level":"ERROR","msg":"log and exit with status 1"}
```
