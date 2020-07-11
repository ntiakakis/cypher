from sqlalchemy import create_engine, Table, Column, Integer, Text, String, DateTime, MetaData
from datetime import datetime
from simpleflake import simpleflake

engine = create_engine('sqlite:///local.db', echo = True)
meta = MetaData() 

user = Table(
    'users',
    meta,
    Column('id', Integer, primary_key = True),
    Column('username', String(16), nullable = False),
    Column('password', String(512), nullable = False),
    Column('email', String(64), nullable = False),
    Column('reg_time', DateTime, default = datetime.utcnow)
)

message = Table(
    'messages',
    meta,
    Column('id', Integer, primary_key = True, default = simpleflake()),
    Column('username', String(16), nullable = False),
    Column('message', Text, nullable = False),
    Column('sent_at', DateTime, default = datetime.utcnow)
)

meta.create_all(engine)