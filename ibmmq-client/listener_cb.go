package main

import (
	"ibmmq-client/utils"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ibm-messaging/mq-golang/v5/ibmmq"
)

var logger = log.New(os.Stdout, "MQ Listener: ", log.LstdFlags)

var mh ibmmq.MQMessageHandle
var ok = true

// The main function just expects to be given a return code for Exit()
func main() {
	os.Exit(mainWithRc())
}

// The real main function is here to set a return code.
func mainWithRc() int {

	utils.EnvSettings.LogSettings()

	// Connect to the queue manager
	qMgr, err := utils.ConnectToQ(utils.FULL_STRING)
	if err != nil {
		logger.Fatalln(err)
		return 1
	}
	defer qMgr.Disc()

	cmho := ibmmq.NewMQCMHO()
	mh, err = qMgr.CrtMH(cmho)
	if err != nil {
		logger.Fatalln(err)
		return 1
	}

	defer dltMh(mh)

	// Open the queue
	qObj, err := utils.OpenQueue(qMgr, utils.OP_Get)
	if err != nil {
		logger.Fatalln(err)
		return 1
	}

	defer qObj.Close(0)

	// Message Descriptor (MQMD) and Get Options (MQPMO)
	gmd := ibmmq.NewMQMD()
	gmo := ibmmq.NewMQGMO()

	gmo.Options = ibmmq.MQGMO_NO_SYNCPOINT
	gmo.Options |= ibmmq.MQGMO_WAIT
	// gmo.WaitInterval = 3 * 1000

	gmo.Options |= ibmmq.MQGMO_PROPERTIES_IN_HANDLE
	gmo.MsgHandle = mh

	// The MQCBD structure is used to specify the function to be invoked
	// when a message arrives on a queue
	cbd := ibmmq.NewMQCBD()
	cbd.CallbackFunction = cb

	err = qObj.CB(ibmmq.MQOP_REGISTER, cbd, gmd, gmo)
	if err != nil {
		logger.Fatalln(err)
		return 1
	}

	defer dereg(qObj, cbd, gmd, gmo)

	// Then we are ready to enable the callback function. Any messages
	// on the queue will be sent to the callback
	ctlo := ibmmq.NewMQCTLO()
	err = qMgr.Ctl(ibmmq.MQOP_START, ctlo)
	if err != nil {
		logger.Fatalln(err)
		os.Exit(1)
	}

	defer stopCB(qMgr)

	d, _ := time.ParseDuration("5s")
	for ok && err == nil {
		time.Sleep(d)
	}

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
func dereg(qObj ibmmq.MQObject, cbd *ibmmq.MQCBD, gmd *ibmmq.MQMD, gmo *ibmmq.MQGMO) error {
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

// This is the callback function invoked when a message arrives on the queue.
func cb(hConn *ibmmq.MQQueueManager, hObj *ibmmq.MQObject, md *ibmmq.MQMD, gmo *ibmmq.MQGMO, buffer []byte, cbc *ibmmq.MQCBC, err *ibmmq.MQReturn) {
	buflen := len(buffer)

	if err.MQCC != ibmmq.MQCC_OK {
		logger.Println(err)
		ok = false
	} else {
		// Assume the message is a printable string, which it will be
		// if it's been created by the amqsput program
		logger.Printf("In callback - Got message of length %d from queue %s: ", buflen, hObj.Name)
		logger.Println(strings.TrimSpace(string(buffer[:buflen])))
	}
}
