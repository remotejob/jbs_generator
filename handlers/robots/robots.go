package robots

import (
	"bytes"
	//	"log/syslog"
	"net/http"
	"strings"
)

func Generate(w http.ResponseWriter, r *http.Request) {

	var buffer bytes.Buffer

	sitefull := r.Host
	site := strings.Split(sitefull, ":")[0]

	//	golog.Info("robots: " + site)

	//	splithost := strings.Split(site, ":")
	//
	//	if len(splithost) == 1 {

	buffer.WriteString("User-agent: *\nAllow: /\nSitemap: http://" + site + "/sitemap.xml\n")

	

	w.Header().Add("Content-type", "text/plain")
	w.Write(buffer.Bytes())

}
