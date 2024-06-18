swag_init:
	swag init -g api/router.go -o api/docs

run:
	go run cmd/main.go

git-push:
	git push origin main