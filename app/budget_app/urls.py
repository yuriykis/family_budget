from django.urls import path
from budget_app.views import (
    register,
)

urlpatterns = [path("register/", register.RegisterView.as_view(), name="register")]
