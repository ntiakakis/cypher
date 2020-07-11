from flask import Flask, Blueprint, request, jsonify
from dotenv import load_dotenv
from flask_sqlalchemy import SQLAlchemy
from ..internal import CheckToken
from ..models import user

import traceback
import jwt
import os

load_dotenv(verbose=True)

users = Blueprint('users', __name__)
db = SQLAlchemy()


@users.route('/users/register', methods=['POST'])
def register():
    if request.method == 'POST':
        if (request.json['username'] is None
                or request.json['password'] is None
                or request.json['email'] is None
                or request.json['email'] == ''
                or request.json['password'] == ''
                or request.json['username'] == ''):
            return "One or more of the required fields are missing", 400
        else:
            data_username = request.json['username']
            data_password = request.json['password']
            data_email = request.json['email']

            if bool(user.query.filter_by(username=data_username)) or bool(user.query.filter_by(email=data_email)):
                return "User already exists", 400
            else:
                try:
                    current_user = user(
                        username=data_username, password=data_password, email=data_email)
                    db.session.add(current_user)
                    db.session.commit()
                    return "Registered! You can now login"
                except:
                    traceback.print_exc()
                    return "Something went wrong", 500
    else:
        return "Method not allowed", 405


@users.route('/users/fetch', methods=['GET'])
def list_users():
    if request.method == 'GET':
        CheckToken(request.headers.get("Authorization"))
        db_users = user.query.order_by(user.id.desc()).all()

        return jsonify({
            "users": db_users
        })
    else:
        return "Method not allowed.", 405
