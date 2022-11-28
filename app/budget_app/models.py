from django.db import models
from django.contrib.auth.models import User


class Budget(models.Model):
    """Main budget model"""

    name = models.CharField(max_length=100)
    description = models.TextField()
    amount = models.DecimalField(max_digits=10, decimal_places=2)
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)

    def __str__(self) -> str:
        return str(self.name)

    class Meta:
        """Order by created_at"""

        ordering = ["-created_at"]


class Category(models.Model):
    """Category model"""

    name = models.CharField(max_length=100, unique=True)
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)

    def __str__(self) -> str:
        return str(self.name)

    class Meta:
        """Order by name"""

        ordering = ["-name"]


class Transaction(models.Model):
    """Transaction model"""

    class TransactionType(models.TextChoices):
        """Transaction type choices"""

        INCOME = "INCOME", "Income"
        EXPENSE = "EXPENSE", "Expense"

    title = models.CharField(max_length=100)
    budget = models.ForeignKey(Budget, on_delete=models.CASCADE)
    category = models.ForeignKey(Category, on_delete=models.CASCADE)
    amount = models.DecimalField(max_digits=10, decimal_places=2)
    type = models.CharField(
        max_length=15, choices=TransactionType.choices, default=TransactionType.EXPENSE
    )
    datetime = models.DateTimeField(auto_now_add=True)
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)

    def __str__(self) -> str:
        return str(self.title)

    class Meta:
        """Order by created_at"""

        ordering = ["-created_at"]


class UserBudget(models.Model):
    """User budget model"""

    user = models.ForeignKey(User, on_delete=models.CASCADE)
    budget = models.ForeignKey(Budget, on_delete=models.CASCADE)
    ownership = models.BooleanField(default=False)
    readonly = models.BooleanField(default=False)
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)

    def __str__(self) -> str:
        return str(self.user)

    class Meta:
        """Order by created_at"""

        ordering = ["-created_at"]
