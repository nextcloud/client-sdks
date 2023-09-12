FROM alpine

RUN apk add \
    npm \
    openjdk17-jre-headless \
    composer php-xml php-simplexml php-xmlreader php-pdo php-tokenizer

RUN npm i -g @openapitools/openapi-generator-cli@2.7.0
RUN openapi-generator-cli version-manager set 7.0.0

ADD generate.sh /usr/local/bin
WORKDIR /src
CMD ["generate.sh"]
