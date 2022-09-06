#!/bin/bash
go build .
sudo mv ./image-converter /usr/bin/imgconvert
echo "Version $(imgconvert --version) successfully installed"