from rest_framework import permissions, status
from rest_framework.response import Response
from rest_framework.generics import GenericAPIView
from django.contrib.auth.models import User
from utilities.logger import get_logger

from budget_app.serializers.user import UserSerializer


class UserView(GenericAPIView):
    """User view"""

    permission_classes = (permissions.IsAuthenticated,)
    serializer_class = UserSerializer
    logger = get_logger()

    def get(self, request, user_id=None):
        """Get my user"""
        try:
            if user_id:
                # Get user by id
                self.logger.debug("Getting user %s", user_id)
                user = User.objects.get(id=user_id)
            else:
                # Get my user
                self.logger.debug("Getting user %s", request.user)
                user = User.objects.get(id=request.user.id)
            self.logger.debug("User found: %s", user)
            serializer = self.serializer_class(user)
            return Response(serializer.data, status=status.HTTP_200_OK)
        except Exception as exception:
            self.logger.error("Error while getting user: %s", exception)
            return Response(
                {"message": str(exception)}, status=status.HTTP_400_BAD_REQUEST
            )


class UsersView(GenericAPIView):
    """Users view"""

    permission_classes = (permissions.IsAuthenticated,)
    serializer_class = UserSerializer
    logger = get_logger()

    def get(self, request):
        """Get all users"""
        try:
            self.logger.debug("Getting all users")
            users = User.objects.all()
            self.logger.debug("Users found: %s", users)
            serializer = self.serializer_class(users, many=True)
            return Response(serializer.data, status=status.HTTP_200_OK)
        except Exception as exception:
            self.logger.error("Error while getting users: %s", exception)
            return Response(
                {"message": str(exception)}, status=status.HTTP_400_BAD_REQUEST
            )
