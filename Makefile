.PHONY: run
run:
	npx nodemon --delay 2.5 -e html,go --exec go run main.go --signal SIGTERM
.PHONY: test
test:
	go test
.PHONY: dcup
dcup:
	docker compose up -d
.PHONY: dcdown
dcdown:
	docker compose down