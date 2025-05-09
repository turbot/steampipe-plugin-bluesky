.PHONY: build install clean restart test

STEAMPIPE_INSTALL_DIR ?= ~/.steampipe
BUILD_TAGS = netgo
install:
	go build -o $(STEAMPIPE_INSTALL_DIR)/plugins/hub.steampipe.io/plugins/turbot/bluesky@latest/steampipe-plugin-bluesky.plugin -tags "${BUILD_TAGS}" *.go

localbuild:
	go build -o build/steampipe-plugin-bluesky.plugin .

localinstall: localbuild
	mkdir -p ~/.steampipe/plugins/local/bluesky
	cp build/steampipe-plugin-bluesky.plugin ~/.steampipe/plugins/local/bluesky/

clean:
	rm -f build/steampipe-plugin-bluesky.plugin
	go clean -cache

restart: install
	steampipe service restart

test:
	./tests/run_tests.sh

