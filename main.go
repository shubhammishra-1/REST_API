package main

import (

    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"

)


type Payload struct{
    
    End_API string `json: "endapi"`
    Method string  `json: "method"`
    Message string `json: "message"`

}

var endPoint string;


//::CREATE

func createData(w http.ResponseWriter, r *http.Request){

payL:=Payload{
    endPoint+"/create",
    "POST",
    "You hitted create Page!!!",
}

json.NewEncoder(w).Encode(payL)

}

//::UPDATE

func updateData(w http.ResponseWriter, r *http.Request){

payL:=Payload{
    endPoint+"/put",
    "PUT",
    "You hitted PUT Page!!!",
}

json.NewEncoder(w).Encode(payL)


}

//::DELETE

func deleteData(w http.ResponseWriter, r *http.Request){

payL:=Payload{
    endPoint+"/delete",
    "DELETE",
    "You hitted delete Page!!!",
}

json.NewEncoder(w).Encode(payL)
    
}

//::READ

func readData(w http.ResponseWriter, r *http.Request){

payL:=Payload{
    endPoint+"/read",
    "GET",
    "You hitted reading Page!!!",
}

json.NewEncoder(w).Encode(payL)
    
}



func home(w http.ResponseWriter, r *http.Request){

payL:=Payload{
    endPoint+"/home",
    "GET",
    "You hitted homepage!!!",
}

json.NewEncoder(w).Encode(payL)


// jsonData,_:=json.Marshal(payL)  ---> it convert data into json byte

// fmt.Println(string(jsonData))  ---> stringifying the byte json


    
}




func main(){

     
    var PORT string;

    PORT=":9000"

    endPoint+=PORT
    
    router := mux.NewRouter()

    router.HandleFunc("/home",home)

    router.HandleFunc("/create",createData).Methods("POST")
    router.HandleFunc("/put",updateData).Methods("PUT") //Update is same as PUT method
    router.HandleFunc("/read",readData).Methods("GET")
    router.HandleFunc("/delete",deleteData).Methods("DELETE")


    log.Fatal(http.ListenAndServe(PORT,router))



}