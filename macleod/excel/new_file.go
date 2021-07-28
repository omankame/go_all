package main

import (
    "fmt"
    "github.com/360EntSecGroup-Skylar/excelize"
)
func main() {

xlsx := excelize.NewFile()

     index := xlsx.NewSheet("Sheet1")
     xlsx.SetCellValue("Sheet1", "A1", "SSH_IP")
     xlsx.SetCellValue("Sheet1", "B1", "USER")
     xlsx.SetCellValue("Sheet1", "C1", "PASSWD")
     xlsx.SetCellValue("Sheet1", "D1", "DEVICE")
     xlsx.SetCellValue("Sheet1", "E1", "DEPENDENTS")
     xlsx.SetCellValue("Sheet1", "F1", "VLAN_IDS")
     xlsx.SetCellValue("Sheet1", "G1", "IPADDR")
     xlsx.SetCellValue("Sheet1", "H1", "NETMASK")
     xlsx.SetCellValue("Sheet1", "I1", "GATEWAY")
     xlsx.SetCellValue("Sheet1", "J1", "IPADDR_V6")
     xlsx.SetCellValue("Sheet1", "K1", "NETMASK_V6")
     xlsx.SetCellValue("Sheet1", "L1", "GATEWAY_V6")
     xlsx.SetCellValue("Sheet1", "M1", "STATUS")
     xlsx.SetCellValue("Sheet1", "N1", "REASON")
     xlsx.SetCellValue("Sheet1", "O1", "REMARKS")


     xlsx.SetActiveSheet(index)

     err := xlsx.SaveAs("./Book1.xlsx")
     if err != nil {
        fmt.Println(err)
     }

}
