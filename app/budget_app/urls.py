from django.urls import path
from budget_app.views import (
    register,
    budget,
    category,
    transaction,
)

urlpatterns = [
    path("register/", register.RegisterView.as_view(), name="register"),
    path("budget/", budget.BudgetView.as_view(), name="budget"),
    path("category/", category.CategoryView.as_view(), name="category"),
    path("transaction/", transaction.TransactionView.as_view(), name="transaction"),
]
