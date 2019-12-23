GO_LINUX := GOOS=linux GOARCH=amd64

clean:
	-rm main
	-rm main.zip

build: clean
	$(GO_LINUX) go build main.go

package: build
	zip main.zip main

deploy-qa: package
	python3 ops/deploy.py qa

deploy-prod: package
	python3 ops/deploy.py prod

rollback-qa:
	python3 ops/rollback.py qa

rollback-prod:
	python3 ops/rollback.py prod


	