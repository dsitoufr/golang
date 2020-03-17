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

findian
The user enter a string and the app checks if it starts with "i" ends with "n" and contains "a".
No matter if the string is in lower or upper case.

Makejson:
Enter your name and adress. The app creates a mapping between name and adress and display 
the mapping in json format.

Slice
Enter a serie a stream of numbers sequentially. The app stores these numbers in a slice
and sorts the slice on fly.

Trunc
check if the user entered a numbers

Read
Read a file containing in each line a name and surname. For each the app creates a struct
stores the struct in a slice.And then all on screen.