#!/bin/bash
go build .
sudo mv ./remix /usr/bin/remix
echo "Version $(remix --version) successfully installed"
