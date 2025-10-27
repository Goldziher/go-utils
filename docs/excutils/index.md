# excutils

Exception-style error handling utilities for Go.

## Overview

The `excutils` package (imported as `exc`) provides utilities for common error handling patterns including panic-on-error, error recovery, and safe error checking. While Go favors explicit error returns, these utilities can make certain patterns more concise.

## Functions

**Panic on Error**: Must, MustResult, MustNotNil, ReturnNotNil
**Recovery**: Try, Catch, RecoverWithValue
**Error Utilities**: FirstErr, AllErr, ReturnAnyErr, IgnoreErr
**Retry**: Retry, RetryWithResult

## Example

```go
import exc "github.com/Goldziher/go-utils/excutils"

// Panic on error (for setup/initialization)
file := exc.MustResult(os.Open("config.json"), "failed to open config")

// Safe recovery
err := exc.Try(func() error {
	// Code that might panic or error
	return riskyOperation()
})

// Get first error
err := exc.FirstErr(err1, err2, err3)

// Retry with exponential backoff
err = exc.Retry(func() error {
	return httpClient.Do(req)
}, 3)
```
