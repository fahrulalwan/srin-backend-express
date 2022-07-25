FROM node:lts-alpine as build-step
RUN mkdir /app
WORKDIR /app
COPY . /app
RUN npm install
COPY . .
RUN npm cache clean --force
RUN npm run start
EXPOSE 3000
