from flask import Blueprint
from .guest_routes import guest_bp
from .cso_routes import cso_bp
from .log_routes import log_bp
from .login_routes import login_bp

def register_routes(app):
    app.register_blueprint(guest_bp, url_prefix='/guests')
    app.register_blueprint(cso_bp, url_prefix='/cso')
    app.register_blueprint(log_bp, url_prefix='/logs')
    app.register_blueprint(login_bp, url_prefix='/login')
