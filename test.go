// go language use package to manage program, every file need a package name at head
package main

// fmt package (I/O package) has Println method, we should be add this package
import "fmt"

func main() {
   fmt.Println("Hello, World!")

   var a, b  int
   a = 10
   var c int = 10
   d := 20
   fmt.Println( a, b, c, d )
}


// go command:
//              go version
//              go run test.go   -> run go application
//              go build test.go  -> build direct application in OS

