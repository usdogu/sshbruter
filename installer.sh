#!/bin/bash
wget https://github.com/usdogu/sshbruter/releases/download/1.0.0/main
chmod +x main
sudo cp ./main /usr/local/bin/sshbruter
echo "Done! you can check with sshbruter --help"

