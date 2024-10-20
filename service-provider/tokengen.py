import jwt
import datetime
import os
import sys


SECRET_KEY = os.environ.get('SECRET_KEY', '0123456789')


def generate_token(user: str, scope: str):
    token = jwt.encode({
        'user': user,
        'scope': scope,
        'exp': datetime.datetime.now(datetime.timezone.utc) + datetime.timedelta(seconds=60)
    }, SECRET_KEY, algorithm="HS256")
    return token


if __name__ == '__main__':
    if len(sys.argv) == 3:
        user = sys.argv[1]
        scope = sys.argv[2]
        token = generate_token(user, scope)
        print(f'User:  \033[35m{user}\033[0m')
        print(f'Scope: \033[33m{scope}\033[0m')
        print(f'Token: \033[32m{token}\033[0m')
    else:
        print(f'Error: \033[31mbad number of arguments!\033[0m')
        print(f'Usage: \033[31m{sys.argv[0]} USER SCOPE\033[0m')
