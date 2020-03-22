# Building a Container with Golang

### Virtualbox with vagrant

`cd vagrant && make up`

## Branchs

### f01: Namespaces

```
make run
hostname
hostname linux-br
```

### f02: PID

- fork exec

```
make run
ls
ps
ls /proc
hostname linux-br
```

### f021: PID - Chroot

```
make run
ls
ps
```

### f022: PID - Chroot - Mount proc

```
make run
ls
ps
ls /proc
```

## Resources

[Julian Friedman - Build your own container](https://www.infoq.com/articles/build-a-container-golang/)
Liz Rice - Building a container from scratch in Go
[Ed King - Namespaces](https://medium.com/@teddyking/linux-namespaces-850489d3ccf)
