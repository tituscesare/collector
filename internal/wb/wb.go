package wb


import (
   
   "fmt"
   "log" 
   "strings"	
   "net/http"
)

/*
type memStorage struct {
  
     store map[string][int64],
}
*/

func updateCounter(w http.ResponseWriter, r *http.Request){
     
   if r.Method != http.MethodPost {
      
      http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
      return
   
   } else {
	   
	   paramUrl := r.URL.RequestURI()
	   
	   
   /*
     tp := r.URL // param type
     
      // При попытке передать запрос с некорректным типом метрики возвращать http.StatusBadRequest. 
      
      if err != nil {
         
         http.Error(w, "mertic type is incorrect", http.StatusBadRequest)
         
         return
      }
      
      nm, err := r.URL.Query().Get("nm") // param name
      
      // При попытке передать запрос без имени метрики возвращать http.StatusNotFound.
      if err != nil {
         
         http.Error(w, "mertic name is empty", http.StatusNotFound)
         
         return
      }
      
     vl, err := r.URL.Query().Get("vl") // param value
     
      // При попытке передать запрос с некорректным значением возвращать http.StatusBadRequest. 
     if err != nil {
         
         http.Error(w, "mertic value is incorrect", http.StatusBadRequest)
         
         return
      }
      
     // При успешном приёме возвращать http.StatusOK.
     */ 
      	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprintf(w, string(strings.Split(paramUrl,"/"))) 
         
   }     
}

func HandleRequests() {

    mux := http.NewServeMux()
    mux.HandleFunc("/", updateCounter)
    
    log.Fatal(http.ListenAndServe(":8080", mux))
}
