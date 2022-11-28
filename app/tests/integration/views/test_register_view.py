import pytest
from django.urls import reverse
from rest_framework import status


@pytest.mark.django_db(transaction=True, reset_sequences=True)
class TestRegisterView:
    def test_register_view(self, client, user):
        url = reverse("register")
        response = client.post(url, user, content_type="application/json")
        assert response.status_code == status.HTTP_201_CREATED
