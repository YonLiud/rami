import pytest
import requests
from datetime import datetime

LOGIN_URL = 'http://localhost:5000/login/'
CREATE_CSO_URL = 'http://localhost:5000/cso/'
CREATE_GUEST_URL = 'http://localhost:5000/guests/'

USERNAME = 'admin'
PASSWORD = 'admin'

@pytest.fixture
def guest_data():
    return {
        "name": "test50",
        'id_number': '0987654321',
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

@pytest.fixture
def cso_data():
    return {
        "name": "test",
        "password": "test",
    }

@pytest.fixture(scope='module')
def access_token():
    data = {
        'name': USERNAME,
        'password': PASSWORD
    }
    response = requests.post(LOGIN_URL, json=data)
    response.raise_for_status()
    return response.json()['access_token']

def test_create_cso(access_token, cso_data):
    headers = {'Authorization': f'Bearer {access_token}'}
    response = requests.post(CREATE_CSO_URL, headers=headers, json=cso_data)
    
    assert response.status_code in [201, 409]

def test_get_cso(access_token, cso_data):
    headers = {'Authorization': f'Bearer {access_token}'}
    
    response = requests.get(CREATE_CSO_URL + cso_data['name'], headers=headers)
    assert response.status_code == 200
    
    retrieved_cso = response.json()
    assert retrieved_cso['name'] == cso_data['name']

def test_login_cso(cso_data):
    response = requests.post(LOGIN_URL, json=cso_data)
    assert response.json()['access_token']

def test_create_guest_with_new_cso(access_token, cso_data, guest_data):
    response = requests.post(LOGIN_URL, json=cso_data)
    access_token = response.json()['access_token']

    global tmp_key
    tmp_key = access_token

    headers = {'Authorization': f'Bearer {access_token}'}

    response = requests.post(CREATE_GUEST_URL, json=guest_data, headers=headers)
    assert response.status_code in [201, 409]

def test_delete_guest_with_new_cso(access_token, cso_data, guest_data):
    response = requests.post(LOGIN_URL, json=cso_data)
    access_token = response.json()['access_token']
    headers = {'Authorization': f'Bearer {access_token}'}

    response = requests.delete(CREATE_GUEST_URL + guest_data['id_number'], headers=headers)
    assert response.status_code == 204

def test_delete_cso(access_token, cso_data):
    headers = {'Authorization': f'Bearer {access_token}'}
    
    response = requests.delete(CREATE_CSO_URL + cso_data['name'], headers=headers)
    assert response.status_code == 204
    
    response = requests.get(CREATE_CSO_URL + cso_data['name'], headers=headers)
    assert response.status_code == 404

def test_create_guest_with_deleted_cso(guest_data):
    headers = {'Authorization': f'Bearer {tmp_key}'}

    response = requests.post(CREATE_GUEST_URL, json=guest_data, headers=headers)
    assert response.status_code == 400