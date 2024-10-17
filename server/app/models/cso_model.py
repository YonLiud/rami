from ..database import db
from flask_bcrypt import Bcrypt

bcrypt = Bcrypt()

class CSO(db.Model):
    uid =               db.Column(db.Integer, primary_key=True)
    name =              db.Column(db.String(50), nullable=False)
    hashed_password =   db.Column(db.String(128), nullable=False)

    def __init__(self, name, password):
        self.name = name
        self.hashed_password = bcrypt.generate_password_hash(password).decode('utf-8')

    def check_password(self, password):
        return bcrypt.check_password_hash(self.hashed_password, password)

    def __repr__(self) -> str:
        return f'<CSO {self.name}>'
    
    def to_dict(self):
        return {
            'uid': self.uid,
            'name': self.name
        }