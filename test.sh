go test -v -coverpkg=./... -coverprofile=cover.out.tmp ./...
cat cover.out.tmp | grep -v "_test.go" | grep -v "main.go" | grep -v "_mock.go" | grep -v "mock_" | grep -v ".pb.go" | grep -v ".pb.validate.go" | grep -v "_easyjson.go" | grep -v "logging.go" | grep -v "config.go" | grep -v "setup.go" | grep -v "router.go"  > cover.out
go tool cover -func cover.out
cd internal/services/auth/delivery && go test -timeout 30s
