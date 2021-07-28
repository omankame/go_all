package main

import (
         "net/http"
         "fmt"
         "html/template" 
         "os"
         "log"
)

var tpl *template.Template

func init() {
     tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
 
}

func main() {

     http.HandleFunc("/", indexPage)
     http.HandleFunc("/filecreate", filecreate)
     http.HandleFunc("/routeadd", routeadd)
     http.HandleFunc("/routeaddIP6", routeaddIP6)
     http.HandleFunc("/validafail", validafail)
     http.ListenAndServe(":8080", nil)

}


func indexPage(w http.ResponseWriter, req *http.Request) {
     
    tpl.ExecuteTemplate(w , "index.gohtml", nil)
    fmt.Println("Index page open")

}

func filecreate(w http.ResponseWriter, req *http.Request) { 


     if req.Method == http.MethodPost {
        fn := req.FormValue("file")
       
        if alreadyFileExists(fn) {
           http.Redirect(w, req, "/", http.StatusSeeOther)
           return
        }

        rFile, err := os.Create(fn)
        if err != nil {
           log.Fatal(err)
           return
        }

        log.Println(rFile)
        rFile.Close()

        http.Redirect(w, req, "/", http.StatusSeeOther)
 
        return
     } 
       
     tpl.ExecuteTemplate(w, "filecreate.gohtml", nil)

}

func routeadd(w http.ResponseWriter, req *http.Request) {

     if req.Method == http.MethodPost {
        dstIp := req.FormValue("dst")
        gw    := req.FormValue("gw")
        mask  := req.FormValue("mask")
        intf  := req.FormValue("intf")

        value := checkIPAddress(dstIp)
        if value == false {
           http.Redirect(w, req, "/validafail", http.StatusSeeOther)
           w.Write([]byte("<script>alert('Invalid IP Address')</script>")) 
        }
//        checkIPAddress(gw)

        

       file, err := os.OpenFile("route.text", os.O_APPEND|os.O_WRONLY, 0644)
        if err != nil {
           fmt.Println(nil)
           http.Redirect(w, req, "/", http.StatusSeeOther)
           return
        }

        defer file.Close()
        
        rt := fmt.Sprintf("ip route add %s%s via %s dev %s \n", dstIp, mask, gw, intf) 
        if _, err := file.WriteString(rt); err != nil {
              fmt.Println(err)
              http.Redirect(w, req, "/", http.StatusSeeOther)
        }

      }
 
        tpl.ExecuteTemplate(w, "routeadd.gohtml", nil)
}


func routeaddIP6(w http.ResponseWriter, req *http.Request) {


      if req.Method == http.MethodPost {
        dstIp := req.FormValue("dst")
        gw    := req.FormValue("gw")
        mask  := req.FormValue("mask")
        intf  := req.FormValue("intf")

        value := checkIPAddress(dstIp)
        if value == false {
           http.Redirect(w, req, "/validafail", http.StatusSeeOther)
           w.Write([]byte("<script>alert('Invalid IP Address')</script>"))
        }
//        checkIPAddress(gw)



       file, err := os.OpenFile("route.text", os.O_APPEND|os.O_WRONLY, 0644)
        if err != nil {
           fmt.Println(nil)
           http.Redirect(w, req, "/", http.StatusSeeOther)
           return
        }

        defer file.Close()

       
        rt := fmt.Sprintf("ip -6 route add %s%s via %s dev %s \n", dstIp, mask, gw, intf)
        if _, err := file.WriteString(rt); err != nil {
              fmt.Println(err)
              http.Redirect(w, req, "/", http.StatusSeeOther)
        }

      }

        tpl.ExecuteTemplate(w, "routeaddIPV6.gohtml", nil)
}

func validafail(w http.ResponseWriter, req *http.Request) {
     Str := "Invalid IP Address" 
     tpl.ExecuteTemplate(w, "validfail.gohtml", Str)
}

