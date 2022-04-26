https://www.youtube.com/watch?v=yyUHQIec83I
at 45 min mark

pkg std lib
https://pkg.go.dev/std

in terminal - go mod init booking-app// to initialize your project

in the main.go file, add package main to the top // all code must belong to a PACKAGE (package is a keyword)

the FIRST statement in the Go file must be 'package ... '

go needs a starting point. it is a function called main where all of your code will reside. (kind of like app)

you will need to import various go functions built in. like:
import "fmt". //fmt = format. it is a collection of go source files. here is a list -- https://pkg.go.dev/std

in terminal - go run main.go

All vars/consts values have data types and can be inferred at creation but you should be explicitly

When using int, decide if the integer can be negative. if not, use uint instead of int
