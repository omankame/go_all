package main

import (
//        "fmt"
)

type Node struct {
     dat    int
     next   *Node
}

type Linkedlist struct {
     head 	*Node
     len	int
}

func initList() *Linkedlist {

     return &Linkedlist{}
}


func main() {

     l := initList()

     l.Insert(5)
     l.showLink()
}

func (li *Linkedlist)Insert(val int) {

     n := &Node {}
     n.dat = val
     if li.len == 0 {
        li.head = n
        li.len++
        return 
     } else {
        curnod := li.head
        for curnod.next != nil {
               curnod = curnod.next
          }

        curnod.next = n
        li.len++
        return
   }
     
/* func (li *Linkedlist)showLink() { 
     curnod := li.head
     if curnod == nil {
        fmt.Println("Link List is empty")
     }

     for curnod.next != nil {
        fmt.Printf("%+v\n", curnod)
        curnod = curnod.next
     }   


}   */         
