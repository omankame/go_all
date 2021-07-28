package main

import (
    "fmt"
    "os"
    "github.com/360EntSecGroup-Skylar/excelize"
//    "strconv"
)

type ACI struct {
     ssh_ip	string
     user	string
     password	string
     device	string
     dependents string
     vlan_ids   string
     ipaddr 	string
     netmask	string
     gateway	string
     ipaddr_v6	string
     netmask_v6 string
     gateway_v6 string
     status	string
     reason	string
     remarks    string
}



func main() {

     args := os.Args[1:]
     if len(args) != 15 {
     fmt.Println("Please check the data you entered")
     return
     } 

     aci := ACI{
     ssh_ip:   args[0],
     user:     args[1],
     password: args[2],
     device:   args[3],
     dependents: args[4],
     vlan_ids: args[5],
     ipaddr:   args[6],
     netmask:  args[7],
     gateway:  args[8],
     ipaddr_v6:  args[9],
     netmask_v6: args[10],
     gateway_v6: args[11],
     status:     args[12],
     reason:     args[13],
     remarks:    args[14], 
     }
     f, err := excelize.OpenFile("./Book1.xlsx")
     if err != nil {
        fmt.Println(err)
        return
     }

    sheetName := f.GetSheetName(1)
    rows, err := f.GetRows(sheetName)
    rlast := len(rows)
    fmt.Println(rlast)

    f.SetCellValue(sheetName, fmt.Sprintf("A%d", rlast+1), aci.ssh_ip)
    f.SetCellValue(sheetName, fmt.Sprintf("A%d", rlast+1), aci.user)
    f.SetCellValue(sheetName, fmt.Sprintf("A%d", rlast+1), aci.password)
    f.SetCellValue(sheetName, fmt.Sprintf("A%d", rlast+1), aci.device)
    f.SetCellValue(sheetName, fmt.Sprintf("A%d", rlast+1), aci.dependents)
    f.SetCellValue(sheetName, fmt.Sprintf("A%d", rlast+1), aci.vlan_ids)
    f.SetCellValue(sheetName, fmt.Sprintf("A%d", rlast+1), aci.ipaddr)
    f.SetCellValue(sheetName, fmt.Sprintf("A%d", rlast+1), aci.netmask)
    f.SetCellValue(sheetName, fmt.Sprintf("A%d", rlast+1), aci.gateway)
    f.SetCellValue(sheetName, fmt.Sprintf("A%d", rlast+1), aci.ssh_ip)
    f.SetCellValue(sheetName, fmt.Sprintf("A%d", rlast+1), aci.ipaddr_v6)
    f.SetCellValue(sheetName, fmt.Sprintf("A%d", rlast+1), aci.netmask_v6)
    f.SetCellValue(sheetName, fmt.Sprintf("A%d", rlast+1), aci.gateway_v6)
    f.SetCellValue(sheetName, fmt.Sprintf("A%d", rlast+1), aci.status)
    f.SetCellValue(sheetName, fmt.Sprintf("A%d", rlast+1), aci.reason)
    f.SetCellValue(sheetName, fmt.Sprintf("A%d", rlast+1), aci.remarks)

    err = f.SaveAs("./Book1.xlsx")
     if err != nil {
        fmt.Println(err)
     }


}
