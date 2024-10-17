import requests

# Constants for API endpoints
LOGIN_URL = 'http://localhost:5000/login'
CREATE_GUEST_URL = 'http://localhost:5000/guests'

# Test cso credentials
USERNAME = 'admin'
PASSWORD = 'admin'

# Function to log in and retrieve the access token
def login():
    data = {
        'name': USERNAME,
        'password': PASSWORD
    }
    response = requests.post(LOGIN_URL, json=data)
    response.raise_for_status()  # Raise an error for bad responses
    return response.json()['access_token']

# Function to create a guest
def create_guest(access_token, guest_data):
    headers = {'Authorization': f'Bearer {access_token}'}
    response = requests.post(CREATE_GUEST_URL, json=guest_data, headers=headers)
    response.raise_for_status()  # Raise an error for bad responses
    return response

# Function to retrieve the created guest information
def get_guest(access_token):
    headers = {'Authorization': f'Bearer {access_token}'}
    response = requests.get(CREATE_GUEST_URL, headers=headers)
    response.raise_for_status()  # Raise an error for bad responses
    return response.json()

# Main test function for creating a guest
def test_create_user():
    access_token = login()

    # Define guest data
    guest_data = {
        "name": "test",
        'id_number': '123456789',
        'vehicle_plate': 'ABC123',
        'inviter': 'sabri',
        'purpose': 'test',
        'security_response': 'test',
        'entry_approved': True,
        'class_level': '2',
        'class_level_expiry': '2022-12-12',
        'security_approval': True,
        'approval_expiry': '2022-12-12',
        'notes': 'test'
    }
    
    # Create guest
    create_response = create_guest(access_token, guest_data)
    assert create_response.status_code == 201  # Check for successful creation

    # Retrieve the created guest
    retrieved_guest = get_guest(access_token)

    # Assertions to validate the retrieved guest data
    assert retrieved_guest['name'] == guest_data['name']
    assert retrieved_guest['id_number'] == guest_data['id_number']
    assert retrieved_guest['vehicle_plate'] == guest_data['vehicle_plate']
    assert retrieved_guest['is_inside'] is False  # Ensure the guest is not inside
    assert retrieved_guest['inviter'] == guest_data['inviter']
    assert retrieved_guest['purpose'] == guest_data['purpose']
    assert retrieved_guest['security_response'] == guest_data['security_response']
    assert retrieved_guest['entry_approved'] == guest_data['entry_approved']
    assert retrieved_guest['class_level'] == guest_data['class_level']
    assert retrieved_guest['class_level_expiry'] == guest_data['class_level_expiry']
    assert retrieved_guest['security_approval'] == guest_data['security_approval']
    assert retrieved_guest['approval_expiry'] == guest_data['approval_expiry']
    assert retrieved_guest['notes'] == guest_data['notes']


if __name__ == '__main__':
    print('Running test: ' + __file__)
    test_create_user()
