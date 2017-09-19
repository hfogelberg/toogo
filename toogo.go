package toogo

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// Read Enviroment variable. If it's missing use default value.
// Ex port := Getenv("PORT", ":3000")
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

// Serve favicon. The file must be in the templates directory
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/favicon.ico")
}

// Formating
// Extract abrevated day from timestamp
func timeToDay(dt int) string {
	layout := "Mon"
	t := int64(dt)
	f := time.Unix(t, 0).Format(layout)
	return f
}

// Extract time in 24 hour format from timestamp
func timeToHour(dt int) string {
	layout := "Mon 15:00"
	t := int64(dt)
	f := time.Unix(t, 0).Format(layout)
	return f
}

// Round float to two decimals and return as string
func roundToTwo(f float64) string {
	return fmt.Sprintf("%0.2f", f)
}
