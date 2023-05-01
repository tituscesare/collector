package wb


import (
   
   "log" 
   //"fmt"
   "strings"	 
   "strconv"
   "net/http"
	
)

type MemStorage struct {
     gauge map[string]float64  //Тип gauge, float64 — новое значение должно замещать предыдущее.
     counter map[string]int64  //Тип counter, int64 — новое значение должно добавляться к предыдущему, если какое-то значение уже было известно серверу.
}


func updateMetrics(w http.ResponseWriter, r *http.Request){
     
   if r.Method != http.MethodPost {
      
      http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
      return
   
   } else {
	   
	   paramsURI := strings.Split(r.URL.RequestURI(),"/")
	   
	   for i, v := range paramsURI {
		   switch i {
			   
		        case 2:
			   if (v != "counter") && (v != "gauge") {
			   // При попытке передать запрос с некорректным типом метрики возвращать http.StatusBadRequest.
			   //http.Error(w, "mertic type is incorrect", http.StatusBadRequest)
		           w.WriteHeader(http.StatusBadRequest)
		           return
			 }
			   
			case 3:
			   if len(v) == 0 {
			   // При попытке передать запрос без имени метрики возвращать http.StatusNotFound.
			   http.Error(w, "mertic name is empty", http.StatusNotFound)
			  }
		        case 4:
			   if _, err := strconv.ParseFloat(v, 64); err != nil {
			   // При попытке передать запрос с некорректным значением возвращать http.StatusBadRequest. 
		           http.Error(w, "mertic value is incorrect", http.StatusBadRequest) 
		           return		   
			  }
			   if _, err := strconv.ParseInt(v, 10, 64); err != nil {
			   // При попытке передать запрос с некорректным значением возвращать http.StatusBadRequest. 
		           http.Error(w, "mertic value is incorrect", http.StatusBadRequest) 
		           return
			  }
		   }	   
		 }
             }     
          }

func HandleRequests() {

    mux := http.NewServeMux()
    mux.HandleFunc("/", updateMetrics)
    
    log.Fatal(http.ListenAndServe(":8080", mux))
}
