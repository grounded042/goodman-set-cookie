package main

import (
	"fmt"
	"net/http"
)

func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "Test", Value: "Yo"}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}

func cookieCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "Test", Value: "Yo"}
	w.Header().Set("Cookie-Cookie", fmt.Sprintf("%s=%s", cookie.Name, cookie.Value))
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/set-cookie", setCookieHandler)
	http.HandleFunc("/cookie-cookie", cookieCookieHandler)
	http.ListenAndServe(":8080", nil)
}
