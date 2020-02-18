package main

import (
"fmt"
"strings"
)

func main() {
   var TextStr string
   var LowerTextStr  string

   fmt.Printf("Enter a double quoted string:")
   _, err := fmt.Scanf("%q", &TextStr)
   
   if err != nil {
      fmt.Printf("Error enter a double quoted string !")
   } else {
      
      LowerTextStr = strings.ToLower(TextStr)
	  fmt.Printf("%s\n",LowerTextStr)
      if strings.HasPrefix(LowerTextStr, "i" ) &&  strings.HasSuffix(LowerTextStr, "n") && strings.Contains(LowerTextStr,"a") {
	       fmt.Printf("Found\n")
	  } else {
	       fmt.Printf("NotFound\n")
	    }  
	     
      
   }
}