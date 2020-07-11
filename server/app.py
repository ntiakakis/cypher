from flask import Flask, jsonify, request
from flask_sqlalchemy import SQLAlchemy
from blueprints import messages, users
from dotenv import load_dotenv
from models import user as User
from datetime import datetime
from internal import CheckToken
import itsdangerous
import traceback
import secrets
import base64
import os

app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = "sqlite:///local.db"
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False
app.config['SECRET_KEY'] = os.getenv("TOKEN_SECRET")
app.db = SQLAlchemy(app)

@app.route("/auth/validate")
def validate():
    CheckToken(request.headers.get("Authorization"))
    return "", 204

@app.route("/auth/generate")
def generate():
    data_user = request.json['username']
    data_password = request.json['password']
    try:
        current_user = User.query.filter_by(username=data_user).first()
        if current_user and current_user.password == data_password:
            return jsonify({"token": itsdangerous.TimestampSigner(app.config["SECRET_KEY"]).sign("cyphchat").decode()})
    except:
        traceback.print_exc()
        return "Something went wrong..", 400

app.register_blueprint(messages)
app.register_blueprint(users)

if __name__ == "__main__":
    app.run(debug = True)