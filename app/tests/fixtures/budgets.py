import pytest


@pytest.fixture(scope="session")
def budget():
    return {
        "name": "test_budget",
        "description": "test_description",
        "amount": 5000,
    }


@pytest.fixture(scope="session")
def budget2():
    return {
        "name": "test_budget2",
        "description": "test_description2",
        "amount": 1000,
    }


@pytest.fixture(scope="session")
def budget3():
    return {
        "name": "test_budget3",
        "description": "test_description3",
        "amount": 123,
    }
