package main

import (
	"encoding/hex"
	"flag"
	"ibmmq-client/mq"
	"io"
	"log"
	"os"
	"strings"

	"github.com/ibm-messaging/mq-golang/v5/ibmmq"
)

var logger = log.New(os.Stdout, "MQ Producer:", log.LstdFlags)

func main() {
	var file string
	flag.StringVar(&file, "f", "message.json", "Specify a file to sent")
	flag.Parse()

	f, err := os.Open(file)
	if err != nil {
		logger.Fatalln(err)
	}
	defer f.Close()

	logger.Println("== Application is starting ==")
	mq.MQSettings.LogSettings()

	// Read file data, ignore error right now
	data, _ := io.ReadAll(f)
	emitMessage(&data)
	logger.Println("== Application is complete ==")
}

func emitMessage(data *[]byte) {

	// Get a MQ Manager
	qMgr, err := mq.ConnectToQ(mq.FULL_STRING)
	if err != nil {
		logger.Fatalln(err)
	}
	defer qMgr.Disc()

	qObj, err := mq.OpenQueue(qMgr, mq.OP_Put)
	if err != nil {
		logger.Fatalln(err)
	}
	defer qObj.Close(0)

	// Message Descriptor (MQMD) and Put Options (MQPMO)
	pmd := ibmmq.NewMQMD()
	pmo := ibmmq.NewMQPMO()

	// Explicit about transactional boundaries as not all platforms behave the same way
	pmo.Options = ibmmq.MQPMO_NO_SYNCPOINT

	// Use text string as message boday format
	pmd.Format = ibmmq.MQFMT_STRING

	logger.Println("Put Message To:", strings.TrimSpace(qObj.Name))
	err = qObj.Put(pmd, pmo, *data)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("Put Message MsgId:", hex.EncodeToString(pmd.MsgId))

}
