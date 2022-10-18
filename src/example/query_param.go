package example

import (
	"io"
	"net/http"
)

func QueryParamDisplayHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "query: "+req.FormValue("name"))
	io.WriteString(res, "\nphone: "+req.FormValue("phone"))
	println("Enter this in your browser:  http://localhost:9000/example?name=mehmet&phone=533-533")
}
