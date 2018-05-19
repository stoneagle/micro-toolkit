.PHONY: run-web, stop-web rm-web

PWD := $(shell pwd)
USER := $(shell id -u)
USERNAME := $(shell id -u -n)
GROUP := $(shell id -g)
PROJECT := toolkit 

run-web: 
	cd hack && docker-compose -p "$(PROJECT)-web-$(USER)" up
stop-web: 
	cd hack && docker-compose -p "$(PROJECT)-web-$(USER)" stop 
rm-web: 
	cd hack && docker-compose -p "$(PROJECT)-web-$(USER)" rm 
