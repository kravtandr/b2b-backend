#!/bin/bash
echo "Start docker-compose"
cd /var/lib/postgresql/backend/b2b-backend
docker-compose up -d --remove-orphans