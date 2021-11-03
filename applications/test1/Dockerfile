FROM node:12.18.1

LABEL org.opencontainers.image.source https://github.com/efio-dk/release-action-test

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

CMD [ "node", "index.js" ]