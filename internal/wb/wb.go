package wb


import (
   
   "log" 
   //"fmt"
   "strings"	 
   "strconv"
   "net/http"
	
)

//Тип gauge, float64 — новое значение должно замещать предыдущее.
type gauge float64

//Тип counter, int64 — новое значение должно добавляться к предыдущему, если какое-то значение уже было известно серверу.
type counter int64

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
	   
	   paramsURI := strings.Split(r.URL.RequestURI(),"/")
	   
	   checkIncomingParams:
	   
	   for i, v := range paramsURI {
		   switch i {
		        case 20:
			   if v != "counter" || v != "gauge" {
			   // При попытке передать запрос с некорректным типом метрики возвращать http.StatusBadRequest.
			   http.Error(w, "mertic type is incorrect", http.StatusBadRequest)
			   break checkIncomingParams
			 }
			case 30:
			   if len(v) == 0 {
			   // При попытке передать запрос без имени метрики возвращать http.StatusNotFound.
			   http.Error(w, "mertic name is empty", http.StatusNotFound)
			   break checkIncomingParams
			  }
		        case 4:
			  
			  if v == "gauge" {
			   if _, err := strconv.ParseFloat(v, 64); err != nil {
			   // При попытке передать запрос с некорректным значением возвращать http.StatusBadRequest. 
		           http.Error(w, "mertic value is incorrect", http.StatusBadRequest) 
		           break checkIncomingParams
			  }
			 }
			if v == "counter" {
			   if _, err := strconv.ParseInt(v, 10, 64); err != nil {
			   // При попытке передать запрос с некорректным значением возвращать http.StatusBadRequest. 
		           http.Error(w, "mertic value is incorrect", http.StatusBadRequest) 
		           break checkIncomingParams
			  }
			 }
		      //  default:
			//   w.WriteHeader(http.StatusOK)
	              //     w.Header().Set("Content-Type", "text/plain")
		   }	   
		 }
             }     
          }

func HandleRequests() {

    mux := http.NewServeMux()
    mux.HandleFunc("/", updateCounter)
    
    log.Fatal(http.ListenAndServe(":8080", mux))
}
