#!/bin/bash

curl --parallel --parallel-immediate -X POST -H "Content-Type: application/json" -d '{"amount":30,"success":true}' --config urls.txt