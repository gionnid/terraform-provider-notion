BASE_DIR = ~/.terraform.d/plugins/terraform.local/local/notion
VERSION := $(shell head -n 1 VERSION.in)
PLATFORMS = linux_amd64

develop:
	mkdir -p ${BASE_DIR}/${VERSION}/${PLATFORMS}
	rm -rf example/.terraform
	rm -f example/.terraform.lock.hcl
	go build -o terraform-provider-notion
	mv terraform-provider-notion ${BASE_DIR}/${VERSION}/${PLATFORMS}

test-plan:
	make develop
	cd example ; \
	rm -rf .terraform* ; \
	terraform init ; \
	terraform plan

test-apply-clean:
	make develop
	cd example ; \
	rm -rf terraform.tfstate ; \
	terraform init ; \
	terraform apply --auto-approve

test-apply:
	make develop
	cd example ; \
	rm -rf .terraform* ; \
	terraform init ; \
	terraform apply --auto-approve

resources:
	cd example ; \
	terraform providers schema -json | \
		jq '.provider_schemas."terraform.local/local/notion".resource_schemas | keys'