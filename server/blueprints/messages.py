from flask import Flask, Blueprint, request
from flask_sqlalchemy import SQLAlchemy
from datetime import datetime
from models import message
import traceback
import pytz

messages = Blueprint('messages', __name__)
db = SQLAlchemy()

@messages.route('/messages', methods=['POST'])
def send_message():
    if request.method == 'POST':
        data_username = str(request.json['user'])
        data_message = str(request.json['message'])

        try:
            current_message = message(username=data_username, message=data_message)
            db.session.add(current_message)
            db.session.commit()
        except:
            traceback.print_exc()
            return "Something went wrong."
        return "Request sent."
    else:
        return "Method not allowed"

@messages.route('/messages/fetch')
def get_messages():
    if request.method == 'GET':
        return str(database)
    else:
        return "Method not allowed"