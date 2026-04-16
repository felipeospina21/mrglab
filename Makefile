start:
	go run .

dev:
	go run . -dev

build:
	go build .

test:
	go test ./...

cover: 
	go test ./... -coverprofile=c.out
	go tool cover -html="c.out"

release:
ifndef v
	$(error Usage: make release v=0.2.0)
endif
	@echo "Releasing v$(v)..."
	git add -A
	git commit -m "chore release v$(v)" --allow-empty
	git tag v$(v)
	git push
	git push origin v$(v)
	@echo "Done. v$(v) released."

update-tuishell:
ifndef v
	$(error Usage: make update-tuishell v=0.3.0)
endif
	GOWORK=off go get github.com/felipeospina21/tuishell@v$(v)
	GOWORK=off go mod tidy
	@echo "Updated tuishell to v$(v)"