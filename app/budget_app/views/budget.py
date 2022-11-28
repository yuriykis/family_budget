from rest_framework import permissions, status
from rest_framework.response import Response
from rest_framework.generics import GenericAPIView
from budget_app.models import Budget, UserBudget, Transaction, User
from utilities.logger import get_logger
from django.db.models import Sum

from budget_app.serializers.budget import GetBudgetSerializer, CreateBudgetSerializer


class BudgetView(GenericAPIView):
    """Budget view"""

    permission_classes = (permissions.IsAuthenticated,)
    serializer_class = GetBudgetSerializer
    logger = get_logger()

    def get_serializer_class(self):
        if self.request.method == "POST":
            return CreateBudgetSerializer
        return self.serializer_class

    def get(self, request):
        """Get budget"""
        try:
            self.logger.debug("Getting budget")
            # obtain budgets for request.user
            budgets = Budget.objects.filter(userbudget__user=request.user)

            # calculate amount_left and total_transactions for each budget
            for budget in budgets:
                user_budget = UserBudget.objects.get(budget=budget, user=request.user)
                if user_budget.ownership == True:
                    budget.ownership = True
                else:
                    budget.ownership = False
                budget_transactions = Transaction.objects.filter(budget=budget)
                transactions_aggregated = budget_transactions.aggregate(Sum("amount"))
                if transactions_aggregated["amount__sum"] is None:
                    budget.amount_left = budget.amount
                else:
                    budget.amount_left = (
                        budget.amount + transactions_aggregated["amount__sum"]
                    )
                budget.total_transactions = budget_transactions.count()

            serializer = self.get_serializer(budgets, many=True)
            self.logger.debug("Budgets obtained")
            return Response(serializer.data, status=status.HTTP_200_OK)
        except Exception as exception:
            self.logger.error("Error while getting budget: %s", exception)
            return Response(
                {"message": str(exception)}, status=status.HTTP_400_BAD_REQUEST
            )

    def post(self, request):
        """Create budget"""
        try:
            self.logger.debug("Creating budget")
            serializer = self.get_serializer(data=request.data)
            serializer.is_valid(raise_exception=True)
            self.logger.debug("Data is valid")
            budget = serializer.save()
            UserBudget.objects.create(user=request.user, budget=budget, ownership=True)
            self.logger.debug("Budget %s created", request.data["name"])
            return Response(serializer.data, status=status.HTTP_201_CREATED)
        except Exception as exception:
            self.logger.error("Error while creating budget: %s", exception)
            return Response(
                {"message": str(exception)}, status=status.HTTP_400_BAD_REQUEST
            )

    def put(self, request, budget_id):
        """Update budget"""
        try:
            self.logger.debug("Updating budget")
            budget = Budget.objects.get(id=budget_id)
            # check if budget belongs to request.user
            if UserBudget.objects.filter(user=request.user, budget=budget).exists():
                serializer = self.get_serializer(budget, data=request.data)
                serializer.is_valid(raise_exception=True)
                self.logger.debug("Data is valid")
                serializer.save()
                self.logger.debug("Budget %s updated", request.data["name"])
                return Response(serializer.data, status=status.HTTP_200_OK)
            else:
                self.logger.error("Budget %s does not belong to user", budget.name)
                return Response(
                    {"message": "Budget does not belong to user"},
                    status=status.HTTP_400_BAD_REQUEST,
                )
        except Exception as exception:
            self.logger.error("Error while updating budget: %s", exception)
            return Response(
                {"message": str(exception)}, status=status.HTTP_400_BAD_REQUEST
            )

    def delete(self, request, budget_id):
        """Delete budget"""
        try:
            self.logger.debug("Deleting budget")
            budget = Budget.objects.get(id=budget_id)
            # check if budget belongs to request.user
            if UserBudget.objects.filter(user=request.user, budget=budget).exists():
                budget.delete()
                self.logger.debug("Budget %s deleted", budget.name)
                return Response(status=status.HTTP_204_NO_CONTENT)
            else:
                self.logger.error("Budget %s does not belong to user", budget.name)
                return Response(
                    {"message": "Budget does not belong to user"},
                    status=status.HTTP_400_BAD_REQUEST,
                )
        except Exception as exception:
            self.logger.error("Error while deleting budget: %s", exception)
            return Response(
                {"message": str(exception)}, status=status.HTTP_400_BAD_REQUEST
            )


class ShareBudgetView(GenericAPIView):
    """Share budget view"""

    permission_classes = (permissions.IsAuthenticated,)
    logger = get_logger()

    def post(self, request, budget_id):
        """Share budget"""
        try:
            self.logger.debug(
                "Sharing budget %s with user %s", budget_id, request.data["user_id"]
            )
            userbudget = UserBudget.objects.get(
                user=request.user, budget=budget_id, ownership=True
            )
            user = User.objects.get(id=request.data["user_id"])
            # check if user is already sharing budget
            if UserBudget.objects.filter(
                user=user, budget=budget_id, ownership=False
            ).exists():
                self.logger.error("User %s is already sharing budget", user.id)
                return Response(
                    {"message": "User is already sharing budget"},
                    status=status.HTTP_400_BAD_REQUEST,
                )
            new_userbudget = UserBudget(
                user=user, budget=userbudget.budget, ownership=False
            )
            new_userbudget.save()
            self.logger.debug(
                "Budget %s shared with user %s", budget_id, request.data["user_id"]
            )
            return Response(status=status.HTTP_200_OK)
        except Exception as exception:
            self.logger.error("Error while sharing budget: %s", exception)
            return Response(
                {"message": str(exception)}, status=status.HTTP_400_BAD_REQUEST
            )
