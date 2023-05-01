package wb


import (
   
   "log" 
   "strings"	
   "strconv"
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
	  
	   serviceName := strings.Split(paramUrl,"/")[1] // имя сервиса
	   
	   metricType  := strings.Split(paramUrl,"/")[2] // тип метрики
	   metricName  := strings.Split(paramUrl,"/")[3] // название метрики
	   metricValue := strings.Split(paramUrl,"/")[4] // значение метрики
	   
	   if serviceName != "update" {
		
              http.Error(w, "service not found", http.StatusNotFound)
              return
	   }
	   
	   // При попытке передать запрос с некорректным типом метрики возвращать http.StatusBadRequest. 
      
      if metricType == " " {
         
         http.Error(w, "mertic type is incorrect", http.StatusBadRequest)
         
         return
      }
	   // При попытке передать запрос без имени метрики возвращать http.StatusNotFound.
           if metricName == " " {
         
            http.Error(w, "mertic name is empty", http.StatusNotFound)
         
           return
        }
	   // При попытке передать запрос с некорректным значением возвращать http.StatusBadRequest. 
        if metricValue == " " {
		
         http.Error(w, "mertic value is incorrect", http.StatusBadRequest)
         
         return
	}
      
     // При успешном приёме возвращать http.StatusOK.
      	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
        
	//fmt.Fprintf(w, 
   }     
}

func HandleRequests() {

    mux := http.NewServeMux()
    mux.HandleFunc("/", updateCounter)
    
    log.Fatal(http.ListenAndServe(":8080", mux))
}
