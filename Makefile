APP_DIR = application-client
OAUTH_DIR = authorization-server
SERV_DIR = service-provider
KEYSTORE_PASSWORD = changeit

.PHONY: setup trust revoke run-app run-oauth run-serv clean help

setup:
	$(MAKE) -C $(APP_DIR)   setup
	$(MAKE) -C $(OAUTH_DIR) setup
	$(MAKE) -C $(SERV_DIR)  setup

trust:
	sudo keytool -cacerts -importcert -noprompt -file $(APP_DIR)/cert.pem   -alias "demo:application-client-certificate"   -storepass $(KEYSTORE_PASSWORD)
	sudo keytool -cacerts -importcert -noprompt -file $(SERV_DIR)/cert.pem  -alias "demo:service-provider-certificate"     -storepass $(KEYSTORE_PASSWORD)
	sudo keytool -cacerts -importcert -noprompt -file $(OAUTH_DIR)/cert.pem -alias "demo:authorization-server-certificate" -storepass $(KEYSTORE_PASSWORD)

revoke:
	sudo keytool -cacerts -delete -alias "demo:application-client-certificate"   -storepass $(KEYSTORE_PASSWORD)
	sudo keytool -cacerts -delete -alias "demo:service-provider-certificate"     -storepass $(KEYSTORE_PASSWORD)
	sudo keytool -cacerts -delete -alias "demo:authorization-server-certificate" -storepass $(KEYSTORE_PASSWORD)

run-app:
	$(MAKE) -C $(APP_DIR) run

run-oauth:
	$(MAKE) -C $(OAUTH_DIR) run

run-serv:
	$(MAKE) -C $(SERV_DIR) run

clean:
	$(MAKE) -C $(APP_DIR)   clean
	$(MAKE) -C $(OAUTH_DIR) clean
	$(MAKE) -C $(SERV_DIR)  clean

help:
	@printf "Makefile usage:\n"
	@printf "\n"
	@printf "\033[32m  build     \033[0m compile the program\n"
	@printf "\033[32m  trust     \033[0m trust our self-signed certificates (need \033[33mroot\033[0m)\n"
	@printf "\033[32m  revoke    \033[0m revoke our self-signed certificates (need \033[33mroot\033[0m)\n"
	@printf "\033[32m  run-app   \033[0m run the \033[35mapplication client\033[0m\n"
	@printf "\033[32m  run-oauth \033[0m run the \033[35mauthorization server\033[0m\n"
	@printf "\033[32m  run-serv  \033[0m run the \033[35mservice provider\033[0m\n"
	@printf "\033[32m  clean     \033[0m remove all artifacts\n"
	@printf "\033[32m  help      \033[0m display this message\n"
	@printf "\n"
