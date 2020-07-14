SYSL  =  ./syslw
ARRAI  =  ./arraiw
CODEGEN_ROOT = internal/gen
SPECS_ROOT = specs/frontend/petdemo

.PHONY: all
gen: backend-gen frontend-gen

# frontend code
.PHONY: frontend-gen
.PHONY: frontend
frontend-gen: petdemo-gen

.PHONY: Petdemo
PETDEMO = Petdemo

frontend = rest-service
.PHONY: petdemo-gen
petdemo-gen: $(PETDEMO) frontend
	$(run-codegen)

# backend code
.PHONY: backend-gen
.PHONY: backend
backend-gen: petstore-gen flickr-gen

.PHONY: Flickr
FLICKR = Flickr

.PHONY: flickr-gen
flickr-gen: $(FLICKR) backend
	$(run-codegen)

.PHONY: Petstore
PETSTORE = Petstore

.PHONY: petstore-gen
petstore-gen: $(PETSTORE) backend
	$(run-codegen)

SYSL_TEMPLATE_ROOT = github.com/anz-bank/sysl-template
backend = rest-service

define run-codegen
	rm -rf $(CODEGEN_ROOT)/$(shell echo $< | tr A-Z a-z)
	rm -rf $(SPECS_ROOT)/sysl.json
	mkdir -p $(CODEGEN_ROOT)/$(shell echo $< | tr A-Z a-z)
	$(SYSL) pb --mode json $(SPECS_ROOT)/petdemo.sysl >> $(SPECS_ROOT)/sysl.json
	$(ARRAI) run service_stub.arrai \
		$(SYSL_TEMPLATE_ROOT)/$(CODEGEN_ROOT) $(SPECS_ROOT)/sysl.json \
		$< $($(word 2,$^)) | tar xf - -C $(CODEGEN_ROOT)/$(shell echo $< | tr A-Z a-z)
	goimports -w $(CODEGEN_ROOT)/$(shell echo $< | tr A-Z a-z)
endef
