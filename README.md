go mod init name

How to add

1)
add import statement in the script

then, go run/build/test/list

// even install pseudo-versions

// go list -m -versions pkg
// ~  ~    -all

2)

go get pkg@v0.0.1

// go mod tidy // remove/tidies the gosum

3)
go get pkg@latest
go get pkg
go get ./...
go get -u // update recursively

4)


