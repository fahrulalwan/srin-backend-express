FROM node:lts-alpine as build-step
RUN mkdir /app
WORKDIR /app
COPY . /app
RUN npm ci
COPY . .
RUN npm cache clean --force
EXPOSE 3000
CMD ["npm", "start"]
