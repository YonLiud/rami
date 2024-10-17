from flask import Blueprint, jsonify, request
from ..models import CSO
from ..database import db
from flask_jwt_extended import jwt_required

cso_bp = Blueprint('cso_bp', __name__)

@cso_bp.route('/', methods=['GET'])
@jwt_required()
def get_csos():
    csos = CSO.query.all()
    return jsonify([cso.to_dict() for cso in csos])

@cso_bp.route('/<uid>', methods=['GET'])
@jwt_required()
def get_cso(uid):
    cso = CSO.query.get(uid)
    if cso:
        return jsonify(cso.to_dict())
    return jsonify({'err': 'CSO not found'}), 404

@cso_bp.route('/', methods=['POST'])
@jwt_required()
def create_cso():
    data = request.json
    try:
        cso = CSO(**data)
        db.session.add(cso)
        db.session.commit()
        return jsonify(cso.to_dict()), 201
    except Exception as e:
        return jsonify({'err': str(e)}), 400