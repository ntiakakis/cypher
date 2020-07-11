import itsdangerous
import base64
import binascii
from .errors import Unauthorized, Forbidden, Ratelimited

def check_token(token):
    if token is None:
        raise Unauthorized("Invalid token")
    try:
        if token.split(".")[0] != "cyphchat":
            raise Unauthorized("Invalid token")
    except IndexError:
        raise Unauthorized("Invalid token")

    try:
        itsdangerous.TimestampSigner(app.config["SECRET_KEY"].unsign(token, max_age=60*60*24))
    except itsdangerous.BadSignature:
        raise Forbidden("Invalid token")