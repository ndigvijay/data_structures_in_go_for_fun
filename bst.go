package main

import "fmt"

type Node struct {
    data  int
    left  *Node
    right *Node
}

func (N *Node) insert(number int) {
    temp := &Node{number, nil, nil}
    if N.left == nil && N.right == nil {
        if number < N.data {
            N.left = temp
            return
        }
        if number > N.data {
            N.right = temp
            return
        }
    }
    curr := N
    var parent *Node
    for curr != nil {
        parent = curr
        if curr.data > number {
            curr = curr.left
        } else if curr.data < number {
            curr = curr.right
        } else {
            return
        }
    }
    if parent.data > number {
        parent.left = temp
    } else {
        parent.right = temp
    }
}

func (N *Node) display() {
    if N == nil {
        return
    }
    N.left.display()
    fmt.Println(N.data)
    N.right.display()
}

func main() {
    head := &Node{0, nil, nil}
    head.insert(7)
    head.insert(3)
    head.insert(10)
    head.insert(1)
    head.insert(5)
    head.insert(9)
    head.insert(8)
    head.display()
}