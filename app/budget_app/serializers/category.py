from rest_framework import serializers
from budget_app.models import Category


class CategorySerializer(serializers.ModelSerializer):
    """Category serializer"""

    class Meta:
        """Meta class"""

        model = Category
        fields = ("id", "name")
