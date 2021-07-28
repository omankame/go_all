package main

import (
        "fmt"
        "bufio"
        "os" 
        "strconv"
        "log"
        "strings"
)

type Node struct {
     data	int
     next	*Node
}

type linkList struct {
     head       *Node
     len	int
}

func initLink() *linkList {
     return &linkList{}
}

func main() { 

     l := initLink()
     reader := bufio.NewReader(os.Stdin)
     fmt.Println("How many nodes do you want?")
     num_str, err := reader.ReadString('\n')
     if err != nil {
        log.Fatalln(err)
     }
     num_str = strings.TrimSuffix(num_str, "\n") 
     num, err := strconv.Atoi(num_str) 
     if err != nil {
        log.Fatalln(err)
     }

     for i := 0; i < num ; i++ {
       reader = bufio.NewReader(os.Stdin)
       val_str, _  :=  reader.ReadString('\n')
       val_str = strings.TrimSuffix(val_str, "\n")
       val, err := strconv.Atoi(val_str)
        if err != nil {
        log.Fatalln(err)
       }
  
        l.addVal(val)
    } 

    l.showLink()
//    beg := 100
    l.addFront(100)

    l.showLink()
    l.addEnd(1000)
    l.showLink()
    l.count()
}

func (li *linkList)addVal(val int) {

     n := &Node{}
     n.data = val
     if li.len == 0 {
             li.head = n
             li.len++
       fmt.Println("First Node Added")
     } else {
        curNod := li.head
        for curNod.next != nil {
            curNod = curNod.next
           }
      curNod.next = n
      li.len++
      fmt.Println("Second and onwards Node Added") 
}

}
func (li *linkList)addFront(val int) {
      n := &Node{}
      n.data = val
      if li.len == 0 {
            li.head = n
            li.len++
       fmt.Println("New Node Added ")
       } else {
            n.next = li.head
            li.head = n
            li.len++
       fmt.Println("New Node Added ")       
       }
}  

func (li  *linkList)addEnd(val int) {
     n := &Node{}
     n.data = val
     if li.len == 0 {
            li.head = n
            li.len++
       fmt.Println("New Node Added ")
       } else {
         curNod := li.head
        for curNod.next != nil {
            curNod = curNod.next
           }
      curNod.next = n
      li.len++
      fmt.Println("node added at the end")
}

}

func (li *linkList)showLink() {
      curnod := li.head
      if curnod == nil {
        fmt.Println("Link List is empty")
      }  
        fmt.Println("Final Link list values are") 
        for i :=0; i < li.len; i++ {
        fmt.Printf(" %v", curnod.data)
        curnod = curnod.next
     }
     fmt.Println()  

}

func (li *linkList)count() {
     fmt.Printf("Total no. of nodes %d \n", li.len) 

}
