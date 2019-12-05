package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	logIn("alex.usavchenko", "")
}

func logIn(login string, password string) {
	client := http.Client{
		Timeout: time.Duration(5*time.Second),
	}
	form := url.Values{}
	form.Add("Login", "alex.usavchenko")
	form.Add("Password", "P@ssw0rd")
	request, _ := http.NewRequest("POST", "https://portalpacjenta.luxmed.pl/PatientPortal/Account/LogIn",
		strings.NewReader(form.Encode()))
	request.Header.Set("Content-type", "application/x-www-form-urlencoded")
	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:37.0) Gecko/20100101 Firefox/37.0")
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Set("Referer", "https://portalpacjenta.luxmed.pl/PatientPortal/Account/LogOn")
	request.AddCookie(&http.Cookie{Name:"LXCookieMonit", Value:"1"})
	authResponse, _ := client.Do(request)
	bytes, _ := ioutil.ReadAll(authResponse.Body)
	log.Println(string(bytes))



}
