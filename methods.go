package main

import "fmt"

type Rect struct{
	length int
	breadth int
}

// func Area(length int,breadth int)int{

// 	length*breadth
// }

func(r Rect)Area()int{
	return r.length*r.breadth
}

func main(){
	r:=Rect{3,4}
	Area:=r.Area()
	fmt.Println(Area)
	
}