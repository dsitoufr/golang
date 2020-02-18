package main

import (
"fmt"
"encoding/json"
)

func main() {

  var name string
  var address  string
  
  m := make(map[string]string)
  
  fmt.Printf("Enter your name double quoted:")
  
  _, err1 := fmt.Scanf("%q\n",&name)
  
  if err1 != nil {
     fmt.Printf("*** Error enter a double quoted name string!\n")
	 
  }  else {
  
			fmt.Printf("Enter your address double quoted:")
  
            _, err2 := fmt.Scanf("%q\n",&address)
		    if err2 != nil {
		          fmt.Printf("*** Error enter a double quoted address string!\n")
           } else {
  
                   m[name] = address
  
                   barr, err3 := json.Marshal(m)
                   if err3 != nil {
                          fmt.Printf("*** Error converting map to json")
                   } else {
                           fmt.Printf("%s", string(barr))
		              }
           }
    }
}