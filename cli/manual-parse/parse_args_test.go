package main

import "testing"

type testConfig struct {
	args []string
	err  error
	config
}

func TestParseArgs(t *testing.T) {

	tests := []testConfig{
		{
			args:   []string{"-h"},
			err:    nil,
			config: config{printUsage: true, nbrOfTimes: 0},
		},
		{
			args:   []string{"10"},
			err:    nil,
			config: config{printUsage: false, nbrOfTimes: 10},
		},
	}

	for _, tc := range tests {
		c, err := parseArgs(tc.args)
		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Fatalf("Expected error to be: %v, got: %v\n", tc.err, err)
		}
		if tc.err == nil && err != nil {
			t.Errorf("Expected nil error, got: %v\n", err)
		}
		if c.printUsage != tc.printUsage {
			t.Errorf("Expected printUsage to be: %v, got: %v\n", tc.printUsage, c.printUsage)
		}
		if c.nbrOfTimes != tc.nbrOfTimes {
			t.Errorf("Expected nbrOfTimes to be: %v, got: %v\n", tc.nbrOfTimes, c.nbrOfTimes)
		}
	}
}
