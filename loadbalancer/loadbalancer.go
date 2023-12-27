package loadbalancer

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"
)

var (baseUrl = "http://localhost:800")
type Targets struct{
	servers []*url.URL
	index int
}

func Loadbalancer(){

	r := mux.NewRouter()
	var tg Targets
	tg.index = 0;
	for i:=0; i<6;i++{
		tg.servers = append(tg.servers, parseUrl(baseUrl,i))
	}
	r.HandleFunc("/loadbalancer",proxyRequests(&tg))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Server on port 8080")
	})
	log.Fatal(http.ListenAndServe(":8080", r))

	
}
func proxyRequests( sr *Targets) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		sr.index = sr.index % len(sr.servers)
		rvproxy := httputil.NewSingleHostReverseProxy(sr.servers[sr.index])
		fmt.Println(sr.servers[sr.index])
		fmt.Println(sr.index)
		sr.index ++;
		rvproxy.ServeHTTP(w,r)

	}
	
}

func parseUrl(baseUrl string, i int) *url.URL{
	link := baseUrl + strconv.Itoa(i)
	parsedURL, _ := url.Parse(link)
    return parsedURL

}
