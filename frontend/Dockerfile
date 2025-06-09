FROM node:22.15.0-alpine AS build
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .

ARG NUXT_PUBLIC_BACKEND_URL
ENV NUXT_PUBLIC_BACKEND_URL=${NUXT_PUBLIC_BACKEND_URL}

RUN npm run generate

CMD ["npm", "run", "preview"]
