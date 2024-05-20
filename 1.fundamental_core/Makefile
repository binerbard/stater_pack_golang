lowercase := $(shell echo $(title) | tr '[:upper:]' '[:lower:]')
generate:
	@mkdir -p tutorial-$(number)-$(title)
	@cd tutorial-$(number)-$(title) && go mod init $(lowercase) && touch main.go
