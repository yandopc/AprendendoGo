package main

import "fmt"

func fatR (n int) int {
   if n < 2 {
      return 1
   } 
   return n * fat (n-1)
}

func fat (n int) int {
   var f, x = n, n - 1
   for x > 1 {
      f, x = f * x, x - 1
   }
   return f
}

func main () {
   
   fmt.Println(fat(5))

}
