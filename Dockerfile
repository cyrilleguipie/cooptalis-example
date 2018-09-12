# Use the official go docker image built on debian.
FROM golang:1.11.0-stretch

COPY . /go/src/cooptalis-example
WORKDIR /go/src/cooptalis-example

# Install go libraries (Revel and Dependencies
RUN set go get github.com/revel/revel \
    && go get github.com/revel/cmd/revel \
    && go get github.com/jinzhu/gorm \
    && go get github.com/casbin/casbin \
    && go get golang.org/x/crypto/bcrypt

EXPOSE 9000
ENTRYPOINT revel run cooptalis-example dev 9000


