# Backend
FROM python:3.7-slim
# update all
RUN apt update && apt upgrade -y --no-install-recommends
# install git
RUN apt install -y git python-pip python3-pip pipenv
# python package install
RUN pip install django

