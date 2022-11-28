from rest_framework import serializers
from budget_app.models import Budget


class GetBudgetSerializer(serializers.ModelSerializer):
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
            "ownership",
        )

    amount_left = serializers.SerializerMethodField()
    total_transactions = serializers.SerializerMethodField()
    ownership = serializers.SerializerMethodField()

    def get_amount_left(self, obj):
        """Get amount left"""
        return obj.amount_left

    def get_total_transactions(self, obj):
        """Get total transactions"""
        return obj.total_transactions

    def get_ownership(self, obj):
        """Get ownership"""
        return obj.ownership


class CreateBudgetSerializer(serializers.ModelSerializer):
    """Budget serializer"""

    class Meta:
        """Meta class"""

        model = Budget
        fields = ("id", "name", "description", "amount")

    def create(self, validated_data):
        """Create budget"""
        budget = Budget.objects.create(**validated_data)
        return budget
