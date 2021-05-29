package main

//import (
//	"fmt"
//	"golang.org/x/time/rate"
//	"io"
//	"net/http"
//	"sync"
//	"testing"
//	"time"
//)
//
//func Run(task_id, sleeptime, timeout int, ch chan string) {
//	ch_run := make(chan string)
//	go run(task_id, sleeptime, ch_run)
//	select {
//	case re := <-ch_run:
//		ch <- re
//	case <-time.After(time.Duration(timeout) * time.Second):
//		re := fmt.Sprintf("task id %d , timeout", task_id)
//		ch <- re
//	}
//}
//
//func run(task_id, sleeptime int, ch chan string) {
//
//	time.Sleep(time.Duration(sleeptime) * time.Second)
//	ch <- fmt.Sprintf("task id %d , sleep %d second", task_id, sleeptime)
//	return
//}
//
//func TestDemo1(t *testing.T) {
//	input := []int{3, 2, 1}
//	timeout := 2
//	chLimit := make(chan bool, 2)
//	chs := make([]chan string, len(input))
//	limitFunc := func(chLimit chan bool, ch chan string, task_id, sleeptime, timeout int) {
//		Run(task_id, sleeptime, timeout, ch)
//		<-chLimit
//	}
//	startTime := time.Now()
//	fmt.Println("Multirun start")
//	for i, sleeptime := range input {
//		chs[i] = make(chan string, 1)
//		chLimit <- true
//		go limitFunc(chLimit, chs[i], i, sleeptime, timeout)
//	}
//
//	for _, ch := range chs {
//		fmt.Println(<-ch)
//	}
//	endTime := time.Now()
//	fmt.Printf("Multissh finished. Process time %s. Number of task is %d", endTime.Sub(startTime), len(input))
//}
//
//type RequestLimitService struct {
//	Interval time.Duration
//	MaxCount int
//	Lock     sync.Mutex
//	ReqCount int
//}
//
//func NewRequestLimitService(interval time.Duration, maxCnt int) *RequestLimitService {
//	reqLimit := &RequestLimitService{
//		Interval: interval,
//		MaxCount: maxCnt,
//	}
//
//	go func() {
//		ticker := time.NewTicker(interval)
//		for {
//			<-ticker.C
//			reqLimit.Lock.Lock()
//			fmt.Println("Reset Count...")
//			reqLimit.ReqCount = 0
//			reqLimit.Lock.Unlock()
//		}
//	}()
//
//	return reqLimit
//}
//
//func (reqLimit *RequestLimitService) Increase() {
//	reqLimit.Lock.Lock()
//	defer reqLimit.Lock.Unlock()
//
//	reqLimit.ReqCount += 1
//}
//
//func (reqLimit *RequestLimitService) IsAvailable() bool {
//	reqLimit.Lock.Lock()
//	defer reqLimit.Lock.Unlock()
//
//	return reqLimit.ReqCount < reqLimit.MaxCount
//}
//
//var RequestLimit = NewRequestLimitService(10*time.Second, 5)
//
//func helloHandler(w http.ResponseWriter, r *http.Request) {
//	if RequestLimit.IsAvailable() {
//		RequestLimit.Increase()
//		fmt.Println(RequestLimit.ReqCount)
//		io.WriteString(w, "Hello world!\n")
//	} else {
//		fmt.Println("Reach request limiting!")
//		io.WriteString(w, "Reach request limit!\n")
//	}
//}
//
//func TestDemo2(t *testing.T) {
//	fmt.Println("Server Started!")
//	http.HandleFunc("/", helloHandler)
//	http.ListenAndServe(":8000", nil)
//}
//
//func TestDemo3(t *testing.T) {
//	mux := http.NewServeMux()
//	mux.HandleFunc("/", okHandler)
//	// Wrap the servemux with the limit middleware.
//	http.ListenAndServe(":4000", limit(mux))
//}
//
//
//var limiter = rate.NewLimiter(2, 5)
//func limit(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if limiter.Allow() == false {
//			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
//			return
//		}
//		next.ServeHTTP(w, r)
//	})
//}
//
//
//func okHandler(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("OK"))
//}
//
//// Create a custom visitor struct which holds the rate limiter for each
//// visitor and the last time that the visitor was seen.
//type visitor struct {
//	limiter  *rate.Limiter
//	lastSeen time.Time
//}
//// Change the the map to hold values of the type visitor.
//var visitors = make(map[string]*visitor)
//var mtx sync.Mutex
//// Run a background goroutine to remove old entries from the visitors map.
//func init() {
//	go cleanupVisitors()
//}
//func addVisitor(ip string) *rate.Limiter {
//	limiter := rate.NewLimiter(2, 5)
//	mtx.Lock()
//	// Include the current time when creating a new visitor.
//	visitors[ip] = &visitor{limiter, time.Now()}
//	mtx.Unlock()
//	return limiter
//}
//func getVisitor(ip string) *rate.Limiter {
//	mtx.Lock()
//	v, exists := visitors[ip]
//	if !exists {
//		mtx.Unlock()
//		return addVisitor(ip)
//	}
//	// Update the last seen time for the visitor.
//	v.lastSeen = time.Now()
//	mtx.Unlock()
//	return v.limiter
//}
//// Every minute check the map for visitors that haven't been seen for
//// more than 3 minutes and delete the entries.
//func cleanupVisitors() {
//	for {
//		time.Sleep(time.Minute)
//		mtx.Lock()
//		for ip, v := range visitors {
//			if time.Now().Sub(v.lastSeen) > 3*time.Minute {
//				delete(visitors, ip)
//			}
//		}
//		mtx.Unlock()
//	}
//}
//func limit(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		limiter := getVisitor(r.RemoteAddr)
//		if limiter.Allow() == false {
//			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
//			return
//		}
//		next.ServeHTTP(w, r)
//	})
//}
