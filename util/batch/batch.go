package batch

import "errors"

type BatchFunc func(start, end int) error

var ErrAbort = errors.New("done")

// Batch calls eachFn for all items
// Returns any error from eachFn except for AbortErr it returns nil.
func Batch(count, size int, eachFn BatchFunc) error {

	last := count - 1
	i := 0

	for i < count {
		end := i + size - 1
		if end > last {
			end = last
		}

		if err := eachFn(i, end); err != nil {
			if err == ErrAbort {
				return nil
			} else {
				return err
			}
		}
		i = end + 1 // move to next batch start
	}

	return nil
}
