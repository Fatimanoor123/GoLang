package main 

import "fmt"

func sum(nums...int){
    fmt.Println("numbers are")
    total:=0
    for _, num:= range  nums{
        total += num        
    }
      fmt.Println("The sum of numbers are", total)
}    

func main(){
    
    sum(1,2)
    sum(12,21)
    
    
}
