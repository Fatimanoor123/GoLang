package main

import (
"fmt"
"errors"

)

func main(){
  err:= errors.New("Error occur")
  if err!=nil{
   fmt.Println(err)
  }


}
