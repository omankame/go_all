/* built in data type, with key value pairs. ex. make(map[type of key]type of value) make(map[string]int) empsal := make(map[string]int)
zero value of map is nil, retriving value from key val := empsal["onkar"] 
checking key exists is very imp.  v, ok := map[key]  , you can execute for loop on map but value retrieval order is not fixed.
delete the map delete(mapname, key) ex. delete(empsal, "onkar") maps of struct ex below
length of map ==> len(empsal) , map are reference type like slice so map assign to another variable they both point to same internal data structure.  maps only compare with nil. another method is compare each element of map with evry other element of map. Or user REFLECTION. */

package main 

import (
         "fmt"
)

type employee struct {
      country	string
      salary    int
}

func main() {
      emp1 := employee{"INDIA", 20000}
      emp2 := employee{"USA", 5000}

      empinfo := make(map[string]employee)

      empinfo["onkar"] = emp1
      empinfo["mike"] = emp2

      for name, info := range empinfo {
         fmt.Printf("Employee Name %s, Employee Country %v, Employee Salary %v\n", name, info.country, info.salary)
      }


} 
