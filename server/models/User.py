from flask_sqlalchemy import SQLAlchemy
from datetime import datetime
from simpleflake import simpleflake

db = SQLAlchemy()


class User(db.Model):
    __tablename__ = "users"
    id = db.Column(db.Integer, primary_key=True, default=simpleflake)
    username = db.Column(db.String(16), nullable=False)
    password = db.Column(db.String(512), nullable=False)
    email = db.Column(db.String(64), nullable=False)
    reg_time = db.Column(db.DateTime, default=datetime.utcnow)