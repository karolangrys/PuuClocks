#!/bin/bash

while ! mysql -uroot -proot -h127.0.0.1 -P3306 -e 'SELECT 1' > /dev/null 2>&1; do
    echo Waiting for db ...
    sleep 2
done

while [ true ]
do
    data=`curl --write-out '%{http_code}' --silent --output /dev/null 127.0.0.1:8080`
    if [[ "$data" -eq 404 ]]; then
        exit 0
    fi
    echo Waiting for app ...
    #docker logs app-tcs
    sleep 2
done
