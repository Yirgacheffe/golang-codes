#!/bin/sh

openssl req -newkey rsa:2048 -nodes -keyout key.key -x509 -days 365 -out key.crt
