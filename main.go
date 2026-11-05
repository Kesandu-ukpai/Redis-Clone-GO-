   package main

   import "fmt"

   func main() {
        fmt.Printf("%q\n", EncodeSimple("Ok"))
       s := NewStore()
       startWeb(s)
   }

