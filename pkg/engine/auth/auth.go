package auth

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/coredevelopment/workflow/internal/models"
	"github.com/coredevelopment/workflow/pkg/util"
)

// Validate : to validate incoming request
func Validate(r *http.Request, auth *models.Authorizer) error {
	if auth == nil {
		log.Println("No authorizer found, continuing")
		return nil
	}

	// Getting token key
	var requestToken string
	switch strings.ToLower(auth.Input) {
	case "header":
		requestToken = r.Header.Get(auth.AKey)
	case "body":
		// TODO: Validate key on nested level
		userData, err := util.ParseData(r.Body)
		if err != nil {
			return err
		}
		var ok bool
		requestToken, ok = userData[auth.AKey].(string)
		if !ok {
			return errors.New("Error parsing auth key")
		}

	}

	// Invoking user's authorizer
	success, err := invokeAuthorizer(requestToken)
	if !success && err != nil {
		return errors.New("Authorization failed")
	} else if err != nil {
		log.Println("Validate : error invoking user authorizer")
		return err
	}

	return nil
}

// invokeAuthorizer : to invoke user authorizer
func invokeAuthorizer(token string) (bool, error) {
	return true, nil
}
