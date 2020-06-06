# sev

*Playing around* to better understand golang, better not use it as is.

Small package that will be deployed on top of AWS lambda and send alerts via SNS describing the state of a certificate expiration date.

## usage

```
package main

import (
	"fmt"
	"log"
	"net/http"

	sev "github.com/Norbinsh/cert-expiration-severity"
)

func main() {
	client := &http.Client{}
	target := sev.Website{"https://google.com", *client}
	cert, err := target.GetCert()
	if err != nil {
		log.Fatal(err)
	}
	severity := sev.GetSev(cert)
	fmt.Println(severity.Level) // 0..4
}

```


## tests
generate the certificates:
```
make gen
```
run local, mocked tests:
```
make test
```



