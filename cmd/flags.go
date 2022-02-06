package cmd

import (
	"crypto/tls"
	"errors"
	"flag"
	"log"
	"os"

	"github.com/drosocode/atvremote/pkg/common"
)

type config struct {
	Version      int
	Certificates tls.Certificate
	IP           string
	Pair         bool
	Command      []common.RemoteKeyCode
	Open         string
	Link         string
	Info         bool
}

// parse flags from the cli, do some checks and return a config struct
func parseFlags() config {
	cert := flag.String("cert", "cert.pem", "path to certificate file")
	key := flag.String("key", "key.pem", "path to key file")
	ip := flag.String("ip", "", "android tv ip")
	version := flag.Int("version", 1, "remote protocol version [1 or 2]")
	pair := flag.Bool("pair", false, "pairing mode")
	command := flag.String("command", "", "command to execute")
	open := flag.String("open", "", "app activity to run (only for v1)")
	link := flag.String("link", "", "link to open (only for v2)")
	info := flag.Bool("info", false, "print device info (only for v2)")

	flag.Parse()

	if *ip == "" {
		flag.Usage()
		log.Fatalf("please provide the ip of your android tv device")
	}

	err := checkPath(*cert, *key)
	if *pair {
		if err != nil {
			err := common.CreateCertificate("androidtv-remote", []string{*ip}, *cert, *key)
			if err != nil {
				log.Fatalf("unable to create certificate: %s", err)
			}
		}
	} else {
		if err != nil {
			flag.Usage()
			log.Fatal(err)
		}
		// check flags
		if *version == 1 {
			if *link != "" {
				flag.Usage()
				log.Fatalf("flag -link is invalid for protocol v1")
			}
			if *info {
				flag.Usage()
				log.Fatalf("flag -info is invalid for protocol v1")
			}
		} else {
			if *open != "" {
				flag.Usage()
				log.Fatalf("flag -open is invalid for protocol v2")
			}
		}
	}

	certs, err := tls.LoadX509KeyPair(*cert, *key)
	if err != nil {
		log.Fatalf("unable to load certificates: %s", err)
	}

	return config{Version: *version, Certificates: certs, IP: *ip, Pair: *pair, Command: common.ParseKeys(*command), Open: *open, Link: *link, Info: *info}
}

// check if the paths to the certs exists
func checkPath(cert string, key string) error {
	if _, err := os.Stat(cert); err != nil && os.IsNotExist(err) {
		return errors.New("path to certificate file does not exists")
	}
	if _, err := os.Stat(key); err != nil && os.IsNotExist(err) {
		return errors.New("path to key file does not exists")
	}
	return nil
}
