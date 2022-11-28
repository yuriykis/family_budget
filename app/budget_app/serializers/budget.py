from rest_framework import serializers
from budget_app.models import Budget


class BudgetSerializer(serializers.ModelSerializer):
    """Budget serializer"""

    class Meta:
        """Meta class"""

        model = Budget
        fields = ("id", "name", "description", "amount")
