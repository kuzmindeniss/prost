FROM node:22.15.0-alpine AS build
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run generate

CMD ["npm", "run", "preview"]
