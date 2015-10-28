Docker Go Base
---

This repo is meant to be used as a starting point for go projects

It uses docker to compile the binaries and the main Dockerfile adds the linux
binary to the busybox image to create an extremely small final image

Building
---

```bash
go build -o sysnet_linux-amd64 -tags netgo *.go
# Or
./script/build
```
> `-tags netgo` will help you achieve static binaries :)

Running
---

```bash
./sysnet_linux-amd64
docker run --rm -ti pdxjohnny/sysnet
```

Changing The Name
---

```bash
./script/change-name $GITHUB_USERNAME $PROJECT_NAME
```


- John Andersen
