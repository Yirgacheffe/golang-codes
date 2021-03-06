package mq

import (
	"github.com/ibm-messaging/mq-golang/v5/ibmmq"
	"github.com/pkg/errors"
)

const (
	OP_Put = "PUT"
	OP_Get = "GET"
)

func ConnectToQ(index int) (ibmmq.MQQueueManager, error) {
	logger.Println("Setting up Connection to MQ")

	// Allocate the MQCNO structure needed for the CONNX call
	env := getEndpoint(index)
	cno := ibmmq.NewMQCNO()

	if user := env.User; user != "" {
		logger.Printf("User %s has been specified\n", user)
		csp := ibmmq.NewMQCSP()
		csp.AuthenticationType = ibmmq.MQCSP_AUTH_USER_ID_AND_PWD
		csp.UserId = user
		csp.Password = env.Password

		cno.SecurityParms = csp // Refer to CSP structure
	}

	// Fill in required fields in the MQCD channel definition structure
	cd := ibmmq.NewMQCD()
	cd.ConnectionName = env.GetConnection(index)
	cd.ChannelName = env.Channel

	logger.Printf("Connect to %s", cd.ConnectionName)

	// Well, set cipher to CD, if no key set, the connect will be failed
	if env.KeyRepo != "" {
		logger.Println("Client is Running in TLS Mode")
		cd.SSLCipherSpec = env.Cipher
		cd.SSLClientAuth = ibmmq.MQSCA_OPTIONAL
	}

	cno.ClientConn = cd

	// Specify the key repository location if it has been provided
	// in the environment json setting
	if env.KeyRepo != "" {
		logger.Println("Key repository has been specified")
		sco := ibmmq.NewMQSCO()
		sco.KeyRepository = env.KeyRepo

		cno.SSLConfig = sco
	}

	cno.Options = ibmmq.MQCNO_CLIENT_BINDING
	logger.Printf("Attempting connect to %s", env.QMgr)

	qMgr, err := ibmmq.Connx(env.QMgr, cno)
	if err != nil {
		return qMgr, errors.Wrap(err, "Connect failed")
	}

	return qMgr, nil // --------------- Succeed!!! ----------------
}

func OpenQueue(qMgr ibmmq.MQQueueManager, opType string) (ibmmq.MQObject, error) {
	return openQ(qMgr, opType, FULL_STRING)
}

func getEndpoint(index int) MQEnv {
	if index == FULL_STRING {
		index = 0
	}

	return Q_EPs.Points[index]
}

func openQ(qMgr ibmmq.MQQueueManager, opType string, idx int) (ibmmq.MQObject, error) {
	// Object Descriptor allows us to set q name
	mqod := ibmmq.NewMQOD()
	qEnv := getEndpoint(idx)

	// Operate depends on 'openOptions' parameter
	var opts int32
	var qObj ibmmq.MQObject

	switch opType {
	case OP_Get:
		opts = ibmmq.MQOO_INPUT_EXCLUSIVE
	case OP_Put:
		opts = ibmmq.MQOO_OUTPUT
	default:
		return qObj, errors.New("Unknown op type, [Put] or [Get]")
	}

	mqod.ObjectType = ibmmq.MQOT_Q
	mqod.ObjectName = qEnv.QName

	logger.Println("Attempting open queue", qEnv.QName)

	qObj, err := qMgr.Open(mqod, opts)
	if err != nil {
		return qObj, errors.Wrap(err, "Unable to open queue "+qEnv.QName)
	}

	logger.Printf("Queue [%s] opened successful", qObj.Name)
	return qObj, nil
}
