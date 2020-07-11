from sys import argv
from flask import Flask, Blueprint, request, jsonify
from flask_sqlalchemy import SQLAlchemy
from ..models import message
from ..internal import CheckToken
import traceback
import pytz
import json

messages = Blueprint('messages', __name__)
db = SQLAlchemy()


@messages.route('/messages', methods=['POST'])
def send_message():
    if request.method == 'POST':
        CheckToken(request.headers.get("Authorization"))
        data_username = str(request.json['user'])
        data_message = str(request.json['message'])

        try:
            current_message = message(
                username=data_username, message=data_message)
            db.session.add(current_message)
            db.session.commit()
        except:
            traceback.print_exc()
            return "Something went wrong", 500

        return "Request sent"
    else:
        return "Method not allowed", 405


@messages.route('/messages/fetch')
def get_messages():
    if request.method == 'GET':
        CheckToken(request.headers.get("Authorization"))
        db_messages = message.query.order_by(message.id.desc()).all()
        resp = [{"id": str(row.id), "message": row.message,
                 "username": row.username} for row in db_messages]

        return jsonify({
            "messages": resp
        })
    else:
        return "Method not allowed", 405
