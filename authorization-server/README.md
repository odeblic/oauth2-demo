# Authorization Server

Generate the X.509 certificate:

```sh
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes
```

Run the program:

```sh
go run .
```

Test authorization grant:

```sh
RESPONSE_TYPE=code
CLIENT_ID=app-0
REDIRECT_URI="https%3A%2F%2Flocalhost%3A5001%2Fcallback"
SCOPE=all
STATE=123456

curl -k "https://localhost:5002/authorize?response_type=${RESPONSE_TYPE}&client_id=${CLIENT_ID}&redirect_uri=${REDIRECT_URI}&scope=${SCOPE}&state=${STATE}"

GRANT_TYPE=authorization_code
AUTHORIZATION_CODE=e192f6
CLIENT_SECRET=secret-000

curl -k -X POST -H 'Content-Type: application/json' "https://localhost:5002/token?grant_type=${GRANT_TYPE}&code=${AUTHORIZATION_CODE}&client_id=${CLIENT_ID}&client_secret=${CLIENT_SECRET}"
```
