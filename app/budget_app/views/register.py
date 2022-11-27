from rest_framework import permissions, status
from rest_framework.response import Response
from rest_framework.generics import GenericAPIView
from utilities.logger import get_logger

from budget_app.serializers.register import RegisterSerializer


class RegisterView(GenericAPIView):
    """Register view"""

    permission_classes = (permissions.AllowAny,)
    authentication_classes = ()
    serializer_class = RegisterSerializer
    logger = get_logger()

    def post(self, request):
        """Create new user"""
        try:
            self.logger.debug("Registering new user")
            serializer = self.serializer_class(data=request.data)
            serializer.is_valid(raise_exception=True)
            self.logger.debug("Data is valid")
            serializer.save()
            self.logger.debug("User %s created", request.data["email"])
            return Response(serializer.data, status=status.HTTP_201_CREATED)
        except Exception as exception:
            self.logger.error("Error while registering user: %s", exception)
            return Response(
                {"message": str(exception)}, status=status.HTTP_400_BAD_REQUEST
            )
