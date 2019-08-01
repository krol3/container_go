# Building a Container with Golang

`cmd := exec.Command("ls", "-lah")`

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
