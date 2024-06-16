package main

import "fmt"

type Node struct{
	data int
	next *Node
}


func(N *Node)insert_Node(number int){
	temp:=&Node{data:number,next:nil}
	if N.next==nil{
		N.next=temp
		return
	}
	curr:=N
	for ;curr.next!=nil;curr=curr.next{
		continue
	}
	curr.next=temp

}


func (N *Node)display(){
	curr:=N.next
	for curr!=nil{
		fmt.Println(curr.data)
		curr=curr.next
	}
}

func main(){
	head:=&Node{0,nil}
	head.insert_Node(7)
	head.insert_Node(8)
	head.insert_Node(9)
	head.insert_Node(10)
	head.display()
}