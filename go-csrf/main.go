package main

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"

	"github.com/justinas/nosurf"
)

var loginTemplate string = `
<!doctype html>
<html>
  <body>
  	<p>Login</p>
    <form action="/" method="POST">
      <input type="hidden" name="csrf_token" value="{{ .CSRFToken }}">
      <input type="text" name="username">
      <input type="submit" value="Send">
    </form>
  </body>
</html>	
`
var login = template.Must(template.New("page").Parse(loginTemplate))

var loggedInTemplate string = `
<!doctype html>
<html>
  <body>
  	<p>If you're seeing this the CSRF token was validated.</p>
  	<p>Hello {{ .Username }}</p>
  </body>
</html>	
`
var loggedIn = template.Must(template.New("loggedIn").Parse(loggedInTemplate))

func protect(h http.Handler) http.Handler {
	csrf := nosurf.New(h)
	// this is optional, if not cookie is set there's a default `csrf_token` one
	// that'll be sent
	csrf.SetBaseCookie(http.Cookie{
		Name:     "_csrf",
		HttpOnly: true,
		Path:     "/",
	})
	return csrf
}

func get(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"CSRFToken": nosurf.Token(r),
	}
	login.Execute(w, data)
}

func post(w http.ResponseWriter, r *http.Request) {
	// this happens automatically, just doing it here to show how it's done
	verified := nosurf.VerifyToken(
		nosurf.Token(r),
		r.PostFormValue("csrf_token"),
	)
	slog.Info("csrf token was sent", "verified", verified)

	data := map[string]string{
		"Username": r.PostFormValue("username"),
	}
	loggedIn.Execute(w, data)
}

func main() {
	http.Handle("GET /", protect(http.HandlerFunc(get)))
	http.Handle("POST /", protect(http.HandlerFunc(post)))

	fmt.Println("make request to :8080/")
	http.ListenAndServe(":8080", nil)
}
