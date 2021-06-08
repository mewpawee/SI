FROM node:14-slim
WORKDIR /usr/src/app
RUN apt update &&\
    apt install libreoffice -y --no-install-recommends
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 3000
CMD [ "node", "server.js" ]

