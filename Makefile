.PHONY: build install

build:
	go build ./kversion.go

install:
	chmod +x ./kversion && mv ./kversion /usr/local/bin/

uninstall:
	rm -rf /usr/local/bin/kversion

clean: 
	rm -rf ./kversion

