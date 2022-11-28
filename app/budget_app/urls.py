from django.urls import path
from budget_app.views import register, budget, category, transaction, user

urlpatterns = [
    path("register/", register.RegisterView.as_view(), name="register"),
    path("budget/", budget.BudgetView.as_view(), name="budget"),
    path("budget/<int:budget_id>/", budget.BudgetView.as_view(), name="budget"),
    path(
        "budget/<int:budget_id>/transaction/",
        transaction.TransactionView.as_view(),
        name="transaction",
    ),
    path(
        "budget/<int:budget_id>/transaction/<int:transaction_id>/",
        transaction.TransactionView.as_view(),
        name="transaction",
    ),
    path(
        "budget/<int:budget_id>/share/", budget.ShareBudgetView.as_view(), name="share"
    ),
    path("category/", category.CategoryView.as_view(), name="category"),
    path("user/", user.UserView.as_view(), name="user"),
    path("user/<int:user_id>/", user.UserView.as_view(), name="user"),
    path("users/", user.UsersView.as_view(), name="user"),
]
