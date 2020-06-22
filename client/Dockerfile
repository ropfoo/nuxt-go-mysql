FROM node:latest

WORKDIR /usr/src/app


EXPOSE 3000

ENV NUXT_HOST=0.0.0.0
ENV NUXT_PORT=3000

COPY package*.json ./
RUN npm install
COPY . .

# build production app
RUN npm run build

# start service
CMD npm start
