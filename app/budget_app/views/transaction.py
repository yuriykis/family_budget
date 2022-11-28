from rest_framework import permissions, status
from rest_framework.response import Response
from rest_framework.generics import GenericAPIView
from budget_app.models import Transaction, Budget, Category
from django.db.models import Sum
from decimal import Decimal
from django.forms.models import model_to_dict
from utilities.logger import get_logger

from budget_app.serializers.transaction import (
    GetTransactionSerializer,
    CreateTransactionSerializer,
)


class TransactionView(GenericAPIView):
    """Transaction view"""

    permission_classes = (permissions.IsAuthenticated,)
    serializer_class = GetTransactionSerializer
    logger = get_logger()

    def get_serializer_class(self):
        if self.request.method == "POST":
            return CreateTransactionSerializer
        return self.serializer_class

    def get(self, request, **kwargs):
        """Get all transactions for budget"""
        if "budget_id" in kwargs and "transaction_id" in kwargs:
            transaction_id = kwargs.get("transaction_id")
            budget_id = kwargs.get("budget_id")
            try:
                self.logger.debug("Getting transaction %s", transaction_id)
                transaction = Transaction.objects.get(
                    id=transaction_id, budget_id=budget_id
                )
                self.logger.debug("Transaction found: %s", transaction)
                serializer = self.get_serializer(transaction)
                return Response(serializer.data, status=status.HTTP_200_OK)
            except Exception as exception:
                self.logger.error("Error while getting transaction: %s", exception)
                return Response(
                    {"message": str(exception)}, status=status.HTTP_400_BAD_REQUEST
                )
        if "budget_id" in kwargs:
            budget_id = kwargs.get("budget_id")
            try:
                self.logger.debug("Getting all transactions for budget %s", budget_id)
                transactions = Transaction.objects.filter(budget=budget_id)
                self.logger.debug("Transactions found: %s", transactions)
                serializer = self.get_serializer(transactions, many=True)
                return Response(serializer.data, status=status.HTTP_200_OK)
            except Exception as exception:
                self.logger.error("Error while getting transactions: %s", exception)
                return Response(
                    {"message": str(exception)}, status=status.HTTP_400_BAD_REQUEST
                )

        return Response(
            {"message": "Invalid request"}, status=status.HTTP_400_BAD_REQUEST
        )

    def post(self, request, budget_id):
        """Create new transaction"""
        try:
            self.logger.debug("Creating new transaction")
            transaction = request.data
            transaction["budget"] = budget_id
            serializer = self.get_serializer(data=transaction)
            serializer.is_valid(raise_exception=True)
            self.logger.debug("Data is valid")

            #  check budget amount left
            budget = Budget.objects.get(id=budget_id)
            budget_transactions = Transaction.objects.filter(budget=budget)
            transactions_aggregated = budget_transactions.aggregate(Sum("amount"))
            if transactions_aggregated["amount__sum"] is None:
                amount_left = budget.amount
            else:
                amount_left = budget.amount + transactions_aggregated["amount__sum"]
            total_transactions = budget_transactions.count()

            if transaction["type"] == Transaction.TransactionType.EXPENSE:
                transaction_amount = Decimal(transaction["amount"])
                if amount_left < transaction_amount:
                    self.logger.error("Not enough money left in the budget")
                    raise Exception("Not enough money left in the budget")

                # expense transaction should be negative
                serializer.validated_data["amount"] = (
                    serializer.validated_data["amount"] * -1
                )

            serializer.save()
            data_to_return = serializer.data
            data_to_return["amount_left"] = amount_left
            data_to_return["total_transactions"] = total_transactions + 1

            self.logger.debug("Transaction %s created", request.data["title"])
            return Response(data_to_return, status=status.HTTP_201_CREATED)
        except Exception as exception:
            self.logger.error("Error while creating transaction: %s", exception)
            return Response(
                {"message": str(exception)}, status=status.HTTP_400_BAD_REQUEST
            )
