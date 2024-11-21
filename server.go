package main
import ("fmt"
        "net/http")

func main(){
    fmt.Println("go server example code");
    var fs= http.FileServer(http.Dir("./dist"));
    http.Handle("/",fs);

    err := http.ListenAndServe(":8080",nil);

    if err!= nil{
        fmt.Printf("Problem starting server %s",err);
    }

}