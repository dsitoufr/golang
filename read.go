package main

import (
"fmt"
"io/ioutil"
"strings"
)

//
//split file data into lines
//

func SplitLines(b []byte) []string {
  w := string(b)
  k := strings.Split(w,"\n")
  return k
}



func main() {

 var filePath string
 
 type name struct {
   fname string
   lname string
 }
 
 s := []name{}
 
 fmt.Printf("Enter the file name: ")
 
 _, err := fmt.Scan(&filePath)
 if err != nil {
 
     fmt.Printf("*** Error reading file name !\n")
	 
 } else {
         
         r, err := ioutil.ReadFile(filePath)
		 
		 if err != nil {
		 
		    fmt.Printf("***Error invalid file !\n")
			
		 } else {
		     
			 //split datas in lines
			 
			 lines := SplitLines(r)
			 
			 for i:= range(lines) {
			 
			       //split line into fields
				   
			      line := strings.Fields(lines[i])
				  
                  // Check field number read by line
     			  if len(line) != 2 {
				  
				     fmt.Printf("*** Error only two fields expected separated by one space!\n")
					 
				  } else {
				            fn := line[0]
                            ln := line[1]
							
							//check firstand last name size
                            if len(fn) > 20  || len(ln) > 20 {
							   fmt.Printf("*** Error fistname and lastname must be 20 chars !\n")
                            } else {
							        
									s = append(s, name{fname: fn , lname: ln })
                              } 							
				      
                      
				    }
				  
				  			  
			 }
			 
		   }
			
   }
   
   for _, val := range(s) {
       fmt.Printf("%s     %s\n", val.fname, val.lname)
   }
 	    
}