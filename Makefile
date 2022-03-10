.SILENT :
.PHONY: docs


GO=$(shell which go)
GOMOD=$(shell echo "$${GO111MODULE:-auto}")


MDs=apis.md \
	chat_info.md \
	dept_info.md \
	external_contact.md \
	media_upload.md \
	oa.md \
	rx_msg.md \
	user_info.md

help:
	echo "make errcodegen | sdkcodegen"

docs:
	$(info docs: $(MDs))

generate:
	GO111MODULE=$(GOMOD) $(GO) generate ./...

errcodegen:
	GO111MODULE=$(GOMOD) $(GO) run -tags=sdkcodegen ./internal/errcodegen errcodes/mod.go

sdkcodegen:
	for name in $(MDs); do \
		echo $${name}; \
		GO111MODULE=$(GOMOD) $(GO) run -tags=sdkcodegen ./internal/sdkcodegen docs/$${name} $${name}.go ; \
	done
