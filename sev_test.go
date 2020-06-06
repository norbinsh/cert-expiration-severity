package sev

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBounds(t *testing.T) {

	cases := []struct {
		Name        string
		CertDir     string
		ExpectedSev Severity
	}{
		{
			"expect sev 0",
			"sev0",
			Severity{0},
		},
		{
			"expect sev 1",
			"sev1",
			Severity{1},
		},
		{
			"expect sev 2",
			"sev2",
			Severity{2},
		},
		{
			"expect sev 3",
			"sev3",
			Severity{3},
		},
		{
			"expect sev 4",
			"sev4",
			Severity{4},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {

			server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))

			localCert, err := tls.LoadX509KeyPair("certs/"+test.CertDir+"/cert.pem", "certs/"+test.CertDir+"/key.pem")

			assertError(t, err)

			server.TLS = &tls.Config{
				Certificates: []tls.Certificate{localCert},
			}

			server.StartTLS()
			defer server.Close()

			client := server.Client()
			want := test.ExpectedSev
			target := Website{server.URL, *client}
			cert, err := target.GetCert()
			assertError(t, err)

			got := GetSev(cert)

			assertDiff(t, got, want)

		})
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func assertDiff(t *testing.T, got, want Severity) {
	t.Helper()
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}
