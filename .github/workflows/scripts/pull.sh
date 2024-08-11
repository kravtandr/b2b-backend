#!/bin/bash
git config --global --add safe.directory /var/lib/postgresql/backend/b2b-backend
cd /var/lib/postgresql/backend/b2b-backend && git pull