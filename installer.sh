#!/bin/bash

if [[ -z `go env GOPATH` ]]; then
echo "Go not found defaulting to binary installation"
binary_install
echo "Done! you can check with sshbruter --help"
else
echo "Go found on $(go env GOPATH). Do you want binary or source install (binary/source)"
read answer
if [[ $answer == "binary" ]]; then
binary_install
else
go get -v -u github.com/usdogu/sshbruter
echo "Done! Make sure your `go env GOPATH`/bin is in the path and check with instalation with sshbruter --help"
fi
fi

binary_install() {
    wget https://github.com/usdogu/sshbruter/releases/download/1.0.0/main
    chmod +x main
    sudo cp ./main /usr/local/bin/sshbruter
}