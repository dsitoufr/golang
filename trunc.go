package main

import (
"fmt"
)

func main() {
   var fnumber float64
   fmt.Printf("Enter a float number: ")
   _,err := fmt.Scan(&fnumber)
   if err != nil {
         fmt.Printf("Please enter a float number.")
   } else {
   	     fmt.Printf("%d", int(fnumber))
	}
}