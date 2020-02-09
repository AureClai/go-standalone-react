ONEHSELL:
PHONY:

clean-and-statik: clean statik

statik:
	cd web && yarn build && cd ..
	statik -src=./web/build

clean:
	rm -rf ./statik
	rm -rf ./web/build

go-build:
	go build -o ./bin/macOSapp