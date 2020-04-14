package view

func View(viewName string) string {
	return "./www/" + viewName + ".html"
}