#!/bin/bash
echo "waiting for MySQL to start......"

while ! nc -z easycreater-mysql 3306; do 
  sleep 1
done

echo "MySQL started. Launching application..."
./main
