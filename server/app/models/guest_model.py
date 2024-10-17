from ..database import db

class Guest(db.Model):
    uid =               db.Column(db.Integer, primary_key=True)
    name =              db.Column(db.String(50),    nullable=False)
    id_number =         db.Column(db.String(9),     nullable=False, unique=True)
    vehicle_plate =     db.Column(db.String(50),    nullable=False)
    inviter =           db.Column(db.String(50),    nullable=False)
    purpose =           db.Column(db.String(50),    nullable=False)
    is_inside =         db.Column(db.Boolean,       nullable=False, default=False)
    security_response = db.Column(db.String(50),    nullable=False)
    entry_approved =    db.Column(db.Boolean,       nullable=False)
    class_level =       db.Column(db.String(50),    nullable=False)
    class_level_expiry =db.Column(db.DateTime,      nullable=False)
    security_approval = db.Column(db.Boolean,       nullable=False)
    approval_expiry =   db.Column(db.DateTime,      nullable=False)
    notes =             db.Column(db.String(50),    nullable=False)

    def __repr__(self) -> str:
        return f'<Guest {self.name}>'
    
    def to_dict(self):
        return {
            'uid': self.uid,
            'name': self.name,
            'id_number': self.id_number,
            'vehicle_plate': self.vehicle_plate,
            'inviter': self.inviter,
            'purpose': self.purpose,
            'is_inside': self.is_inside,
            'security_response': self.security_response,
            'entry_approved': self.entry_approved,
            'class_level': self.class_level,
            'class_level_expiry': self.class_level_expiry.isoformat() if self.class_level_expiry else None,
            'security_approval': self.security_approval,
            'approval_expiry': self.approval_expiry.isoformat() if self.approval_expiry else None,
            'notes': self.notes
    }