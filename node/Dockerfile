FROM arm32v7/node:lts-alpine

WORKDIR /app

COPY . .

RUN npm ci

CMD ["npm", "start"]