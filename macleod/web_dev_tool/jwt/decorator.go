/* decorator is important concept which is used in JWT. Its function wrapping mechanism

func main() {

    fmt.Printf("Type of myfunc%T", myfunc)
    coolfunc(myfunc)
}
func myfunc() {
    fmt.Println("Hello in Myfunc")
}

func coolfunc(a func()) {
     fmt.Printf("Start of myfunc as a()")
     a()
     fmt.Printf("End of myfunc as a()")
}

the above coolfunc is concept of decorators */

/* now will see real life decorators

func main() {

     http.Handle("/", isauthorized(homepage))
     http.ListenANdServe(":8080", nil)
}

func homepage(w http.ResponseWriter, req *http.Request) {

}

func isauthorized(endpoint func(http.ResponseWriter, *http.Request) http.Handler {
     return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {   ==> real trick as handler return by http.HandlerFunc(f) hence took here
      
      logic
     })

    another logic
  })

}
