# go-http-proxy
A simple go program to proxy http request through a server with caching. Based on go-fiber

## Install requirements
For now, since there are no prebuilt binaries, you need to have go installed
```bash
apt install golang
```


## Usage/Installation
### Using go get
```bash
go get -v https://github.com/GQDeltex/go-http-proxy.git
```
### Using git and go install
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
### Adding to PATH
You can even add ~/go/bin to your path to have it even easier
```bash
# ~/.bashrc
PATH=$PATH:~/go/bin
```
Now you can run by typing
```bash
go-http-proxy
```
