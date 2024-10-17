from flask import Flask
from .config import Config
from .database import db
from .models import CSO, Guest, Log
from .routes import register_routes
from flask_jwt_extended import JWTManager


from datetime import datetime

def create_app():
    app = Flask(__name__)
    app.config.from_object(Config)

    db.init_app(app)

    jwt = JWTManager(app)

    with app.app_context():
        db.create_all()

        if not CSO.query.first():
            print('Creating default CSO')        
            cso = CSO(name='admin', password='admin')
            db.session.add(cso)
            db.session.commit()
        else:
            print('CSO already exists, skipping')
        

    register_routes(app)

    return app
