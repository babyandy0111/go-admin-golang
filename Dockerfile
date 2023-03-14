FROM alpine

RUN apk update --no-cache
RUN apk add --update gcc g++ libc6-compat
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ Asia/Taipei

COPY ./main /main
COPY config/dev.settings.yml /config/settings.yml
COPY ./docs/ /docs/

# 這邊要用PVC
COPY ./static /static

EXPOSE 8000
RUN  chmod +x /main
CMD ["/main","server","-c", "/config/settings.yml"]