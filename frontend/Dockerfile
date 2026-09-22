FROM node:latest

WORKDIR /app
COPY . .
RUN npm install
RUN npm run build

ENV API_KEY="7d572541135f588432728f7fa12b3b320abde07bab2dbe84"
EXPOSE 5173
CMD ["npm", "run", "preview"]
