from flask_sqlalchemy import SQLAlchemy
from datetime import datetime
from simpleflake import simpleflake
db = SQLAlchemy()


class Message(db.Model):
    __tablename__ = "messages"
    id = db.Column(db.Integer, primary_key=True, default=simpleflake)
    username = db.Column(db.String(16), nullable=False)
    message = db.Column(db.Text, nullable=False)
    sent_at = db.Column(db.DateTime, default=datetime.utcnow)