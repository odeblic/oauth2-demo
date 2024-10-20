KEYSTORE_LOCATION = src/main/resources/keystore.jks
KEYSTORE_PASSWORD = 123456

.PHONY: setup certificate check build run clean help

setup: certificate check build

certificate:
	keytool -keystore $(KEYSTORE_LOCATION) -genkeypair -alias my-key -keyalg RSA -storepass $(KEYSTORE_PASSWORD) -dname "CN=, OU=, O=Home, L=, ST=, C=HK"
	keytool -keystore $(KEYSTORE_LOCATION) -exportcert -alias my-key -rfc -file cert.pem -storepass $(KEYSTORE_PASSWORD)

check:
	mvn validate
	keytool -keystore $(KEYSTORE_LOCATION) -list -storepass $(KEYSTORE_PASSWORD)

build:
	mvn compile

run:
	mvn exec:java -Dexec.mainClass="com.nalfiro.Application"
	#mvn exec:java
	#mvn spring-boot:run

clean:
	mvn clean
	rm -f cert.pem
	rm -f $(KEYSTORE_LOCATION)

help:
	@printf "Makefile usage:\n"
	@printf "\n"
	@printf "\033[32m  setup       \033[0m invoke \033[32mcertificate\033[0m, \033[32mcheck\033[0m and \033[32mbuild\033[0m\n"
	@printf "\033[32m  certificate \033[0m generate the X.509 certificate\n"
	@printf "\033[32m  check       \033[0m verify the project configuration\n"
	@printf "\033[32m  build       \033[0m compile the program\n"
	@printf "\033[32m  run         \033[0m run the program\n"
	@printf "\033[32m  clean       \033[0m remove all build artifacts\n"
	@printf "\033[32m  help        \033[0m display this message\n"
	@printf "\n"
