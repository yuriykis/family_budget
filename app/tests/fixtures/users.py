import pytest


@pytest.fixture(scope="session")
def user():
    return {
        "username": "test_user",
        "password": "test_password",
        "first_name": "test_first_name",
        "last_name": "test_last_name",
        "email": "test_email",
    }


@pytest.fixture(scope="session")
def user2():
    return {
        "username": "test_user2",
        "password": "test_password2",
        "first_name": "test_first_name2",
        "last_name": "test_last_name2",
        "email": "test_email2",
    }


@pytest.fixture(scope="session")
def user3():
    return {
        "username": "test_user3",
        "password": "test_password3",
        "first_name": "test_first_name3",
        "last_name": "test_last_name3",
        "email": "test_email3",
    }
