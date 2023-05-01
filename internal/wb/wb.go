package wb


import (
   
   "log" 
   "strings"	 
   "fmt"	
   //"strconv"
   "reflect"	
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
	   
	   paramUrl := r.URL.RequestURI()
	  
	   serviceName := strings.Split(paramUrl,"/")[1] // имя сервиса
	   
	   metricType  := strings.Split(paramUrl,"/")[2] // тип метрики      -- только строка
	   metricName  := strings.Split(paramUrl,"/")[3] // название метрики -- только строка
	   metricValue := strings.Split(paramUrl,"/")[4] // значение метрики -- только не строка 
	   
	   if serviceName != "update" {
		
              http.Error(w, "service not found", http.StatusNotFound)
              return
	   }
	   
	   
	   fmt.Println(reflect.TypeOf(metricType))
	   fmt.Println(reflect.TypeOf(metricName))
	   fmt.Println(reflect.TypeOf(metricValue))
	   
	   
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

   }     
}

func HandleRequests() {

    mux := http.NewServeMux()
    mux.HandleFunc("/", updateCounter)
    
    log.Fatal(http.ListenAndServe(":8080", mux))
}
