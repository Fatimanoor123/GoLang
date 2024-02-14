package main

import "fmt"

func circumference[r int| float64] (radius r){
   c:= 2*3*radius
   fmt.Println("The circumference is", c)
}

func main(){
  radius1:=45
  radius2:=12.1

  circumference(radius1)
  circumference(radius2)


}
