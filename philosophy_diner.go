package main

import(
"log"
"math/rand"
"os"
"sync"
"time"
"hash/fnv"
)


var ph = []string{"Mark","Russel","Rocky","Haris","Root"}

const(
   hunger = 3
   think = time.Second/100
   eat = time.Second/100  
)

var fmt = log.New(os.Stdout,"",0)
var dining sync.WaitGroup

func diningProblem(phName string, dominantHand, otherHand * sync.Mutex){

       fmt.Println(phName, "Seated")
       h := fnv.New64a()
       h.Write( []byte(phName))
       rg := rand.New(rand.NewSource(int64(h.Sum64())))
      
       rSleep := func(t time.Duration) {

            time.Sleep(t/2 + time.Duration(rg.Int63n(int64(t))))
       }

       for h := hunger; h > 0; h-- {

           fmt.Println(phName,"Hungry")
           dominantHand.Lock()  //pick up forks
           otherHand.Lock()  
           fmt.Println(phName, "Eating")
           rSleep(eat)
           dominantHand.Unlock() //put down forks
           otherHand.Unlock()
           fmt.Println(phName,"Thinking")
           rSleep(think)         
       }

       fmt.Println(phName,"Stisfied")
       dining.Done()
       fmt.Println(phName,"Left the table")
        
}


func main() {
    fmt.Println("Table empty")

    dining.Add(5)
    fork0 := &sync.Mutex{}
    forkLeft := fork0
    
     for i := 1;  i < len(ph); i++ {
         forkRight := &sync.Mutex{}
         go diningProblem(ph[i], forkLeft, forkRight)
         forkLeft = forkRight
     }

     go diningProblem(ph[0], fork0, forkLeft)

     dining.Wait()  //wait all philosopher to finish
     fmt.Println("Table empty")

}

