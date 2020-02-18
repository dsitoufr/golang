package main

import (
"fmt"
"strconv"
)

func sortSlice( s10 []int) []int{
 var N int = len(s10)
 for i := N - 1; i >= 0; i-- {
    for j := 0; j <= i - 1; j++ {
	   if s10[j] > s10[j+1] {
           v := s10[j]
		   w := s10[j + 1]
		   
		   s10[j] = w
		   s10[j + 1] = v
	   }
	      
	}
 }
 return s10
}




func main() {
   
  var Str string
  var StrInt int
  const SlSize = 3
  s0 := make([]int, SlSize)
  

  i := 0
   
   for {
      fmt.Printf("Enter a number:")
	  fmt.Scan(&Str)
	  
	  _, err := strconv.Atoi(Str)
	  if err != nil {

     	  //not integer
		  if Str == "X" {
		     break
		  }
	  } else {
	     //integer
		 StrInt, _ = strconv.Atoi(Str)
		 if i < SlSize {
		 
		    if i == (SlSize - 1) {
                 s0[0] = StrInt
			}  else {
                 s0[i] = StrInt
				}
			i++
		 } else {
		        
		        s0 = append(s0, StrInt)
				}  
           k := sortSlice(s0)
           fmt.Printf("tried: %v\n", k) 		   
		 
	  }
   }

}