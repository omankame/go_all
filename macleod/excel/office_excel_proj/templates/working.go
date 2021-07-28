package main

import (
//        "fmt"
        "net/http"
        "html/template"
)

var tpl *template.Template

func init() {
     tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

type acidata struct {
     Ssh_ip	string
     User	string
     Passwd	string
     Device	string
     Dependents	string
     Vlan_ids	string
     Ipaddr	string
     Netmask	string
     Gateway	string
     Ipaddr_v6	string
     Netmask_v6	string
     Gateway_v6	string
     Status	string
     Reason	string
     Remarks	string
}


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

     tpl.ExecuteTemplate(w, "index.gohtml", acidata{si, u, p, dev, dep, vl, ip, ne, gw, ip_6, ne_6, gw_6, st, rs, re})
    
}

