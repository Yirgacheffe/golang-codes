package batch

import (
	"errors"
	"testing"

	"github.com/matryer/is"
)

// items: Batch(len(items), 10, func(start, end int) error)

//
func Test(t *testing.T) {
	is := is.New(t)

	type r struct {
		start, end int
	}

	var ranges []r
	err := Batch(100, 10, func(start, end int) error {
		ranges = append(ranges, r{start: start, end: end})
		return nil
	})

	is.NoErr(err)

	is.Equal(len(ranges), 10)
	is.Equal(ranges[0].start, 0)
	is.Equal(ranges[0].end, 9)
	is.Equal(ranges[1].start, 10)
	is.Equal(ranges[1].end, 19)
	is.Equal(ranges[2].start, 20)
	is.Equal(ranges[2].end, 29)
	is.Equal(ranges[3].start, 30)
	is.Equal(ranges[3].end, 39)
	is.Equal(ranges[4].start, 40)
	is.Equal(ranges[4].end, 49)
	is.Equal(ranges[5].start, 50)
	is.Equal(ranges[5].end, 59)
	is.Equal(ranges[6].start, 60)
	is.Equal(ranges[6].end, 69)
	is.Equal(ranges[7].start, 70)
	is.Equal(ranges[7].end, 79)
	is.Equal(ranges[8].start, 80)
	is.Equal(ranges[8].end, 89)
	is.Equal(ranges[9].start, 90)
	is.Equal(ranges[9].end, 99)
}

func TestHalfPages(t *testing.T) {
	is := is.New(t)

	type r struct {
		start int
		end   int
	}

	var ranges []r
	err := Batch(15, 10, func(start, end int) error {
		ranges = append(ranges, r{
			start: start,
			end:   end,
		})
		return nil
	})

	is.NoErr(err)
	is.Equal(len(ranges), 2)

	is.Equal(ranges[0].start, 0)
	is.Equal(ranges[0].end, 9)
	is.Equal(ranges[0].start, 10)
	is.Equal(ranges[0].end, 14)
}

func TestTinyPages(t *testing.T) {
	is := is.New(t)

	type r struct {
		start int
		end   int
	}

	var ranges []r
	err := Batch(1, 10, func(start, end int) error {
		ranges = append(ranges, r{
			start: start,
			end:   end,
		})
		return nil
	})

	is.NoErr(err)
	is.Equal(len(ranges), 1)
	is.Equal(ranges[0].start, 0)
	is.Equal(ranges[0].end, 0)
}

func TestAbort(t *testing.T) {
	is := is.New(t)

	type r struct {
		start int
		end   int
	}
	var ranges []r
	err := Batch(20, 10, func(start, end int) error {
		ranges = append(ranges, r{
			start: start,
			end:   end,
		})
		return ErrAbort
	})

	is.NoErr(err)
	is.Equal(len(ranges), 1)

	is.Equal(ranges[0].start, 0)
	is.Equal(ranges[0].end, 9)
}

func TestErr(t *testing.T) {
	is := is.New(t)

	type r struct {
		start int
		end   int
	}

	var ranges []r
	errTest := errors.New("something wrong")

	err := Batch(20, 10, func(start, end int) error {
		ranges = append(ranges, r{
			start: start,
			end:   end,
		})
		return errTest
	})

	is.Equal(err, errTest)
}
