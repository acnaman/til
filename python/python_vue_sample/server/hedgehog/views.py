from django.shortcuts import render
from django.contrib.auth.models import User
from rest_framework import generics, permissions
from rest_framework.permissions import IsAdminUser, IsAuthenticated

from .models import PileColor, HedgeHog, Comment
from .serializers import UserSerializer, PileColorSerializer, HedgeHogSerializer, CommentSerializer

class UserList(generics.ListAPIView):
    """ View to list all Users """

    queryset = User.objects.all().order_by('first_name')
    serializer_class = UserSerializer
    permission_classes = [IsAuthenticated]

class UserCreate(generics.CreateAPIView):
    """ View to create new user. Only accepts POST requests """

    queryset = User.objects.all()
    serializer_class = UserSerializer
    permission_classes = [IsAdminUser]

class UserRetrieveUpdate(generics.RetrieveUpdateAPIView):
    """ Retrieve a user or update user information 
    Accepts GET and PUT requests and the record id must be provided in the request """

    queryset = User.objects.all()
    serializer_class = UserSerializer
    permission_classes = [IsAuthenticated]

class PileColorListCreate(generics.ListCreateAPIView):
    """ List and Create PileColors """

    queryset = PileColor.objects.all()
    serializer_class = PileColorSerializer
    permission_classes = [IsAuthenticated]

class PileColorRetrieveUpdate(generics.RetrieveUpdateAPIView):
    """ Retrieve and Update PileColor information """

    queryset = PileColor.objects.all()
    serializer_class = PileColorSerializer
    permission_classes = [IsAuthenticated]

class HedgeHogListCreate(generics.ListCreateAPIView):
    """ List and Create HedgeHogs """

    queryset = HedgeHog.objects.all().order_by('name')
    serializer_class = HedgeHogSerializer
    permission_classes = [IsAuthenticated]

class HedgeHogRetrieveUpdate(generics.RetrieveUpdateAPIView):
    """ Retrieve and Update HedgeHog information """

    queryset = HedgeHog.objects.all()
    serializer_class = HedgeHogSerializer
    permission_classes = [IsAuthenticated]

class CommentListCreate(generics.ListCreateAPIView):
    """ List and Create Comments """

    queryset = Comment.objects.all()
    serializer_class = CommentSerializer
    permission_classes = [IsAuthenticated]

class CommentRetrieveUpdate(generics.RetrieveUpdateAPIView):
    """ Retrieve and Update Comment information """

    queryset = Comment.objects.all()
    serializer_class = CommentSerializer
    permission_classes = [IsAuthenticated]

