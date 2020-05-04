package http

import "net/http"

func DefaultQuery(r *http.Request, key, defaultValue string) string {
	if value, ok := GetQuery(r, key); ok {
		return value
	}
	return defaultValue
}

func GetQuery(r *http.Request, key string) (string, bool) {
	if values, ok := GetQueryArray(r, key); ok {
		return values[0], ok
	}
	return "", false
}

func GetQueryArray(r *http.Request, key string) ([]string, bool) {
	if values, ok := r.URL.Query()[key]; ok && len(values) > 0 {
		return values, true
	}
	return []string{}, false
}
