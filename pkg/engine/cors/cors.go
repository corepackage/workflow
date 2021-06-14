package cors

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/coredevelopment/workflow/pkg/util"
)

// Validate : to validate the cors policy for the workflow
func Validate(r *http.Request, w http.ResponseWriter, cors map[string]interface{}) error {

	// Setting origin
	w.Header().Set("Access-Control-Allow-Origin", getOrigins(cors["allow-origin"]))

	// validating method
	methods := getMethods(cors["allow-method"])
	if methods != "*" {
		if !strings.Contains(methods, r.Method) {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return errors.New("Method not allowed")
		}
	}
	return nil
}

// getOrigins : to check for allowed origins
func getOrigins(origin interface{}) string {
	allowedOrigins := "*"
	if origin == nil {
		return allowedOrigins
	}
	value, ok := origin.(string)
	if !ok {
		log.Println("getOrigins : Invalid string, setting default value")
		return allowedOrigins
	}
	if len(value) == 1 && value == "*" {
		return allowedOrigins
	}

	origins := strings.Split(value, ",")
	validOrigins := make([]string, 0)
	for _, o := range origins {
		_, err := url.ParseRequestURI(o)
		if err != nil {
			log.Printf("Invalid origin %v, removed with error %v\n", o, err)
			continue
		}
		validOrigins = append(validOrigins, o)

	}
	if len(validOrigins) == 0 {
		return allowedOrigins
	}
	allowedOrigins = strings.Join(validOrigins, ",")
	return allowedOrigins
}

// getMethods : to check for allowed methods
func getMethods(methods interface{}) string {
	defaultMethod := "*"
	if methods == nil {
		return defaultMethod
	}
	methodListString, ok := methods.(string)
	if !ok {
		log.Println("getMethods : Invalid string, setting default value")
		return defaultMethod
	}
	if len(methodListString) == 1 && methodListString == "*" {
		return defaultMethod
	}

	methodList := strings.Split(methodListString, ",")

	// Validating user methods
	defaultMethods := []string{"POST", "GET", "PUT", "DELETE"}
	allowedMethods := make([]string, 0)
	for _, m := range methodList {

		temp := strings.ToUpper(m)
		_, ok := util.FindInArray(temp, defaultMethods)
		if !ok {
			log.Println("Invalid method type, neglecting it")
			continue
		}
		allowedMethods = append(allowedMethods, temp)
	}

	if len(allowedMethods) != 0 {
		return strings.Join(allowedMethods, ",")
	}
	return defaultMethod
}
