package main

import (
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	logIn("alex.usavchenko", "P@ssw0rd")
}

func logIn(login string, password string) {

	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			req.Response.Header.Get("Set-Cookie")
			for k, v := range req.Response.Header {
				if k == "Set-Cookie" && len(v) == 4 {
					lxtoken := v[3]
					log.Println(http.Cookie{Name: "LXToken",
						Value:  lxtoken[8:strings.IndexByte(lxtoken, ';')],
						Secure: true, HttpOnly: true, Path: "/"})
					sessionToken := v[0]
					sessionIdCookie := http.Cookie{Name: "ASP.NET_SessionId",
						Value:  sessionToken[18:strings.IndexByte(lxtoken, ';')],
						Secure: true, HttpOnly: true, Path: "/"}
					log.Println(sessionIdCookie)
				}
				//headers[strings.ToLower(k)] = string(v[0])
				//log.Println(k)
				//log.Println(v)

			}
			return nil
		},
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
	request.AddCookie(&http.Cookie{Name: "LXCookieMonit", Value: "1"})
	authResponse, _ := client.Do(request)
	for k, v := range authResponse.Header {
		if k == "Set-Cookie" {
			log.Println(v[2])
		}
	}
	defer authResponse.Body.Close()
	//document, _ := goquery.NewDocumentFromReader(authResponse.Body)
	//document.Find("input").Each(func(i int, selection *goquery.Selection) {
	//	log.Println(selection.Attr("value"))
	//})

}
