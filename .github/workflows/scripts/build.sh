#!/bin/bash
echo "Build..."
cd /root/b2b-backend
docker-compose build --parallel --no-cache
echo "Build completed"
