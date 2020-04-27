FROM node:14-slim

# update all
RUN apt update && apt upgrade -y --no-install-recommends

# install git
RUN apt install -y git python procps yarn build-essential python-pip python3-pip

# install vue-cli
RUN npm install -g @vue/cli

# python package install
RUN pip install django

# install yarn package
RUN yarn add axios sweetalert2 vue-session
