#!/bin/bash
echo "Build..."
cd /var/lib/postgresql/backend/b2b-backend && git pull
docker-compose build
echo "Build completed"
