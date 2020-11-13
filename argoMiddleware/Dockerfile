FROM node:14.0.0-alpine3.10  
ENV ARGO_SERVER='argo-server'

WORKDIR /usr/src/app  
COPY package*json ./  
RUN npm install  
COPY . . 
EXPOSE 3000 
CMD ["npm","run","startdev"]