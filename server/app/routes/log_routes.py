from flask import Blueprint, jsonify, request
from ..models import CSO, Guest, Log
from ..database import db

log_bp = Blueprint('log_bp', __name__)

@log_bp.route('/', methods=['GET'])
def get_logs():
    logs = Log.query.all()
    return jsonify([log.to_dict() for log in logs])

@log_bp.route('/<guest_id>', methods=['GET'])
def get_guest_logs(guest_id):
    logs = Log.query.filter_by(guest_id=guest_id).all()
    return jsonify([log.to_dict() for log in logs])
