package main

import (
        "fmt"
        "net/http"
        "html/template"
        "github.com/360EntSecGroup-Skylar/excelize"
)

var tpl *template.Template

func init() {
     tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

type M map[string]interface{}

func main() {
     http.HandleFunc("/", web)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func web(w http.ResponseWriter, req *http.Request) {
//    aci := ACIDATA{}
    
//    if req.Method == http.MethodPost {
      si := req.FormValue("ssh")
      u := req.FormValue("user")
      p := req.FormValue("passwd")
      dev := req.FormValue("device")
      dep := req.FormValue("dependents")
      vl := req.FormValue("vlan_ids")
      ip := req.FormValue("ipaddr")
      ne := req.FormValue("netmask")
      gw := req.FormValue("gateway")
      ip_6 := req.FormValue("ipaddr_v6")
      ne_6 := req.FormValue("netmask_v6")
      gw_6 := req.FormValue("gateway_v6")
      st := req.FormValue("status")
      rs := req.FormValue("reason")
      re := req.FormValue("remarks")

      fmt.Println(si, u , p, re)
//     tpl.ExecuteTemplate(w, "index.gohtml", acidata{si, u, p, dev, dep, vl, ip, ne, gw, ip_6, ne_6, gw_6, st, rs, re})

     var data = M{"SSH_IP": si, "USER": u, "PASSWD": p, "DEVICE": dev, "DEPENDENTS": dep, "VLAN_IDS": vl, "IPADDR": ip, "NETMASK": ne, "GATEWAY": gw, "IPADDR_V6":  ip_6, "NETMASK_V6": ne_6, "GATEWAY_V6": gw_6, "STATUS": st, "REASON": rs, "REMARKS": re  }
    
    
     tpl.ExecuteTemplate(w, "index.gohtml", data)

     f, err := excelize.OpenFile("./Book1.xlsx")
     if err != nil {
     fmt.Println(err)
     return
     }


    sheetName := f.GetSheetName(0)
    fmt.Println(sheetName)
    rows, err := f.GetRows(sheetName)
    rlast := len(rows)
    f.SetCellValue(sheetName, fmt.Sprintf("A%d", rlast+1), data["SSH_IP"])
    f.SetCellValue(sheetName, fmt.Sprintf("B%d", rlast+1), data["USER"])
    f.SetCellValue(sheetName, fmt.Sprintf("C%d", rlast+1), data["PASSWD"])
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

     err = f.SaveAs("./Book1.xlsx")
     if err != nil {
        fmt.Println(err)
     }

     
}


