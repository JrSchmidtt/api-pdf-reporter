FROM surnet/alpine-wkhtmltopdf:3.9-0.12.5-full as wkhtmltopdf
FROM openjdk:8-jdk-alpine3.9

# Install dependencies for wkhtmltopdf
RUN apk add --no-cache \
  libstdc++ \
  libx11 \
  libxrender \
  libxext \
  libssl1.1 \
  ca-certificates \
  fontconfig \
  freetype \
  ttf-dejavu \
  ttf-droid \
  ttf-freefont \
  ttf-liberation \
  ttf-ubuntu-font-family \
&& apk add --no-cache --virtual .build-deps \
  msttcorefonts-installer \
\
# Install microsoft fonts
&& update-ms-fonts \
&& fc-cache -f \
\
# Clean up when done
&& rm -rf /tmp/* \
&& apk del .build-deps

#app builder
FROM golang:1.19 as appbuild

WORKDIR /app
COPY . ./
RUN go build -o /server

# Image
FROM minidocks/wkhtmltopdf

WORKDIR /app

COPY fonts/ /usr/share/fonts
COPY templates/ ./templates/
RUN mkdir tmp
COPY --from=appbuild /server ./server
COPY --from=wkhtmltopdf /bin/wkhtmltopdf /bin/wkhtmltopdf

EXPOSE 3000

ENTRYPOINT [ "./server" ]