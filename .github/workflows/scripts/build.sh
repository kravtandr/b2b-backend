#!/bin/bash
echo "Build..."
cd /var/lib/postgresql/backend/b2b-backend && git pull
docker-compose build --parallel --no-cache --env-file=.env
echo "Build completed"
