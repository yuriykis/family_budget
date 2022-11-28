from rest_framework import serializers
from budget_app.serializers.category import CategorySerializer
from budget_app.models import Transaction


class GetTransactionSerializer(serializers.ModelSerializer):
    """Transaction serializer"""

    class Meta:
        """Meta class"""

        model = Transaction
        fields = ("id", "budget", "category", "amount", "datetime", "title", "type")

    category = CategorySerializer()


class CreateTransactionSerializer(serializers.ModelSerializer):
    """Transaction serializer"""

    class Meta:
        """Meta class"""

        model = Transaction
        fields = ("id", "budget", "category", "amount", "datetime", "title", "type")
