#!/bin/bash
echo "Clean and rights"
cd /var/lib/postgresql/backend && sudo chown -R postgres b2b-backend/
cd /var/lib/postgresql/backend/b2b-backend && chmod -R 777 pgdata/
docker image prune -a -f