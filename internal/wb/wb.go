package wb


import (
   
   "log" 
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
	   
	   paramsURI := r.URL.RequestURI()
	   
	   checkIncomingParams:
	   
	   for i, v := range strings.Split(paramsURI,"/"){
		   switch i {
		        case 20:
			   // При попытке передать запрос с некорректным типом метрики возвращать http.StatusBadRequest.
			   http.Error(w, "mertic type is incorrect", http.StatusBadRequest)
			   break checkIncomingParams
			case 30:
			   // При попытке передать запрос без имени метрики возвращать http.StatusNotFound.
			   http.Error(w, "mertic name is empty", http.StatusNotFound)
			   break checkIncomingParams
		        case 4:
			   if _, err := strconv.ParseFloat(v, 64); err != nil {
			   // При попытке передать запрос с некорректным значением возвращать http.StatusBadRequest. 
			   w.WriteHeader(http.StatusBadRequest)
		           w.Header().Set("Content-Type", "text/plain")		   
			   break checkIncomingParams
			  }
		        default:
			   w.WriteHeader(http.StatusOK)
	                   w.Header().Set("Content-Type", "text/plain")
		   }	   
		 }
             }     
          }

func HandleRequests() {

    mux := http.NewServeMux()
    mux.HandleFunc("/", updateCounter)
    
    log.Fatal(http.ListenAndServe(":8080", mux))
}
