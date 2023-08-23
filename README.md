#b2b-backend

run:           docker-compose up

Delite b2b-backend/pgdata folder before re-build for postgres schema reset!!!
re-build:                 docker-compose up --build

path where change origin: cmd/gateway/main.go
	local: ctx.Response.Header.Set("Access-Control-Allow-Origin", "http://localhost:3000")
	deploy: //ctx.Response.Header.Set("Access-Control-Allow-Origin", "https://bi-tu-bi.ru")
    after the change necessary re-build gateway service 

re-build changed origin:  docker-compose up -d --force-recreate --no-deps --build gateway

full re-build: docker system prune -a
               docker-compose up 


There is no test data in initial state of postgres!!!
How to add test data:
    Enter postgres container: docker exec -it postgres_container /bin/sh
    Enter psql in container: psql b2b b2b

    add test categories: 
    COPY categories(name) FROM '/var/lib/postgresql/data/export_base_categories.csv' DELIMITER ',' CSV HEADER;

    add test products: 
    COPY products(name, description, price, photo) FROM '/var/lib/postgresql/data/test_products.csv' DELIMITER ';' CSV HEADER;

Minio:
    http://127.0.0.1:9001/
    b2b
    b2b_pass