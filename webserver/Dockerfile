FROM node:10.21-alpine
ARG BUILD_ENV

#create destination directory
RUN mkdir -p /usr/src/app
WORKDIR /user/src/app


COPY . .
RUN npm install 
RUN npm run build

EXPOSE 3000

CMD [ "npm", "start"]


