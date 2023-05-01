package rq


import (
   
   "log"
   "strconv"  
   "net/http"
)


type memStorage struct (
  
     store map[string][int64]
)


func updateCounter(w http.ResponseWriter, r *http.Request){
     
   if r.Method != http.MethodPost {
      
      return
   
   } else {
   
     tp, err := r.URL.Query().Get("tp") // param type
     nm, err := r.URL.Query().Get("nm") // param name
     vl,err  := r.URL.Query().Get("vl") // param value
     
     
     // При успешном приёме возвращать http.StatusOK.

     // При попытке передать запрос без имени метрики возвращать http.StatusNotFound.

     //При попытке передать запрос с некорректным типом метрики или значением возвращать http.StatusBadRequest.
 
   }     
}

func HandleRequests() {

    mux := http.NewServeMux()
    mux.HandleFunc('/update/{tp}/{nm}/{vl}', updateCounter)
    
    log.Fatal(http.ListenAndServe(":8080", mux))
}
