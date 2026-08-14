FROM nginx:1.27-alpine

COPY services/frontend/nginx.conf /etc/nginx/nginx.conf
COPY services/frontend/html/ /usr/share/nginx/html/

EXPOSE 80

ENTRYPOINT ["nginx", "-g", "daemon off;"]
