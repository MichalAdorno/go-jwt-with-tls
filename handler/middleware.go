package handler

import (
	"log"
	"net/http"
)

func TlsResponseHeaderHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	}
}

func LogHttpHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%p] %s %s", r, r.Method, r.URL)
		handler(w, r)
	}
}

// func LogHandler(fn http.HandlerFunc) http.HandlerFunc {
//     return func(w http.ResponseWriter, r *http.Request) {
//         x, err := httputil.DumpRequest(r, true)
//         if err != nil {
//             http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
//             return
//         }
//         log.Println(fmt.Sprintf("%q", x))
//         rec := httptest.NewRecorder()
//         fn(rec, r)
//         log.Println(fmt.Sprintf("%q", rec.Body))
//     }
// }
