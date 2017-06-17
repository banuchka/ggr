package main

import (
        "io/ioutil"
        "log"
        "net/http"
        "strings"
        "time"
        "fmt"
        "github.com/antonholmquist/jason"
//	"github.com/patrickmn/go-cache"
)
	var (
    		httpClientChecker *http.Client
	)

	const (
    		MaxIdleConnections int = 1000
    		RequestTimeout     int = 5
	)

	// init HTTPClient
	func init() {
    		httpClientChecker = createHTTPClient()
	}

	// createHTTPClient for connection re-use
	func createHTTPClient() *http.Client {
    		client := &http.Client{
        		Transport: &http.Transport{
            				MaxIdleConnsPerHost: MaxIdleConnections,
        		},
        	Timeout: time.Duration(RequestTimeout) * time.Second,
    	}

    	return client
	}
func check(h *Host, b string) bool {
        var browser = strings.ToUpper(b)
        var nodestate bool
	//key_queuesize := fmt.Sprintf("%s_%d_%s_queue",h.Name, h.Port, browser)
	//key_freeslots := fmt.Sprintf("%s_%d_%s_free",h.Name, h.Port, browser)
	//key_unavailable := fmt.Sprintf("%s_%d_%s_unavailable",h.Name, h.Port, browser)
        
	
	//fkey_queuesize, found := c.Get(key_queuesize)
	//if found {}
	//log.Printf("[DEBUG] queue size in cache %d for %s:%d and %s", fkey_queuesize, h.Name, h.Port, browser)
	//fkey_freeslots, found := c.Get(key_freeslots)
        //if found {}
 
	//if fkey_freeslots == nil || fkey_queuesize == nil {
        //fkey_unavailable, found := c.Get(key_unavailable) 
	//if x, found := c.Get(key_unavailable);found {
         //       fkey_unavailable = x.(int)+1
 	//	c.Set(key_unavailable, fkey_unavailable, cache.DefaultExpiration)	
	//	nodestate := false
         //       return nodestate
        //} else {
		url := fmt.Sprintf("http://%s:%d/grid/api/hub", h.Name, h.Port)
        	request, err := http.NewRequest("GET", url, strings.NewReader("{\"configuration\":[\"browserSlotsCount\", \"newSessionRequestCount\"]}"))
 		resp, err := httpClientChecker.Do(request)
        	if err != nil {
	//		c.Set(key_unavailable, 1, cache.DefaultExpiration)
                	nodestate := false
                	log.Printf("[DEBUG] failed request #1 %s", err)
                	return nodestate
        	}
		defer resp.Body.Close()
        	buffer, err := ioutil.ReadAll(resp.Body)
		v, err := jason.NewObjectFromBytes(buffer)
	        if err != nil {
	//		c.Set(key_unavailable, 1, cache.DefaultExpiration)
         	       	nodestate := false
               		log.Printf("[DEBUG] failed request #3 %s", err)
                	return nodestate
        	}
        	freeslots, err := v.GetInt64("browserSlotsCount", browser, "free")
        	if err != nil {
	//		c.Set(key_unavailable, 1, cache.DefaultExpiration)
                	nodestate := false
			log.Printf("[DEBUG] failed request #4 %s", err)
                	return nodestate
        	}
		//c.Set(key_freeslots, freeslots, cache.DefaultExpiration)
		queuesize, err := v.GetInt64("newSessionRequestCount")
        	if err != nil {
	//		c.Set(key_unavailable, 1, cache.DefaultExpiration)
                	nodestate := false
			log.Printf("[DEBUG] failed request #5 %s", err)
               	 	return nodestate
        	}
		//c.Set(key_queuesize, queuesize, cache.DefaultExpiration)
	//	fkey_freeslots = freeslots
	//	fkey_queuesize = queuesize
	//}
	//var freeslots int64
	//var queuesize int64
	//if x, found := c.Get(key_freeslots);found {
	//	freeslots = x.(int64)
	//}
	//if x, found := c.Get(key_queuesize);found {
         //       queuesize = x.(int64)
        //}
	if (freeslots > 0) && (queuesize < 2) {
        	nodestate := true
		//log.Printf("[DEBUG] node %s:%d is OK (free: %d for %s) queue size: %d.", h.Name, h.Port, freeslots, browser, queuesize)
      		return nodestate
	} else {
                nodestate := false
		//log.Printf("[DEBUG] node %s:%d is out of free slots (free: %d for %s) or queue is too large: %d. Ignore it!", h.Name, h.Port, freeslots, browser, queuesize)
		return nodestate
        }
	return nodestate
}
