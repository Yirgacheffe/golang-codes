# IBMMQ Client

This tool is for internal testing purpose, include 3 parts, rightnow please use the 'producer' and 'nisemono'.

## Prerequisites

You are going to need:
- **Go Installation**
- **IBM MQ**

## How to run it locally

- Setup your MQ endpoint in 'configs/mq.json', if SSL used, make sure 'clientkey.kdb' in right place.
- Run 'build.sh'.
- Run 'producer -f ./tests/greeting.json' for smoke test.
- Run 'nisemono' to mock vendor api.

- Run 'clean.sh' to remove all executable file.

NOTES: Please feel free to find the useful scripts in folder 'scripts'.
