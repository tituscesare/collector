package rq


import (
     
   "net/http"
)


type memStorage struct (
  
     store map[string][int64]
)


func updateCounter(w http.ResponseWriter, r *http.Request){
  
    
}

func HandleRequests() {

    mux := http.NewServeMux()
    mux.HandleFunc('/update/{type}/{name}/{value}', updateCounter).Methods("POST")
    
    log.Fatal(http.ListenAndServe(":8080", mux))
}
