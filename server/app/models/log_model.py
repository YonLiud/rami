from ..database import db

class Log(db.Model):
    uid =               db.Column(db.Integer, primary_key=True)
    timestamp =         db.Column(db.DateTime, nullable=False)
    event =             db.Column(db.String(50), nullable=False)
    guest_id =          db.Column(db.Integer, nullable=False)

    def __repr__(self) -> str:
        return f"<Log {self.event}> | {self.timestamp} | {self.guest_id}"
    
    def to_dict(self):
        return {
            'uid': self.uid,
            'timestamp': self.timestamp.isoformat(),
            'event': self.event,
            'guest_id': self.guest_id
        }

    # log table with types of events
event_types = {
    'guest_created':    'Guest created',
    'guest_deleted':    'Guest deleted',
    'guest_updated':    'Guest updated',
    'guest_entry':      'Guest entered',
    'guest_exit':       'Guest left',
}