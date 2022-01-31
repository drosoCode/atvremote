package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"github.com/drosocode/atvremote/internal/cert"
	pairing "github.com/drosocode/atvremote/internal/v1/pairing"
)

func main() {
	cert.CreateCertificate("androidtv-remote", []string{"192.168.1.20"})

	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatal(err)
	}

	p := pairing.New("192.168.1.20", 6467, &cert)

	err = p.Connect()
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter code: ")
	code, _ := reader.ReadString('\n')

	secret, err := p.Secret(code)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(secret)

	//secret := "="
}
