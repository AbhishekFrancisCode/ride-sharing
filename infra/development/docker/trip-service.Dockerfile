FROM alpine
WORKDIR /app

ADD shared shared
ADD build build

RUN chmod +x build/trip-service

ENTRYPOINT build/trip-service