package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"github.com/drosocode/atvremote/internal/cert"
	"github.com/drosocode/atvremote/internal/remote"
	"github.com/drosocode/atvremote/internal/v1/command"
	"github.com/drosocode/atvremote/internal/v1/pairing"

	c2 "github.com/drosocode/atvremote/internal/v2/command"
)

func main() {
	//v1_test()
	v2_test()
}

func v1_test() {
	ip := "192.168.1.20"
	// create certificates
	cert.CreateCertificate("androidtv-remote", []string{ip}, "cert.pem", "key.pem")

	// load certificates
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatal(err)
	}

	// start a new pairing process
	p := pairing.New(ip, 6467, &cert)

	err = p.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// prompts the user to enter the pairing code
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter code: ")
	code, _ := reader.ReadString('\n')

	// finish the pairing process
	err = p.Secret(code)
	if err != nil {
		log.Fatal(err)
	}

	// connect in command mode
	cmd := command.New(ip, 6466, &cert)
	cmd.Connect()
	// send volume down keypress
	cmd.SendKey(remote.RemoteKeyCode_KEYCODE_VOLUME_DOWN)
	// open netflix
	cmd.SendIntent("com.netflix.ninja/.MainActivity")
}

func v2_test() {
	ip := "192.168.1.20"
	// create certificates
	//cert.CreateCertificate("androidtv-remote", []string{ip}, "cert.pem", "key.pem")

	// load certificates
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatal(err)
	}

	/*
		// start a new pairing process
		p := pv2.New(ip, 6467, &cert)

		err = p.Connect()
		if err != nil {
			log.Fatal(err)
		}

		// prompts the user to enter the pairing code
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter code: ")
		code, _ := reader.ReadString('\n')

		// finish the pairing process
		err = p.Secret(code)
		if err != nil {
			log.Fatal(err)
		}
	*/

	// connect in command mode
	cmd := c2.New(ip, 6466, &cert)
	cmd.Connect()
	// send volume down keypress
	//cmd.SendKey(remote.RemoteKeyCode_KEYCODE_VOLUME_DOWN)
	// open netflix
	//cmd.SendIntent("com.netflix.ninja/.MainActivity")
}
