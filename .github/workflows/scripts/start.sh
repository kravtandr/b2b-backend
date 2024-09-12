#!/bin/bash
echo "Start docker-compose"
cd /root/b2b-backend
docker compose up -d --remove-orphans