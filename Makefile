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
SYSL_GO=anzbank/sysl-go:v0.59.0
SYSL=anzbank/sysl:v0.152.0
.PHONY: setup gen downstream
all: setup gen downstream format

# try to clone, then try to fetch and pull
setup:
	docker run -v $$(pwd):/mount:ro $(SYSL) pb --mode json /mount/api/project.sysl > api/project.json
	$(foreach path, $(deps), $(shell mkdir -p ${outdir}/$(path)))
    $(foreach path, $(apps), $(shell mkdir -p ${outdir}/$(path)))
	

# Generate files with internal git service
gen:
	$(foreach app, $(apps), $(shell echo "docker run --rm -v $$(pwd):/mount:ro  $(SYSL_GO) /sysl-go/codegen/arrai/service.arrai $(basepath)/$(outdir) /mount/api/project.json $(app) rest-app | tar xf - -C $(outdir)/$(app)"))

downstream:
	$(foreach app, $(deps), $(shell docker run --rm -v $$(pwd):/mount:ro $(SYSL_GO) /sysl-go/codegen/arrai/service.arrai $(basepath)/$(outdir) /mount/api/project.json $(app) rest-client | tar xf - -C $(outdir)/$(app)))

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