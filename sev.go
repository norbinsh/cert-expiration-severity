package sev

import (
	"crypto/x509"
	"fmt"
	"math"
	"net/http"
	"time"
)

// Severity ranges from 0-4, 0 being high amount of days until certificate expiration
type Severity struct {
	Level int
}

// Website hostname of which we are are checking the certificate for
type Website struct {
	Fullurl string
	Client  http.Client
}

// GetSev parse the x509 certificate and return a Severity
func GetSev(c *x509.Certificate) Severity {
	expirationTime := c.NotAfter
	daysLeft := int(math.Round(time.Until(expirationTime).Hours() / 24))
	fmt.Println(daysLeft)
	switch {
	case daysLeft <= 20:
		return Severity{4}
	case daysLeft <= 40:
		return Severity{3}
	case daysLeft <= 60:
		return Severity{2}
	case daysLeft <= 80:
		return Severity{1}
	default:
		return Severity{0}
	}
}

// GetCert returns the x509 certificate from the server
func (w *Website) GetCert() (*x509.Certificate, error) {
	res, err := w.Client.Get(w.Fullurl)
	if err != nil {
		return nil, err
	}
	return res.TLS.PeerCertificates[0], nil
}
