package main

import (
	"encoding/hex"
	"flag"
	"ibmmq-client/utils"
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
		log.Fatalln(err)
		os.Exit(1)
	}

	defer f.Close()

	data, _ := io.ReadAll(f)

	logger.Println("== Application is starting ==")
	utils.EnvSettings.LogSettings()
	emitMessage(&data)
	logger.Println("== Application is complete ==")
}

func emitMessage(data *[]byte) {

	// Get a MQ Manager
	qMgr, err := utils.ConnectToQ(utils.FULL_STRING)
	if err != nil {
		logger.Println(err)
		logger.Fatalln("Unable connect to Server")
		os.Exit(1)
	}

	defer qMgr.Disc()

	qObj, err := utils.OpenQueue(qMgr, utils.OP_Put)
	if err != nil {
		log.Fatalln("Unable to open message queue")
		os.Exit(1)
	}

	defer qObj.Close(0)

	logger.Println("Writing message to Queue ...")

	// Message Descriptor (MQMD) and Put Options (MQPMO)
	// Create those with default values
	pmd := ibmmq.NewMQMD()
	pmo := ibmmq.NewMQPMO()

	// Explicit about transactional boundaries as not all platforms behave the same way
	pmo.Options = ibmmq.MQPMO_NO_SYNCPOINT

	// Use text string as message boday format
	pmd.Format = ibmmq.MQFMT_STRING

	logger.Printf("Sending message %s", *data)
	err = qObj.Put(pmd, pmo, *data)

	if err != nil {
		logger.Println(err)
	} else {
		logger.Println("Put Message To:", strings.TrimSpace(qObj.Name))
		logger.Println("MsgId:", hex.EncodeToString(pmd.MsgId))
	}

}
