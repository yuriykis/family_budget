from rest_framework import serializers
from budget_app.models import Budget


class BudgetSerializer(serializers.ModelSerializer):
    """Budget serializer"""

    class Meta:
        """Meta class"""

        model = Budget
        fields = (
            "id",
            "name",
            "description",
            "amount",
            "amount_left",
            "total_transactions",
        )

    amount_left = serializers.SerializerMethodField()
    total_transactions = serializers.SerializerMethodField()

    def get_amount_left(self, obj):
        """Get amount left"""
        return obj.amount_left

    def get_total_transactions(self, obj):
        """Get total transactions"""
        return obj.total_transactions
