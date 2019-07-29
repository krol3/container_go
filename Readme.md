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
ls /proc
hostname linux-br
```
