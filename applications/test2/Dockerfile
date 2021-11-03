FROM node:12.18.1

LABEL org.opencontainers.image.source https://github.com/andreasvikke/CPH-Businnes-LS-Exam

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

CMD [ "node", "index.js" ]