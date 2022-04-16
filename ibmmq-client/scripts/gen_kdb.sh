#!/bin/sh

#
# Shell not tested, combine the command to run it is OK
#
CMD="/opt/mqm/bin/runmqakm"

KDB="clientkey.kdb"
CRT_FILE="key.crt"

PASSWORD="tru5passw00d" ## Change to your password
LABEL="ati.cert"

${CMD} -keydb -create -db ${KDB} -pw ${PASSWORD} -type pkcs12 -expire 1000 -stash
${CMD} -cert -add -label ${LABEL_NAME} -db ${KDB} -stashed -trust enable -file ${CRT_FILE}
