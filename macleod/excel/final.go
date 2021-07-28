package main

import (
    "fmt"
    "os"
    "github.com/360EntSecGroup-Skylar/excelize"
//    "strconv"
//     "strings"
)


type M map[string]interface{}



func main() {

     args := os.Args[1:]
     if len(args) != 15 {
     fmt.Println("Please check the data you entered")
     return
     } 

//    sl_DEPENDENTS := strings.Split(args[4], "\n") 

     var data = M{"SSH_IP": args[0], "User": args[1], "Password": args[2], "DEVICE":args[3], "DEPENDENTS": args[4], "VLAN_IDS": args[5], "IPADDR": args[6], "NETMASK": args[7], "GATEWAY": args[8], "IPADDR_V6":  args[9], "NETMASK_V6": args[10], "GATEWAY_V6": args[11], "STATUS": args[12], "REASON": args[13], "REMARKS": args[14]  }

     f, err := excelize.OpenFile("./Book1.xlsx")
     if err != nil {
        fmt.Println(err)
        return
     }

    sheetName := f.GetSheetName(0)
    fmt.Println(sheetName)
    rows, err := f.GetRows(sheetName)
    rlast := len(rows)
    fmt.Println(rlast)

    f.SetCellValue(sheetName, fmt.Sprintf("A%d", rlast+1), data["SSH_IP"])
    f.SetCellValue(sheetName, fmt.Sprintf("B%d", rlast+1), data["User"])
    f.SetCellValue(sheetName, fmt.Sprintf("C%d", rlast+1), data["Password"])
    f.SetCellValue(sheetName, fmt.Sprintf("D%d", rlast+1), data["DEVICE"])
    f.SetCellValue(sheetName, fmt.Sprintf("E%d", rlast+1), data["DEPENDENTS"])
    f.SetCellValue(sheetName, fmt.Sprintf("F%d", rlast+1), data["VLAN_IDS"])
    f.SetCellValue(sheetName, fmt.Sprintf("G%d", rlast+1), data["IPADDR"])
    f.SetCellValue(sheetName, fmt.Sprintf("H%d", rlast+1), data["NETMASK"])
    f.SetCellValue(sheetName, fmt.Sprintf("I%d", rlast+1), data["GATEWAY"])
    f.SetCellValue(sheetName, fmt.Sprintf("J%d", rlast+1), data["IPADDR_V6"])
    f.SetCellValue(sheetName, fmt.Sprintf("K%d", rlast+1), data["NETMASK_V6"])
    f.SetCellValue(sheetName, fmt.Sprintf("L%d", rlast+1), data["GATEWAY_V6"])
    f.SetCellValue(sheetName, fmt.Sprintf("M%d", rlast+1), data["STATUS"])
    f.SetCellValue(sheetName, fmt.Sprintf("N%d", rlast+1), data["REASON"])
    f.SetCellValue(sheetName, fmt.Sprintf("O%d", rlast+1), data["REMARKS"])
//    f.SetCellValue(sheetName, fmt.Sprintf("P%d", rlast+1), data["remarks"])

    err = f.SaveAs("./Book1.xlsx")
     if err != nil {
        fmt.Println(err)
     }


}
