# Go
```
root@ubuntu:~/Container-Registry/docker-ovs-plugin-go# go build   -o ovs-plugin main.go
/root/go/src/github.com/docker/go-plugins-helpers/sdk/unix_listener_systemd.go:10:2: cannot find package "github.com/coreos/go-systemd/activation" in any of:
        /usr/local/go/src/github.com/coreos/go-systemd/activation (from $GOROOT)
        /root/go/src/github.com/coreos/go-systemd/activation (from $GOPATH)
/root/go/src/github.com/docker/go-plugins-helpers/sdk/tcp_listener.go:8:2: cannot find package "github.com/docker/go-connections/sockets" in any of:
        /usr/local/go/src/github.com/docker/go-connections/sockets (from $GOROOT)
        /root/go/src/github.com/docker/go-connections/sockets (from $GOPATH)
root@ubuntu:~/Container-Registry/docker-ovs-plugin-go# cd ~/go/src/github.com/docker/
root@ubuntu:~/go/src/github.com/docker#
root@ubuntu:~/go/src/github.com/docker#
root@ubuntu:~/go/src/github.com/docker#
root@ubuntu:~/go/src/github.com/docker#
root@ubuntu:~/go/src/github.com/docker# git clone https://github.com/coreos/go-systemd.git
Cloning into 'go-systemd'...
remote: Enumerating objects: 2688, done.
remote: Counting objects: 100% (959/959), done.
remote: Compressing objects: 100% (303/303), done.
remote: Total 2688 (delta 745), reused 734 (delta 649), pack-reused 1729 (from 2)
Receiving objects: 100% (2688/2688), 785.14 KiB | 329.00 KiB/s, done.
Resolving deltas: 100% (1443/1443), done.
root@ubuntu:~/go/src/github.com/docker# git clone https://github.com/docker/go-connections.git
Cloning into 'go-connections'...
remote: Enumerating objects: 9333, done.
remote: Counting objects: 100% (228/228), done.
remote: Compressing objects: 100% (56/56), done.
remote: Total 9333 (delta 198), reused 172 (delta 172), pack-reused 9105 (from 2)
Receiving objects: 100% (9333/9333), 2.26 MiB | 28.00 KiB/s, done.
Resolving deltas: 100% (690/690), done.
root@ubuntu:~/go/src/github.com/docker#

```
