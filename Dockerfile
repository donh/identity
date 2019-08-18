# Node image
FROM node:latest
# Maintainer
MAINTAINER Don Hsieh <don@idhub.network>

# basic tool
RUN apt-get -qq update
RUN apt-get -qqy install vim
RUN apt-get -qqy install sudo
RUN apt-get -qqy install python python-dev
RUN apt-get -qqy install python3.5-dev
RUN apt -qqy install python3-setuptools
RUN apt-get -qqy install python3-pip
RUN apt-get -qqy install ufw
RUN apt-get -qqy install mysql-server mysql-client

# Install Go
RUN cd /tmp && \
wget https://storage.googleapis.com/golang/go1.12.linux-amd64.tar.gz && \
tar -xvf go1.12.linux-amd64.tar.gz && \
mv go /usr/local

# Install the packages
RUN npm install -g --save ethereumjs-testrpc
# Expose port
EXPOSE 8545
# Start TestRPC
ENTRYPOINT ["testrpc"]
# ADD dev user
RUN useradd -m localadmin && echo "localadmin:admin_user" | chpasswd && adduser localadmin sudo
USER localadmin
# Install python-web3 by localadmin
### RUN pip3 install web3 flask flask-cors
# Create code directory
RUN mkdir /home/localadmin/code && \
mkdir /home/localadmin/goLib && \
echo 'export GOPATH=$HOME/goLib\nexport PATH=$PATH:/usr/local/go/bin\nexport PATH=$PATH:$GOPATH/bin' >> /home/localadmin/.bashrc && \
##### /bin/bash -c "source /home/localadmin/.bashrc" && \
export GOPATH=$HOME/goLib && \
export PATH=$PATH:/usr/local/go/bin && \
export PATH=$PATH:$GOPATH/bin && \
/usr/local/go/bin/go get github.com/ethereum/go-ethereum && \
/usr/local/go/bin/go get github.com/gorilla/mux && \
/usr/local/go/bin/go get github.com/gorilla/websocket && \
/usr/local/go/bin/go get github.com/jmoiron/sqlx && \
/usr/local/go/bin/go get github.com/gofrs/uuid && \
/usr/local/go/bin/go get github.com/go-sql-driver/mysql && \
/usr/local/go/bin/go get github.com/rs/cors && \
/usr/local/go/bin/go get gopkg.in/gomail.v2 && \
/usr/local/go/bin/go get gopkg.in/yaml.v2 && \
/usr/local/go/bin/go get golang.org/x/crypto/nacl/box && \
/usr/local/go/bin/go get github.com/golangci/golangci-lint/cmd/golangci-lint

# Set working directory
WORKDIR /home/localadmin/code
RUN npm install --save-exact openzeppelin-solidity
# Change root to install Truffle
USER root
RUN npm install -g truffle
RUN npm install --save-exact openzeppelin-solidity
RUN npm install -g solium
RUN npm install --save truffle-hdwallet-provider

# Change dev user back
USER localadmin
