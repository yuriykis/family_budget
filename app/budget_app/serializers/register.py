from rest_framework import serializers
from django.contrib.auth.models import User


class RegisterSerializer(serializers.ModelSerializer):
    """Register serializer"""

    class Meta:
        """Meta class"""

        model = User
        fields = ("id", "email", "username", "password", "first_name", "last_name")
        extra_kwargs = {
            "password": {"write_only": True},
            "first_name": {"required": True},
            "last_name": {"required": True},
            "email": {"required": True},
        }

    def create(self, validated_data):
        """Create user"""
        user = User.objects.create_user(**validated_data)
        return user
