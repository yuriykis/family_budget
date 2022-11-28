import pytest


@pytest.fixture(scope="session")
def transaction():
    return {
        "budget": 1,
        "category": 1,
        "date": "2020-01-01",
        "type": "EXPENSE",
        "title": "test_transaction",
        "description": "test_description",
        "amount": 5000,
    }


@pytest.fixture(scope="session")
def transaction2():
    return {
        "budget": 2,
        "category": 2,
        "date": "2020-01-01",
        "type": "EXPENSE",
        "title": "test_transaction2",
        "description": "test_description2",
        "amount": 1000,
    }


@pytest.fixture(scope="session")
def transaction3():
    return {
        "budget": 1,
        "category": 1,
        "date": "2020-01-01",
        "type": "INCOME",
        "title": "test_transaction3",
        "description": "test_description3",
        "amount": 123,
    }
