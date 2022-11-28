from rest_framework import permissions, status
from rest_framework.response import Response
from rest_framework.generics import GenericAPIView
from budget_app.models import Transaction
from utilities.logger import get_logger

from budget_app.serializers.transaction import TransactionSerializer


class TransactionView(GenericAPIView):
    """Transaction view"""

    permission_classes = (permissions.IsAuthenticated,)
    serializer_class = TransactionSerializer
    logger = get_logger()

    def get(self, request, budget):
        """Get all transactions for budget"""
        try:
            self.logger.debug("Getting all transactions for budget %s", budget)
            transactions = Transaction.objects.filter(budget=budget)
            self.logger.debug("Transactions found: %s", transactions)
            serializer = self.serializer_class(transactions, many=True)
            return Response(serializer.data, status=status.HTTP_200_OK)
        except Exception as exception:
            self.logger.error("Error while getting transactions: %s", exception)
            return Response(
                {"message": str(exception)}, status=status.HTTP_400_BAD_REQUEST
            )

    def post(self, request):
        """Create new transaction"""
        try:
            self.logger.debug("Creating new transaction")
            serializer = self.serializer_class(data=request.data)
            serializer.is_valid(raise_exception=True)
            self.logger.debug("Data is valid")
            serializer.save()
            self.logger.debug("Transaction %s created", request.data["title"])
            return Response(serializer.data, status=status.HTTP_201_CREATED)
        except Exception as exception:
            self.logger.error("Error while creating transaction: %s", exception)
            return Response(
                {"message": str(exception)}, status=status.HTTP_400_BAD_REQUEST
            )
