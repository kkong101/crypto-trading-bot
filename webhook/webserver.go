package webhook

// StartWebServer gin web server start
func StartWebServer() {
	r := NewRouter()
	err := r.Run(":7579")
	if err != nil {
		panic(err)
	}
}
