# build environment
FROM node:14.18.1 as build
WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

ARG REACT_APP_BASE_USER_URL
ENV REACT_APP_BASE_USER_URL $REACT_APP_BASE_USER_URL
ARG REACT_APP_BASE_MOVIE_URL
ENV REACT_APP_BASE_MOVIE_URL $REACT_APP_BASE_MOVIE_URL
ARG REACT_APP_BASE_SHOWTIME_URL
ENV REACT_APP_BASE_SHOWTIME_URL $REACT_APP_BASE_SHOWTIME_URL
ARG REACT_APP_BASE_RESERVATION_URL
ENV REACT_APP_BASE_RESERVATION_URL $REACT_APP_BASE_RESERVATION_URL
ARG REACT_APP_BASE_CINEMA_URL
ENV REACT_APP_BASE_CINEMA_URL $REACT_APP_BASE_CINEMA_URL

COPY package.json /app/
# COPY package-lock.json ./
COPY yarn.lock /app/

RUN yarn add sass
COPY . /app
RUN yarn
RUN yarn build && yarn --production

# production environment
FROM nginx:1.17-alpine
COPY --from=build /app/build/ /usr/share/nginx/html
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx/nginx.conf /etc/nginx/conf.d
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
