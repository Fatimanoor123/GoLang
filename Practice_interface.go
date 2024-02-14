package main 

import (
   "fmt"
   "math"

)
type geometry interface {
   area() float64
   perim() float64

}

type rect struct{
  height, width float64
}
 
type circle struct {
   radius float64

}

func (r rect)area() float64{
	return r.width*r.height
}
func (r rect) perim() float64{
   return 2 * r.height + r.width

} 

func (cir circle) area() float64{
  return 2*math.Pi* cir.radius*cir.radius
}
func (cir circle) perim() float64{
  return 2*math.Pi* cir.radius
}
func measure(g geometry)  {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func main(){
  r:= rect{width:3, height: 6}
  cir:=circle{radius:4}  
  measure(r)
measure(cir) 
}
