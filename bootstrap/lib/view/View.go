package view

func View(viewName string) string {
	return "./www/" + viewName + ".html" // The path returned which will be then server with http.ServeFile
}