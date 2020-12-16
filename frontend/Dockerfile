FROM node:alpine

WORKDIR /frontend

COPY package.json /frontend

RUN yarn install

COPY . /frontend

CMD ["yarn", "run", "start"]

