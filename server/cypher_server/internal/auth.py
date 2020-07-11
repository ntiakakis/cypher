import itsdangerous
import base64
import binascii
import os
from dotenv import load_dotenv
from .errors import Unauthorized, Forbidden, Ratelimited

def check_token(token):
    if token is None:
        raise Unauthorized("Invalid token")

    try:
        itsdangerous.TimestampSigner(os.getenv("TOKEN_SECRET")).unsign(token, max_age=60*60*24)
    except itsdangerous.BadSignature:
        raise Forbidden("Invalid token")