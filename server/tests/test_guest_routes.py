import pytest
import requests
from datetime import datetime

LOGIN_URL = 'http://localhost:5000/login/'
CREATE_GUEST_URL = 'http://localhost:5000/guests/'

USERNAME = 'admin'
PASSWORD = 'admin'

@pytest.fixture(scope='module')
def access_token():
    data = {
        'name': USERNAME,
        'password': PASSWORD
    }
    response = requests.post(LOGIN_URL, json=data)
    response.raise_for_status()
    return response.json()['access_token']

@pytest.fixture
def guest_data():
    return {
        "name": "test",
        'id_number': '123456789',
        'vehicle_plate': 'ABC123',
        'inviter': 'sabri',
        'purpose': 'test',
        'security_response': 'test',
        'entry_approved': True,
        'class_level': '2',
        'class_level_expiry': datetime.now().isoformat(),
        'security_approval': True,
        'approval_expiry': datetime.now().isoformat(),
        'notes': 'test'
    }

def create_guest(access_token, guest_data):
    headers = {'Authorization': f'Bearer {access_token}'}
    response = requests.post(CREATE_GUEST_URL, json=guest_data, headers=headers)
    return response

def get_guest(access_token, guest_id):
    headers = {'Authorization': f'Bearer {access_token}'}
    response = requests.get(f'{CREATE_GUEST_URL}{guest_id}', headers=headers)
    return response

def delete_guest(access_token, guest_id):
    headers = {'Authorization': f'Bearer {access_token}'}
    response = requests.delete(f'{CREATE_GUEST_URL}{guest_id}', headers=headers)
    return response

def test_create_guest(access_token, guest_data):
    create_response = create_guest(access_token, guest_data)
    assert create_response.status_code == 201

def test_get_guest(access_token, guest_data):
    retrieved_guest = get_guest(access_token, guest_data['id_number']).json()   

    assert retrieved_guest['name'] == guest_data['name']
    assert retrieved_guest['id_number'] == guest_data['id_number']
    assert retrieved_guest['vehicle_plate'] == guest_data['vehicle_plate']
    assert retrieved_guest['is_inside'] is False
    assert retrieved_guest['inviter'] == guest_data['inviter']
    assert retrieved_guest['purpose'] == guest_data['purpose']
    assert retrieved_guest['security_response'] == guest_data['security_response']
    assert retrieved_guest['entry_approved'] == guest_data['entry_approved']
    assert retrieved_guest['class_level'] == guest_data['class_level']
    assert retrieved_guest['class_level_expiry'] == guest_data['class_level_expiry']
    assert retrieved_guest['security_approval'] == guest_data['security_approval']
    assert retrieved_guest['approval_expiry'] == guest_data['approval_expiry']
    assert retrieved_guest['notes'] == guest_data['notes']


def test_mark_entry(access_token, guest_data):
    retrieved_guest = get_guest(access_token, guest_data['id_number']).json()
    assert retrieved_guest['is_inside'] is False

    requests.put(CREATE_GUEST_URL+guest_data['id_number'])
    retrieved_guest = get_guest(access_token, guest_data['id_number']).json()
    assert retrieved_guest['is_inside'] is True

def test_conflict_create_guest(access_token, guest_data):
    create_response = create_guest(access_token, guest_data)
    assert create_response.status_code == 409

def test_delete_guest(access_token, guest_data):
    delete_response = delete_guest(access_token, guest_data['id_number'])
    assert delete_response.status_code == 204

    retrieved_guest = get_guest(access_token, guest_data['id_number'])
    assert retrieved_guest.status_code == 404

if __name__ == '__main__':
    pytest.main()
