FROM node:18-alpine
WORKDIR /app
COPY package.json app.js ./
EXPOSE 3000
CMD ["node", "app.js"]
