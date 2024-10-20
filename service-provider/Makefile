.PHONY: setup environment certificate run test-ko test-ok clean help

setup: environment certificate

environment:
	python3 -m venv environment
	. environment/bin/activate ; pip install pyjwt flask

certificate:
	openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -days 365 -nodes -subj "/C=HK/ST=/L=/O=Home/OU=/CN=localhost/emailAddress=myself@localhost"

run:
	. environment/bin/activate ; python server.py

test-ko:
	curl -k -H "Authorization: Bearer xxxxx" https://localhost:5003/resource

test-ok:
	curl -k -H "Authorization: Bearer $(shell python tokengen.py alice all | grep Token: | cut -d' ' -f2 | sed -e 's/\x1b\[[0-9]\+m//g')" https://localhost:5003/resource

clean:
	rm -f key.pem
	rm -f cert.pem
	rm -fr __pycache__
	rm -fr environment

help:
	@printf "Makefile usage:\n"
	@printf "\n"
	@printf "\033[32m  setup       \033[0m invoke \033[32menvironment\033[0m and \033[32mcertificate\033[0m\n"
	@printf "\033[32m  environment \033[0m create the virtual environment\n"
	@printf "\033[32m  certificate \033[0m generate the X.509 certificate\n"
	@printf "\033[32m  run         \033[0m run the program\n"
	@printf "\033[32m  test-ko     \033[0m access resource with failure\n"
	@printf "\033[32m  test-ok     \033[0m access resource with success\n"
	@printf "\033[32m  clean       \033[0m cleanup all artifacts\n"
	@printf "\033[32m  help        \033[0m display this message\n"
	@printf "\n"
