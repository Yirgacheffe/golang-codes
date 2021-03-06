package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime/debug"
)

const (
	jobBinPath = "/bad/job/binary"
	binPath    = "/bab/job/binary"
)

// Define a new type of MyError ...... demo for error handle
type MyError struct {
	Inner      error
	Message    string
	StackTrace string
	Misc       map[string]interface{}
}

func (err MyError) Error() string {
	return err.Message
}

func wrapError(err error, messagef string, msgArgs ...interface{}) MyError {
	return MyError{
		Inner:      err,
		Message:    fmt.Sprintf(messagef, msgArgs...),
		StackTrace: string(debug.Stack()),
		Misc:       make(map[string]interface{}),
	}
}

// Lowlevel Module
type LowLevelErr struct {
	error
}

func isGloballyExec(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, LowLevelErr{wrapError(err, err.Error())}
	}
	return info.Mode()&0100 == 0100, nil
}

// Intermediate Module
type IntermediateErr struct {
	error
}

func runJob(id string) error {
	isExecutable, err := isGloballyExec(jobBinPath)
	if err != nil {
		return err
	}
	if !isExecutable {
		return wrapError(nil, "job binary is not executable")
	}

	return exec.Command(jobBinPath, "--id="+id).Run()
}

// Use this when client don't care the low level error
func runJobAndHideLowlevelError(id string) error {
	isExecutable, err := isGloballyExec(binPath)
	if err != nil {
		return IntermediateErr{wrapError(
			err,
			"cannot run job %q: requisite binaries not available",
			id,
		)}
	}
	if !isExecutable {
		return wrapError(
			err,
			"cannot run job %q: requisite binaries not available",
			id,
		)
	}

	return exec.Command(binPath, "--id="+id).Run()
}

// Func main for demo
func handleError(key int, err error, message string) {
	log.SetPrefix(fmt.Sprintf("[logID: %v]:", key))
	log.Printf("%#v\n", err)
	fmt.Printf("[%v] %v", key, message)
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	const id = 1
	err := runJob(fmt.Sprint(id))

	if err != nil {
		msg := "There was an unexpected issue; please report this as a bug."
		if _, ok := err.(IntermediateErr); ok {
			msg = err.Error()
		}
		handleError(id, err, msg)
	}
}
