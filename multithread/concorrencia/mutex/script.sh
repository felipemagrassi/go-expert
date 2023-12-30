#!/bin/sh

echo "start"
ab -n 1000 -c 100 http://localhost:8080/ 
curl http://localhost:8080/
echo "end"
