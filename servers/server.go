package servers

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type Servers struct {
	Ports []int
}

func (s *Servers) addServers(num int) {
	if num > 10 {
		fmt.Printf("Servers can't be more than 10")
	}
	for i := 0; i < num; i++ {
		s.Ports = append(s.Ports, 8000+i)
	}
}

func mkServer(port int, wg sync.WaitGroup){
	defer wg.Done()
	r := mux.NewRouter();
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Server on port %d", port)
	})
	address := fmt.Sprintf(":%d", port)
	http.ListenAndServe(address, r)
	

}

func Runservers() {
	var myServers Servers
	myServers.addServers(6)
	var wg sync.WaitGroup
	wg.Add(6)
	defer wg.Wait()

	for _, port := range myServers.Ports {
		go mkServer(port, wg)
	}	
}
