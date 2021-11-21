# go-http-proxy
A simple go program to proxy http requests through a server with caching

## Usage
All cli options are optional, and have the default values of
```
secret: 'secret'
cachetime: '30m'
allowedheaders: 'content-type'
```
To change these defaults, just run the command with options
```bash
go-http-server --secret 'supersecret' --cachetime 30m --allowedheaders 'content-type,etag,date'
```
After running the server you can access the service on port 3000 `http://localhost:3000/?url=<the-url-to-proxy&token=<the auth token>&expires=<token expiry time>`.
The token is a SHA256 Hash of the url + secret + expiry time. So as long as nobody knows your secret key nobody else can access the router.

## Installation
### Download latest release
go to the [releases](https://github.com/GQDeltex/go-http-proxy/releases/latest) page and download the latest release
```bash
# Untar the files
tar -xfv go-http-proxy-linux-amd64.tar.gz
# And start the executable
./go-http-server
```
### Using go get
```bash
go get -v https://github.com/GQDeltex/go-http-proxy.git
~/go/bin/go-http-server
```

## Building yourself
Install requirements 
```bash
apt install golang git
```
Clone the repo into your go workspace and get the dependencies
```bash
git clone https://github.com/GQDeltex/go-http-proxy.git ~/go/src/github.com/GQDeltex/go-http-proxyk
cd ~/go/src/github.com/GQDeltex/go-http-proxy
go mod download
```
After that you can either build the binary in that folder or install it into ~/go/bin/go-http-proxy
```bash
go build # Executable now in ./go-http-proxy
# Or install
go install # Executable now in ~/go/bin/go-http-proxy
```
Now you can run the proxy by starting the Executable
```bash
./go-http-proxy
# Or if you've used go install
~/go/bin/go-http-proxy
```
You can even add ~/go/bin to your path to have it even easier
```bash
# ~/.bashrc
PATH=$PATH:~/go/bin
```
Now you can run by typing
```bash
go-http-proxy
```
