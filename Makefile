COVERAGE_DIR ?= .coverage

# cp from: https://github.com/yyle88/osexistpath/blob/09939b8aa19005f8f8f4a936235347883c6375bd/Makefile#L4
test:
	@-rm -r $(COVERAGE_DIR)
	@mkdir $(COVERAGE_DIR)
	make test-with-flags TEST_FLAGS='-v -race -covermode atomic -coverprofile $$(COVERAGE_DIR)/combined.txt -bench=. -benchmem -timeout 20m'

test-with-flags:
	@go test $(TEST_FLAGS) ./...
