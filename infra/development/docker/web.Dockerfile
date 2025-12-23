FROM node:20-alpine

WORKDIR /app

COPY web/package*.json ./

RUN npm install

COPY web ./

RUN npm run build

EXPOSE 3000

<<<<<<< HEAD
CMD ["npm", "start"] 
=======
CMD ["npm", "start"]
>>>>>>> 583e922a9184104830417e5b6115f8c54ac8bac7
