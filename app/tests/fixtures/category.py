import pytest


@pytest.fixture(scope="session")
def category():
    return {"name": "test_category"}


@pytest.fixture(scope="session")
def category2():
    return {"name": "test_category2"}


@pytest.fixture(scope="session")
def category3():
    return {"name": "test_category3"}
