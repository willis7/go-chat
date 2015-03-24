package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte(`
			<html>
				<head>
					<title>Chat</title>
				</head>
				<body>
					Lets chat!
				</body>
			</html>
		`))
	})
	// start the webserver
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
