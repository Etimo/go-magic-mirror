FROM golang:1.15.3

RUN apt-get update
RUN apt-get -y install curl gnupg
RUN curl -sL https://deb.nodesource.com/setup_15.x  | bash -
RUN apt-get -y install nodejs
RUN npm install yarn -g

WORKDIR /go/src/go-magic-mirror
COPY package.json yarn.lock ./
COPY main.go .

RUN yarn
RUN yarn install-go