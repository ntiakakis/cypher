from flask import Flask, Blueprint, request
from dotenv import load_dotenv
from flask_sqlalchemy import SQLAlchemy
from models import user
import traceback
import jwt
import os

load_dotenv(verbose = True)

users = Blueprint('users', __name__)
db = SQLAlchemy()

@users.route('/users/register', methods=['POST'])
def register():
    if request.method == 'POST':
        data_username = request.json['username']
        data_password = request.json['password']
        data_email = request.json['email']

        try:
            current_user = user(username=data_username, password=data_password, email=data_email)
            db.session.add(current_user)
            db.session.commit()    
            return "Request sent."
        except:
            traceback.print_exc()
            return "Something went wrong."
    else:
        return "Method not allowed."

@users.route('/users/login', methods=['POST'])
def login():
    print('a')

@users.route('/users/fetch', methods=['GET'])
def list_users():
    if request.method == 'GET':
        return str(users_db)
    else:
        return "Method not allowed."