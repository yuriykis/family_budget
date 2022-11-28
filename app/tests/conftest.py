import pytest
import json

from django.urls import reverse


@pytest.fixture(scope="function")
def login_user(client, user):
    url = reverse("register")
    client.post(url, user, content_type="application/json")
    client.login(email=user["email"], password=user["password"])
