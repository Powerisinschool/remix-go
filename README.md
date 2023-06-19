# Remix - Image Converter built with go

## Usage:
```shell
remix [format1] [format2] [--args]
```

## Build & Installation:
Run the following in your terminal

```shell
# Clone the repository
git clone https://github.com/Powerisinschool/remix-go.git
cd remix-go
go build .

# Optionally, add the directory to your PATH or move the
# generated file to your custom directory on the PATH
# For Mac/Linux:
cp ./remix /usr/bin/
```

## Example Usage:
Assuming a file is existent in the current directory called image.jpg.

To convert image.jpg to WebP format, run:

```shell
remix image.jpg image.webp
```

To open the new file after conversion, pass the --open argument when running the program like:

```shell
remix image.jpg image.webp --open
```
