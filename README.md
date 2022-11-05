b2b-backend

run:           docker-compose up

re-build:      docker-compose up --build
               docker-compose up -d --force-recreate --no-deps --build company_service

full re-build: docker system prune -a
               docker-compose up 
