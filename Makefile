# Copyright (C) 2019 Steven Stallion <sstallion@gmail.com>
#
# Redistribution and use in source and binary forms, with or without
# modification, are permitted provided that the following conditions
# are met:
# 1. Redistributions of source code must retain the above copyright
#    notice, this list of conditions and the following disclaimer.
# 2. Redistributions in binary form must reproduce the above copyright
#    notice, this list of conditions and the following disclaimer in the
#    documentation and/or other materials provided with the distribution.
#
# THIS SOFTWARE IS PROVIDED BY THE AUTHOR AND CONTRIBUTORS ``AS IS'' AND
# ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
# IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
# ARE DISCLAIMED.  IN NO EVENT SHALL THE AUTHOR OR CONTRIBUTORS BE LIABLE
# FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
# DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
# OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
# HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
# LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
# OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
# SUCH DAMAGE.

GO = go
GOLINT = golint

GOOS = $(shell $(GO) env GOOS)

HIDAPI_DIR = $(PWD)/vendor/github.com/signal11/hidapi

HIDAPI_OS = $(if $(filter darwin,$(GOOS)),mac,$(GOOS))
HIDAPI_OS_DIR = $(HIDAPI_DIR)/$(HIDAPI_OS)

export CGO_CPPFLAGS = -I$(HIDAPI_DIR)
export CGO_LDFLAGS = -L$(HIDAPI_OS_DIR)

.PHONY: all
all:	cgo dep check

.PHONY: clean
clean:
	$(MAKE) -C $(HIDAPI_OS_DIR) -f Makefile.cgo clean
	$(GO) clean ./...

.PHONY: cgo
cgo:
	$(MAKE) -C $(HIDAPI_OS_DIR) -f Makefile.cgo all

.PHONY: dep
dep:
	$(GO) get -d ./...

.PHONY: check
check:
	$(GO) test -c # no tests to run

.PHONY: lint
lint:
	$(GO) vet ./...
	$(GOLINT) ./...

.PHONY: install
install:
	$(GO) install ./...

.PHONY: uninstall
uninstall:
	$(GO) clean -i ./...
