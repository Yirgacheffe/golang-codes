package main

import (
	"encoding/json"
	"fmt"
	"ibmmq-client/mq"
	"ibmmq-client/types"
	"log"
	"os"
	"time"

	"github.com/ibm-messaging/mq-golang/v5/ibmmq"
)

var logger = log.New(os.Stdout, "MQ Consumer: ", log.LstdFlags)

var mh ibmmq.MQMessageHandle
var ok = true

const (
	exit_err = 1
)

// The main function just expects to be given a return code for Exit()
func main() {
	os.Exit(mainWithRc())
}

// The real main function is here to set a return code.
func mainWithRc() int {

	mq.MQSettings.LogSettings()

	// Connect to the queue manager
	qMgr, err := mq.ConnectToQ(mq.FULL_STRING)
	if err != nil {
		logger.Fatalln(err)
		return exit_err
	}
	defer qMgr.Disc()

	// Open the queue
	qObj, err := mq.OpenQueue(qMgr, mq.OP_Get)
	if err != nil {
		logger.Fatalln(err)
		return exit_err
	}
	defer qObj.Close(0)

	// Create message handler object
	cmho := ibmmq.NewMQCMHO()
	mh, err = qMgr.CrtMH(cmho)
	if err != nil {
		logger.Fatalln(err)
		return exit_err
	}
	defer dltMh(mh)

	// Message Descriptor (MQMD) and Get Options (MQPMO)
	gmd := ibmmq.NewMQMD()
	gmo := ibmmq.NewMQGMO()

	gmo.Options = ibmmq.MQGMO_NO_SYNCPOINT
	gmo.Options |= ibmmq.MQGMO_WAIT
	gmo.Options |= ibmmq.MQGMO_PROPERTIES_IN_HANDLE

	gmo.MsgHandle = mh

	// The MQCBD structure is used to specify the function to be invoked
	// when a message arrives on a queue
	cbd := ibmmq.NewMQCBD()
	cbd.CallbackFunction = cb

	err = qObj.CB(ibmmq.MQOP_REGISTER, cbd, gmd, gmo)
	if err != nil {
		logger.Fatalln(err)
		return exit_err
	}
	defer deReg(qObj, cbd, gmd, gmo)

	// Enable the callback function
	// Any messages on the Queue will be sent to the Callback
	ctlo := ibmmq.NewMQCTLO()
	err = qMgr.Ctl(ibmmq.MQOP_START, ctlo)
	if err != nil {
		logger.Fatalln(err)
		return exit_err
	}

	defer stopCB(qMgr)

	// Keep the program running until the callback has indicated there are
	// no more messages.
	d, _ := time.ParseDuration("5s")
	for ok && err == nil {
		time.Sleep(d)
	}

	// Checking MQ return
	mqRet := 0
	if err != nil {
		mqRet = int((err.(*ibmmq.MQReturn)).MQCC)
	}
	return mqRet
}

// Deallocate the message handle
func dltMh(mh ibmmq.MQMessageHandle) error {
	dmho := ibmmq.NewMQDMHO()
	err := mh.DltMH(dmho)
	if err != nil {
		logger.Println(err)
	} else {
		logger.Println("Closed a Message Handler")
	}
	return err
}

// Deregister the callback function - have to do this before the message handle can be
// successfully deleted
func deReg(qObj ibmmq.MQObject, cbd *ibmmq.MQCBD, gmd *ibmmq.MQMD, gmo *ibmmq.MQGMO) error {
	err := qObj.CB(ibmmq.MQOP_DEREGISTER, cbd, gmd, gmo)
	if err != nil {
		logger.Println(err)
	} else {
		logger.Println("Deregistered callback")
	}
	return err
}

// Stop the callback function from being called again
func stopCB(qMgr ibmmq.MQQueueManager) {
	ctlo := ibmmq.NewMQCTLO()
	err := qMgr.Ctl(ibmmq.MQOP_STOP, ctlo)
	if err != nil {
		logger.Println(err)
	} else {
		logger.Printf("Stopped callback function\n")
	}
}

// Callback function invoked when a message arrives on the queue
func cb(qMgr *ibmmq.MQQueueManager, qObj *ibmmq.MQObject, md *ibmmq.MQMD, gmo *ibmmq.MQGMO, buffer []byte, cbc *ibmmq.MQCBC, err *ibmmq.MQReturn) {
	buflen := len(buffer)

	if err.MQCC != ibmmq.MQCC_OK {
		logger.Println(err)
		ok = false
	} else {
		// Assume the message is a printable string
		logger.Printf("Got Message: Length %d from Queue %s\n", buflen, qObj.Name)

		var req types.Request
		err := json.Unmarshal(buffer, &req)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Printf("Entity received: %s, %s\n", req.ReqId, req.TimeStamp)
		}
		// Add message dealing logic here ... DB Operate, File Operate
	}
}
