#!/bin/bash

# Install and run mkcert because fetching from http://localhost/ might be blocked
# due to the CORS and it bothers development.
sudo wget https://dl.filippo.io/mkcert/latest?for=linux/arm64 -O /usr/local/bin/mkcert
sudo chmod 755 /usr/local/bin/mkcert
mkcert -install
mkcert localhost

yarn install