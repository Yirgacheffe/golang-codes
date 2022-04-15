package main

import (
	"encoding/hex"
	"ibmmq-client/utils"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ibm-messaging/mq-golang/v5/ibmmq"
)

var logger = log.New(os.Stdout, "MQ Listener: ", log.LstdFlags)

// The main function just call handleMessage to retrieve message
func main() {
	os.Exit(mainWithRc())
}

// The real main function is here to set a return code
func mainWithRc() int {

	utils.EnvSettings.LogSettings()

	// Connect to the queue manager
	qMgr, err := utils.ConnectToQ(utils.FULL_STRING)
	if err != nil {
		logger.Fatalln(err)
		return 1
	}
	defer qMgr.Disc()

	// Open queue
	qObj, err := utils.OpenQueue(qMgr, utils.OP_Get)
	if err != nil {
		logger.Fatalln(err)
		return 1
	}

	defer qObj.Close(0)

	msgAvail := true
	for msgAvail == true && err == nil {
		var datalen int

		gotMsg := false

		// Message Descriptor (MQMD) and Get Options (MQPMO)
		gmd := ibmmq.NewMQMD()
		gmo := ibmmq.NewMQGMO()

		gmo.Options = ibmmq.MQGMO_NO_SYNCPOINT
		gmo.Options |= ibmmq.MQGMO_WAIT
		gmo.Options |= ibmmq.MQGMO_ACCEPT_TRUNCATED_MSG

		var msgId string
		if msgId != "" {
			logger.Println("Setting Match Option for MsgId")
			gmo.MatchOptions = ibmmq.MQMO_MATCH_MSG_ID
			gmd.MsgId, _ = hex.DecodeString(msgId)

			// Will only try to get a single message with the MsgId as there should
			// never be more than one. So set the flag to not retry after the first attempt.
			msgAvail = false
		}

		buffer := make([]byte, 4096)

		datalen, err = qObj.Get(gmd, gmo, buffer)
		if err != nil {
			logger.Println(err)
			msgAvail = false
			mqret := err.(*ibmmq.MQReturn)
			if mqret.MQRC == ibmmq.MQRC_NO_MSG_AVAILABLE {
				err = nil
			}
		} else {
			logger.Printf("Got message of lenght %d: ", datalen)
			logger.Println(strings.TrimSpace(string(buffer[:datalen])))
			gotMsg = true
		}

		if gotMsg {
			t := gmd.PutDateTime
			if !t.IsZero() {
				diff := time.Now().Sub(t)
				round, _ := time.ParseDuration("1s")
				diff = diff.Round(round)

				logger.Printf("Message was put %d seconds ago\n", int(diff.Seconds()))
			} else {
				logger.Printf("Message has empty PutDateTime - MQMD PutDate: 's'\n", gmd.PutDateTime)
			}
		}
	}

	mqret := 0
	if err != nil {
		mqret = int((err.(*ibmmq.MQReturn)).MQCC)
	}

	return mqret
}
