package main

import (
         "net"
         "log"

         "google.golang.org/grpc"
)

func main() {
     l, err := net.Listen("tcp", ":8080")
     if err != nil {
        log.Fatal(err)
     }

     gRPCServer := grpc.NewServer()

     if err := gRPCServer.Serve(l); err != nil {
        log.Fatal(err)
     }

     log.Println("Minimum GRPC server implemented")    
}      
