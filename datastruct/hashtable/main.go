//https://dev.to/ishankhare07/you-think-you-understand-key-value-pairs-m7l

package main

import (
        "fmt"
)

const MAP_SIZE = 50


type Node struct {
     key   string
     value string
     next  *Node
}

type HashMap struct {
     Data []*Node
}


func NewDict() *HashMap {
     return &HashMap{ Data: make([]*Node, MAP_SIZE) }
}

func main() {
     a := NewDict()

     a.Insert("name", "ishan")
     a.Insert("gender", "male")
     a.Insert("city", "mumbai")
     a.Insert("lastname", "khare") 

     if value, ok := a.Get("name"); ok {
       fmt.Println(value);
     } else {
       fmt.Println("Value did not match!")
    }

     fmt.Println("%q", a)
}    

func hash(key string) (hash uint32) {
    hash = 0
    for _, ch := range key {
        hash += uint32(ch)
        hash += hash << 10
        hash ^= hash >> 6
    }

    hash += hash << 3
    hash ^= hash >> 11
    hash += hash << 15

    return
}

func getIndex(key string) (index int) {
     return int(hash((key)))  % MAP_SIZE
}

func (h *HashMap) Insert( key string, value string) {
     index := getIndex(key) 

     if h.Data[index] == nil {
          h.Data[index] = &Node{key: key, value: value}
     } else  {
       starting_node := h.Data[index]

       for ; starting_node != nil; starting_node = starting_node.next {
             if starting_node.key == key {
                starting_node.value = value
                return
             }
       }

       starting_node.next = &Node{key: key, value: value }
     }

  }   

func (h *HashMap) Get( key string ) (string, bool) {
     index := getIndex(key)

     if h.Data[index] != nil {
        starting_node := h.Data[index]
      
        for ; ; starting_node = starting_node.next {
              if starting_node.key == key {
                 return starting_node.value , true
              }
         	if starting_node.next == nil {
            	break
              }
         }
      }

     return "", false
} 

         
