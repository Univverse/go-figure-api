#!/bin/sh

if [ $DEVELOPMENT == true ]; then
  cd /go/src/app

  export APP_PORT=8081

  touch /var/healthy

  gin run main.go
else
  touch /var/healthy

  supervisord -c '/etc/supervisor.d/supervisord.ini'
fi