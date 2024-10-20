# Service Provider

Create the virtual environment:

```sh
python -m venv environment
. environment/bin/activate
pip install pyjwt flask
```

Generate the X.509 certificate:

```sh
openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -days 365 -nodes
```

Run the program:

```sh
. environment/bin/activate
python server.py
```

Test resource access:

```sh
. environment/bin/activate

BAD_TOKEN="xxxxx"
curl -k -H "Authorization: Bearer ${BAD_TOKEN}" https://localhost:5003/resource

GOOD_TOKEN=$(python tokengen.py alice all | grep Token: | cut -d' ' -f2 | sed -e 's/\x1b\[[0-9]\+m//g')
curl -k -H "Authorization: Bearer ${GOOD_TOKEN}" https://localhost:5003/resource
```
