server:
	go run ./cmd/api/server/main.go

client:
	go run ./cmd/api/client/main.go

generateAuth:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./internal/proto/auth/auth.proto

generateUpload:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./internal/proto/upload/upload.proto