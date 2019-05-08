all: generate

generate: validate clean
	swagger generate client -f swagger.yaml --name unifi --target ./pkg

clean:
	rm -rf pkg/*

validate:
	swagger validate ./swagger.yaml