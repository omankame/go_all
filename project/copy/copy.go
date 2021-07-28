package main

import (
         "fmt"
//         "io"
         "io/ioutil"
         "os"
         "path/filepath"    
         "syscall"
         
)

var srcfd *os.File

var dstfd *os.File



func main() {
     args := os.Args[1:]
     if len(args) != 2 {
        fmt.Println("Please enter source file name/path and dst path")
        return
     }
         var src, dst string

     	 src = args[0]
     	 dst = args[1]

         fileInfo, err := os.Stat(src)
         if err != nil {
          fmt.Println(err)
         }

         _, err = os.Stat(dst)  
         if err != nil {   // if dirrectory does not exist then it will give error and will create new directory
              if err = CreateIfNotExists(dst, 0755); err != nil {
                 fmt.Println(err)
              }
         }
            

         stat, ok := fileInfo.Sys().(*syscall.Stat_t)   //type assertion topic of interface
         if !ok {
               fmt.Errorf("failed to get raw syscall.Stat_t data for '%s'", src)
         }
         _ = stat
         switch fileInfo.Mode() & os.ModeType {
                case os.ModeDir:
                     CopyDirectory(src, dst )           
                
                  
                default:
                      Copy(src, dst)
          }
}
         
func Copy(src string, dest string) {
     srcf := filepath.Base(src) // for geeting last argument of source file path
     dstf := filepath.Join(dest, srcf)   //absolute path representation else current file     
     fmt.Println(srcf, dstf)
     srcb, err  :=  ioutil.ReadFile(src)  //reading all data as []byte slice
     if err != nil {
        fmt.Println(err)
        return
     }
     
     ioutil.WriteFile(dstf, srcb, 0664) 
     fmt.Println("File copied sucessfully")   
}  


func CopyDirectory(src string, dst string)  {
     entries, err := ioutil.ReadDir(src) // getting all the file and dir info from source directory location
     if err != nil {
        fmt.Println(err)
        return
     } 

     for _, entry := range entries {
         sourcepath := filepath.Join(src, entry.Name())
         destpath   := filepath.Join(dst, entry.Name())
         
         fmt.Println(sourcepath, destpath)
         fileInfo, err := os.Stat(sourcepath) 
         if err != nil {
            fmt.Println(err)
         }

         stat, ok := fileInfo.Sys().(*syscall.Stat_t)   //type assertion topic of interface
         if !ok {
                fmt.Errorf("failed to get raw syscall.Stat_t data for '%s'", sourcepath)
          }

          _ = stat
//          fmt.Println(fileInfo.Mode() & os.ModeType)  //vimp to find out the status "fileInfo.Mode() & os.ModeType"
       
          switch fileInfo.Mode() & os.ModeType {
       
              case os.ModeDir:
                   err := CreateIfNotExists(destpath, 0755)
                   if err != nil {
                      fmt.Println(err)
                   }
 
                  CopyDirectory(sourcepath, destpath)
              
//              case os.ModeSymlink:
//                    err := copySymlink(sourcepath, destpath)
//                    if err != nil {
//                     return err
//                  }

              default:
                  Copy(sourcepath, dst)

                 
            }
      }
}

func CreateIfNotExists(destpath string, perm os.FileMode) error {
     _, err  := os.Stat(destpath)
     if err != nil {
        fmt.Println("Directory is not present, creating directory")
     }
     
    fmt.Println("zzz")
    err = os.Mkdir(destpath, perm) 
          if err != nil {
             return err
          }
    fmt.Println("Directory", destpath, "created sucessfully")
    return nil
} 
