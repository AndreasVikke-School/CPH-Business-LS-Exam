FROM node:12

LABEL org.opencontainers.image.source https://github.com/andreasvikke/cph-business-ls-exam

# Create app directory
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

# Installing dependencies
COPY package*.json /usr/src/app/
RUN npm install

# Copying source files
COPY . /usr/src/app

# Building app
RUN npm run build

# Running the app
CMD "npm" "run" "dev"