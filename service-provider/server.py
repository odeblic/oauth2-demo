from flask import Flask, request, jsonify
import jwt
import os


app = Flask(__name__)
app.config['SECRET_KEY'] = os.environ.get('SECRET_KEY', '0123456789')


resources = {
    'alice': 'I have a secret boyfriend.',
    'bob': 'I am actually gay.',
    'charlie': 'I smoke weed.',
    'denis': 'I hacked the school\'s server.',
    'emily': 'I have a crush on Charlie.',
    'felix': 'I believe in aliens.',
}


def token_required(f):
    def decorator(*args, **kwargs):
        print(f'Authorization: \033[36m{request.headers["Authorization"]}\033[0m')
        token = None
        if 'Authorization' in request.headers:
            token = request.headers['Authorization'].split(" ")[1]
        if not token:
            return jsonify({'message': 'Token is missing!'}), 401
        try:
            data = jwt.decode(token, app.secret_key, algorithms=["HS256"])
            cliend_id = data.get('client_id', '')
            scope = data.get('scope', '')
        except jwt.ExpiredSignatureError:
            return jsonify({'message': 'Token has expired!'}), 401
        except jwt.InvalidTokenError:
            return jsonify({'message': 'Token is invalid!'}), 401

        return f(cliend_id, scope, *args, **kwargs)
    return decorator


def build_message(user: str | None = None):
    message = ""
    if user is None:
        for user, secret in resources.items():
            message += f'The secret of user \"{user}\" is: \"{secret}\"\n'
    else:
        secret = resources[user]
        message = f'The secret of user \"{user}\" is: \"{secret}\"\n'
    return message


@app.route('/resource', methods=['GET'])
@token_required
def get_resource(cliend_id, scope):
    print(f'Request from client \033[32m<{cliend_id}>\033[0m with scope \033[34m<{scope}>\033[0m.')
    if scope == "all":
        message = build_message()
    elif scope in resources.keys():
        message = build_message(scope)
    else:
        return jsonify({'message': 'Invalid scope!'}), 401
    print(f'Response is "\033[33m{message}\033[0m".')
    return jsonify({'message': message})


if __name__ == '__main__':
    app.run(ssl_context=('cert.pem', 'key.pem'), port=5003)
