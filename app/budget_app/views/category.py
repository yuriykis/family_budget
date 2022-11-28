from rest_framework import permissions, status
from rest_framework.response import Response
from rest_framework.generics import GenericAPIView
from budget_app.models import Category
from utilities.logger import get_logger

from budget_app.serializers.category import CategorySerializer


class CategoryView(GenericAPIView):
    """Category view"""

    permission_classes = (permissions.IsAuthenticated,)
    serializer_class = CategorySerializer
    logger = get_logger()

    def get(self, request):
        """Get all categories"""
        try:
            self.logger.debug("Getting all categories")
            categories = Category.objects.all()
            self.logger.debug("Categories found: %s", categories)
            serializer = self.serializer_class(categories, many=True)
            return Response(serializer.data, status=status.HTTP_200_OK)
        except Exception as exception:
            self.logger.error("Error while getting categories: %s", exception)
            return Response(
                {"message": str(exception)}, status=status.HTTP_400_BAD_REQUEST
            )

    def post(self, request):
        """Create new category"""
        try:
            self.logger.debug("Creating new category")
            serializer = self.serializer_class(data=request.data)
            serializer.is_valid(raise_exception=True)
            self.logger.debug("Data is valid")
            serializer.save()
            self.logger.debug("Category %s created", request.data["name"])
            return Response(serializer.data, status=status.HTTP_201_CREATED)
        except Exception as exception:
            self.logger.error("Error while creating category: %s", exception)
            return Response(
                {"message": str(exception)}, status=status.HTTP_400_BAD_REQUEST
            )
