package batch

import "errors"

type BatchFunc func(start, end int) error

var AbortErr = errors.New("done")

// Batch calls eachFn for all items
// Returns any error from eachFn except for AbortErr it returns nil.
func Batch(count, size int, eachFn BatchFunc) error {

}
