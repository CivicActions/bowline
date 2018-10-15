```
cd $GOPATH/src/github.com/docker
git clone git@github.com:docker/engine.git docker
cd docker
git checkout v18.06.1-ce
cd $GOPATH/src/github.com/CivicActions/bowline
go get
cd fixtures
go run ../main.go
```