FROM node:alpine as builder

WORKDIR /build

COPY . .

RUN npm install

RUN npm run build

FROM nginx

# Install dependencies
RUN apt-get update -y

# install app
RUN rm -rf /var/www/*
RUN mkdir /var/www/
RUN mkdir /var/www/uwalexaskills.org
COPY --from=builder /build/build /var/www/uwalexaskills.org/

# setup nginx
COPY uwalexaskills.org.conf /etc/nginx/conf.d/
RUN service nginx start

EXPOSE 80