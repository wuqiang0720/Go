# Go
```
wget https://dl.google.com/go/go1.21.0.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

GOOS=windows GOARCH=amd64 go build -o myapp.exe main.go

root@ubuntu:~# tree -d -L 4 go/
go/
├── pkg
│   └── mod
│       └── cache
│           └── vcs
└── src
    └── github.com
        ├── coreos
        │   └── go-systemd
        └── docker
            ├── coreos
            ├── go-connections
            └── go-plugins-helpers

```
