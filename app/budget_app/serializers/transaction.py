from rest_framework import serializers
from budget_app.models import Transaction


class TransactionSerializer(serializers.ModelSerializer):
    """Transaction serializer"""

    class Meta:
        """Meta class"""

        model = Transaction
        fields = (
            "id",
            "budget",
            "category",
            "amount",
            "datetime",
            "title",
            "type",
        )
