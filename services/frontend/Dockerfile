FROM node:20-alpine AS build
WORKDIR /app

COPY services/frontend/package*.json ./
RUN npm install

COPY services/frontend/ .
RUN npm run build

FROM nginx:1.27-alpine

COPY services/frontend/nginx.conf /etc/nginx/nginx.conf
COPY --from=build /app/dist /usr/share/nginx/html/

EXPOSE 80

ENTRYPOINT ["nginx", "-g", "daemon off;"]
