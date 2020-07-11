from dataclasses import dataclass
from flask_sqlalchemy import SQLAlchemy
from datetime import datetime
from simpleflake import simpleflake

db = SQLAlchemy()


@dataclass
class Message(db.Model):
    id: str
    username: str
    message: str

    __tablename__ = "messages"
    id = db.Column(db.Integer, primary_key=True, default=simpleflake)
    username = db.Column(db.String(16), nullable=False)
    message = db.Column(db.Text, nullable=False)
