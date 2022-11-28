from rest_framework import permissions, status
from rest_framework.response import Response
from rest_framework.generics import GenericAPIView
from budget_app.models import Budget, UserBudget, Transaction
from utilities.logger import get_logger
from django.db.models import Sum

from budget_app.serializers.budget import BudgetSerializer


class BudgetView(GenericAPIView):
    """Budget view"""

    permission_classes = (permissions.IsAuthenticated,)
    serializer_class = BudgetSerializer
    logger = get_logger()

    def get(self, request):
        """Get budget"""
        try:
            self.logger.debug("Getting budget")
            # obtain budgets for request.user
            budgets = Budget.objects.filter(userbudget__user=request.user)

            # calculate amount_left and total_transactions for each budget
            for budget in budgets:
                budget_transactions = Transaction.objects.filter(budget=budget)
                transactions_aggregated = budget_transactions.aggregate(Sum("amount"))
                if transactions_aggregated["amount__sum"] is None:
                    budget.amount_left = budget.amount
                else:
                    budget.amount_left = (
                        budget.amount + transactions_aggregated["amount__sum"]
                    )
                budget.total_transactions = budget_transactions.count()

            serializer = self.serializer_class(budgets, many=True)
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
            serializer = self.serializer_class(data=request.data)
            serializer.is_valid(raise_exception=True)
            self.logger.debug("Data is valid")
            budget = serializer.save()
            UserBudget.objects.create(user=request.user, budget=budget)
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
            serializer = self.serializer_class(budget, data=request.data)
            serializer.is_valid(raise_exception=True)
            self.logger.debug("Data is valid")
            serializer.save()
            self.logger.debug("Budget %s updated", request.data["name"])
            return Response(serializer.data, status=status.HTTP_200_OK)
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
            budget.delete()
            self.logger.debug("Budget %s deleted", budget_id)
            return Response(status=status.HTTP_204_NO_CONTENT)
        except Exception as exception:
            self.logger.error("Error while deleting budget: %s", exception)
            return Response(
                {"message": str(exception)}, status=status.HTTP_400_BAD_REQUEST
            )
