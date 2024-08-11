#!/bin/bash
echo "Build..."
cd /var/lib/postgresql/backend/b2b-backend
docker-compose build --parallel --no-cache
echo "Build completed"
