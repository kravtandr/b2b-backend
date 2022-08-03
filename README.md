# b2b-backend

psql postgres
CREATE USER b2b WITH PASSWORD 'b2b';
create database b2b;
psql -U b2b -d b2b 

cd main
go run amin.go