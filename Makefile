GOPATH := $(shell go env GOPATH)
GOBIN := $(PWD)/bin

PATH := $(GOBIN):$(PATH)
SHELL := $(shell which bash)

.PHONY: deps
deps:
	go install github.com/xo/xo@latest
	brew install sqlcipher

.PHONY: genmodels
genmodels:
	@mkdir -p rekordbox/models
	@cp $(shell go run cmd/getdbpath/main.go) /tmp/master.db
	@sqlcipher "file:/tmp/master.db" "PRAGMA KEY = '$(shell go run cmd/getencryptionkey/main.go)'; ATTACH DATABASE 'plaintext.db' AS plaintext KEY '';SELECT sqlcipher_export('plaintext');" || true
	@xo schema "file:plaintext.db" -o "rekordbox/" --src tpl/
	@sqlite3 "file:plaintext.db" ".schema" > db/schema.sql
	@rm plaintext.db
	@echo "generated rekordbox/"