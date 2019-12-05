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
	logIn("alex.usavchenko", "P@ssw0rd")
}

func logIn(login string, password string) {
	var sessionIdCookie http.Cookie
	var lxTokenCookie http.Cookie
	var requestVerificationCookie http.Cookie
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			req.Response.Header.Get("Set-Cookie")
			for k, v := range req.Response.Header {
				if k == "Set-Cookie" && len(v) == 4 {
					lxtoken := v[3]
					lxTokenCookie = http.Cookie{Name: "LXToken",
						Value:  lxtoken[8:strings.IndexByte(lxtoken, ';')],
						Secure: true, HttpOnly: true, Path: "/", Domain: "portalpacjenta.luxmed.pl"}
					sessionToken := v[0]
					sessionIdCookie = http.Cookie{Name: "ASP.NET_SessionId",
						Value:  sessionToken[18 : strings.IndexByte(lxtoken, ';')-2],
						Secure: true, HttpOnly: true, Path: "/", Domain: "portalpacjenta.luxmed.pl"}
					log.Println(sessionIdCookie)
					log.Println(v[0])
					log.Println(lxTokenCookie)
				}
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
			requestVerificationCookie = http.Cookie{Name: "__RequestVerificationToken_L1BhdGllbnRQb3J0YWw1",
				Value:  v[2][strings.IndexByte(v[2], '=')+1 : strings.IndexByte(v[2], ';')],
				Secure: true, HttpOnly: true, Path: "/", Domain: "portalpacjenta.luxmed.pl"}
			log.Println(requestVerificationCookie)
			//log.Println(v[2])
		}
	}
	defer authResponse.Body.Close()

	searchForm := url.Values{}
	searchForm.Add("DateOption", "SelectedDate")
	searchForm.Add("FilterType", "Coordination")
	searchForm.Add("CoordinationActivityId", "90")
	searchForm.Add("IsFFS", "False")
	searchForm.Add("IsDisabled", "False")
	searchForm.Add("MaxPeriodLength", "0")
	searchForm.Add("PayersCount", "0")
	searchForm.Add("FromDate", "06.12.2019")
	searchForm.Add("ToDate", "18.12.2019")
	searchForm.Add("DefaultSearchPeriod", "14")
	searchForm.Add("SelectedSearchPeriod", "14")
	searchForm.Add("CustomRangeSelected", "False")
	searchForm.Add("CityId", "70")
	searchForm.Add("DateRangePickerButtonDefaultLabel", "Inny zakres")
	searchForm.Add("ServiceId", "4430")
	searchForm.Add("TimeOption", "0")
	searchForm.Add("PayerId", "167225")
	searchRequest, _ := http.NewRequest("POST",
		"https://portalpacjenta.luxmed.pl/PatientPortal/Reservations/Reservation/PartialSearch",
		strings.NewReader(searchForm.Encode()))
	searchRequest.AddCookie(&requestVerificationCookie)
	searchRequest.AddCookie(&lxTokenCookie)
	searchRequest.AddCookie(&sessionIdCookie)
	searchRequest.Header.Set("Content-type", "application/x-www-form-urlencoded")
	searchRequest.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:37.0) Gecko/20100101 Firefox/37.0")
	searchRequest.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	searchReponse, _ := client.Do(searchRequest)
	//log.Println(err)
	bytes, _ := ioutil.ReadAll(searchReponse.Body)
	log.Println(string(bytes))
	//document, _ := goquery.NewDocumentFromReader(authResponse.Body)
	//document.Find("input").Each(func(i int, selection *goquery.Selection) {
	//	log.Println(selection.Attr("value"))
	//})

}
