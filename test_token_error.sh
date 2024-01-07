#!/bin/bash

TOKEN="any"
URL="http://localhost:8080/"
CONCURRENT_REQUESTS=100
TOTAL_REQUESTS=30

for ((i=1; i<=$TOTAL_REQUESTS; i++)); do
    TIMESTAMP=$(date +"%H:%M:%S")

    STATUS_CODE=$(curl -s -o /dev/null -w "%{http_code}\n" -H "API_KEY: $TOKEN" "$URL")
    
    if [ $STATUS_CODE -eq 000 ]; then
        echo "$TIMESTAMP - Erro Request $i: server is down"
    else
       echo "$TIMESTAMP - Request $i - Status Code: $STATUS_CODE"
    fi
    
    if [ $((i % $CONCURRENT_REQUESTS)) -eq 0 ]; then
        wait
    fi
done

wait
