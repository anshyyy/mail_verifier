package mailService

import (
	"fmt"
	"log"
	"os"

	emailverifier "github.com/AfterShip/email-verifier"
)

var (
	proxyURI = os.Getenv("PROXY_URI")
)

func VerifySingleMail(email string, useProxy bool) (*emailverifier.Result, error) {
	verifier := emailverifier.NewVerifier()
	log.Print("proxy use ",useProxy)
	if useProxy {
		log.Print("Using proxy true",proxyURI)
		verifier = verifier.Proxy(proxyURI)
	}

	result, err := verifier.Verify(email)
	if err != nil {
		return nil, fmt.Errorf("failed to verify email: %w", err)
	}

	return result, nil
}

func VerifyBulkMail(emails []string, useProxy bool) (map[string]*emailverifier.Result, error) {
	verifier := emailverifier.NewVerifier()
	if useProxy && proxyURI != "" {
		log.Print("Using proxy true",proxyURI)
		verifier = verifier.Proxy(proxyURI)
	}


	results := make(map[string]*emailverifier.Result)
	for _, email := range emails {
		result, err := verifier.Verify(email)
		if err != nil {
			results[email] = nil
			continue
		}
		results[email] = result
	}

	return results, nil
}
