package rq


import (
     
   "net/http"
)


type memStorage struct (
  
     store map[string][int64]
)


func updateCounter(w http.ResponseWriter, r *http.Request){
      
     params := mux.Vars(r)
     
     paramType  := params["tp"] // param type
     paramName  := params["nm"] // param name
     paramValue := params["vl"] // param value
     
     
     // При успешном приёме возвращать http.StatusOK.

     // При попытке передать запрос без имени метрики возвращать http.StatusNotFound.

     //При попытке передать запрос с некорректным типом метрики или значением возвращать http.StatusBadRequest.

    
}

func HandleRequests() {

    mux := http.NewServeMux()
    mux.HandleFunc('/update/{tp}/{nm}/{vl}', updateCounter).Methods("POST")
    
    log.Fatal(http.ListenAndServe(":8080", mux))
}
