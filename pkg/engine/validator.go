package engine

import (
	"errors"
	"log"
	"net/http"
)

var (
	errValidationFailed = errors.New("validation failed")
)

// Validate : Validating access request
func Validate(r *http.Request, w http.ResponseWriter, wf *Workflow) error {

	// headers := (map[string][]string)(r.Header)
	err := CORSValidate(r, w, wf.CORS)
	if err != nil {
		log.Println("Error in CORS Policy")
		w.Write([]byte("Error in CORS Policy"))
		return errValidationFailed
	}

	//TODO: Validate request if authorizer present
	err = AuthValidate(r, wf.Authorizer)
	if err != nil {
		log.Println("Request Not Valid")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Request Not Valid"))
		return errValidationFailed
	}
	return nil
}
