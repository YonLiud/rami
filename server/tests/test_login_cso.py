import pytest
import requests

LOGIN_URL = 'http://localhost:5000/login'
USERNAME = 'admin'
PASSWORD = 'admin'

@pytest.fixture
def valid_login_data():
    return {
        'name': USERNAME,
        'password': PASSWORD
    }

@pytest.fixture
def invalid_login_data():
    return {
        'name': 'wrong',
        'password': 'wrong'
    }

@pytest.fixture
def missing_data():
    return {
        'name': USERNAME
    }

@pytest.fixture
def empty_data():
    return {}

def test_login_success_with_valid_credentials(valid_login_data):
    response = requests.post(LOGIN_URL, json=valid_login_data)
    assert response.status_code == 200
    assert 'access_token' in response.json()

def test_login_failure_with_invalid_credentials(invalid_login_data):
    response = requests.post(LOGIN_URL, json=invalid_login_data)
    assert response.status_code == 401
    assert 'err' in response.json()

def test_login_failure_with_missing_data(missing_data):
    response = requests.post(LOGIN_URL, json=missing_data)
    assert response.status_code == 400
    assert 'err' in response.json()

def test_login_failure_with_empty_data(empty_data):
    response = requests.post(LOGIN_URL, json=empty_data)
    assert response.status_code == 400
    assert 'err' in response.json()

if __name__ == '__main__':
    pytest.main()
