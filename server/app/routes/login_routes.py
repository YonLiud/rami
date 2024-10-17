from flask import Blueprint, jsonify, request
from ..models import CSO
from ..database import db
from flask_jwt_extended import create_access_token, jwt_required, JWTManager

login_bp = Blueprint('login_bp', __name__)
jwt = JWTManager()

@login_bp.route('/', methods=['POST'])
def login():
    data = request.json
    try:
        data['name']
        data['password']
    except KeyError:
        return jsonify({'err': 'Missing name or password'}), 400

    cso = CSO.query.filter_by(name=data['name']).first()
    if cso and cso.check_password(data['password']):
        access_token = create_access_token(identity=cso.uid)
        return jsonify(access_token=access_token), 200
    return jsonify({'err': 'Bad username or password'}), 401