input = api/project.sysl
apps = simple
deps = jsonplaceholder # this can be a list separated by a space or left empty
outdir = gen
# Current go import path
basepath = github.com/anz-bank/sysl-template

####################################################################
#                                                                  #
#                                                                  #
#                                                                  #
# START SYSL MAKEFILE: you shouldn't need to edit anything below   #
#                                                                  #
#                                                                  #
#                                                                  #
####################################################################

SERVERLIB = /var/tmp/sysl-go
.PHONY: setup gen downstream
all: setup gen downstream format

# try to clone, then try to fetch and pull
setup:
	# Syncing server-lib to $(SERVERLIB)
	git clone --depth 1 --branch joshcarp/arrai-fix https://github.com/anz-bank/sysl-go/ $(SERVERLIB) || true  # Don't fail
	cd $(SERVERLIB) && git pull
	$(foreach path, $(deps), $(shell mkdir -p ${outdir}/$(path)))
    $(foreach path, $(apps), $(shell mkdir -p ${outdir}/$(path)))

# Generate files with internal git service
gen:
	$(foreach app, $(apps), $(shell $(SERVERLIB)/codegen/arrai/service.arrai $(basepath)/$(outdir) api/project.json $(app) rest-app | tar xf - -C $(outdir)/$(app)))

downstream:
	$(foreach app, $(deps), $(shell $(SERVERLIB)/codegen/arrai/service.arrai $(basepath)/$(outdir) api/project.json $(app) rest-client | tar xf - -C $(outdir)/$(app)))
.PHONY: format
format:
	$(foreach path, $(deps), $(shell gofmt -s -w ${outdir}/${path}/))
	$(foreach path, $(deps), $(shell goimports -w ${outdir}/${path}/))
	$(foreach path, $(apps), $(shell gofmt -s -w ${outdir}/${path}/))
	$(foreach path, $(apps), $(shell goimports -w ${outdir}/${path}/))

docker:
	GOOS=linux GOARCH=amd64 go build main.go
	docker build -t joshcarp/sysltemplate .
	docker run -p 8080:8080 joshcarp/sysltemplate
