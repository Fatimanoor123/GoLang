package main

import "fmt"

type adress struct{
   City, state string

}

type Person struct{
  name string
  adress //embeded function
}

func main(){
   p:=Person{
      name: "Ismail",
      adress: adress{
      City: "Wah",
      state: "Punjab",
   },
}
  fmt.Println(p.name)
  fmt.Println(p.adress)
}
