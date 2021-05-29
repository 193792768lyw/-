package main

//func HelloHandler(w http.ResponseWriter, req *http.Request) {
//	w.Write([]byte("Hello, World!"))
//}
//
//func main() {
//	// Create a request limiter per handler.
//	lmt := tollbooth.NewLimiter(2, nil)
//
//	http.Handle("/", HeaderLimiterHandler(lmt, HelloHandler))
//
//	http.ListenAndServe(":12345", nil)
//
//}
//
//// HeaderLimiterHandler 通用的基于header key的限流
//func HeaderLimiterHandler(lmt *limiter.Limiter, f func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
//	return func(req http.ResponseWriter, rp *http.Request) {
//
//		httpError := tollbooth.LimitByRequest(lmt, req, rp)
//		if httpError != nil {
//			fmt.Println(httpError)
//			return
//		}
//		f(req, rp)
//	}
//}
