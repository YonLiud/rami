from flask import Blueprint, jsonify, request
from ..models import CSO
from ..database import db
from flask_jwt_extended import jwt_required
from sqlalchemy.exc import IntegrityError
from flask_jwt_extended import jwt_required, get_jwt

cso_bp = Blueprint('cso_bp', __name__)

@cso_bp.route('/', methods=['GET'])
@jwt_required()
def get_csos():
    csos = CSO.query.all()
    return jsonify([cso.to_dict() for cso in csos])

@cso_bp.route('/<name>', methods=['GET'])
@jwt_required()
def get_cso(name):
    cso = CSO.query.filter_by(name=name).first()
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
    except IntegrityError:
        db.session.rollback()
        return jsonify({'err': 'A CSO with this name already exists.'}), 409
    except Exception as e:
        print(e)
        return jsonify({'err': str(e)}), 400
    
@cso_bp.route('/<name>', methods=['DELETE'])
@jwt_required()
def delete_cso(name):
    # Implement JWT token revocation 
    cso = CSO.query.filter_by(name=name).first()
    if cso:
        db.session.delete(cso)
        db.session.commit()
        return '', 204
    return jsonify({'err': 'CSO not found'}), 404
