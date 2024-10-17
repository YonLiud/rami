from flask import Blueprint, jsonify, request
from ..models import CSO, Guest, Log
from ..models.log_model import event_types as et
from ..database import db
from sqlalchemy.exc import IntegrityError
from flask_jwt_extended import jwt_required
from datetime import datetime

guest_bp = Blueprint('guest_bp', __name__)

def log_event(guest_id, event):
    log_entry = Log(guest_id=guest_id, event=event, timestamp=datetime.now())
    db.session.add(log_entry)
    db.session.commit()

@guest_bp.route('/', methods=['GET'])
def get_guests():
    guests = Guest.query.all()
    return jsonify([guest.to_dict() for guest in guests])

# Get all Guests inside
@guest_bp.route('/inside', methods=['GET'])
def get_guests_inside():
    guests = Guest.query.filter_by(is_inside=True).all()
    return jsonify([guest.to_dict() for guest in guests])

# Get a Guest by id_number
@guest_bp.route('/<id_number>', methods=['GET'])
def get_guest(id_number):
    guest = Guest.query.filter_by(id_number=id_number).first()
    if guest:
        return jsonify(guest.to_dict())
    return jsonify({'err': 'Guest not found'}), 404

# Create a Guest ( PROTECTED )
@guest_bp.route('/', methods=['POST'])
@jwt_required()
def create_guest():
    data = request.json
    try:
        data['class_level_expiry'] = datetime.fromisoformat(data['class_level_expiry'])
        data['approval_expiry'] = datetime.fromisoformat(data['approval_expiry'])
    except ValueError:
        return jsonify({'err': 'Invalid date format. Use ISO 8601 format.'}), 400
    try:
        guest = Guest(**data)
        db.session.add(guest)
        db.session.commit()
        log_event(guest.id_number, 'Guest created')
        return jsonify(guest.to_dict()), 201
    except IntegrityError:
        db.session.rollback()
        return jsonify({'err': 'A guest with this ID number already exists.'}), 409
    except Exception as e:
        db.session.rollback()
        return jsonify({'err': str(e)}), 400

# Delete a Guest ( PROTECTED )
@guest_bp.route('/<id_number>', methods=['DELETE'])
@jwt_required()
def delete_guest(id_number):
    guest = Guest.query.filter_by(id_number=id_number).first()
    if guest:
        log_event(guest.id_number, et['guest_deleted'])
        db.session.delete(guest)
        db.session.commit()
        return jsonify({'msg': 'Guest deleted'}), 204
    return jsonify({'err': 'Guest not found'}), 404

# Mark entry / exit of a Guest
@guest_bp.route('/<id_number>', methods=['PUT'])
def mark_guest(id_number):
    guest = Guest.query.filter_by(id_number=id_number).first()
    if guest:
        if guest.is_inside:
            guest.is_inside = False
            log_event(guest.id_number, et['guest_exit'])
        else:
            guest.is_inside = True
            log_event(guest.id_number, et['guest_entry'])
        db.session.commit()
        return jsonify(guest.to_dict())
    return jsonify({'err': 'Guest not found'}), 404