package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	v1c "github.com/drosocode/atvremote/pkg/v1/command"
	v1p "github.com/drosocode/atvremote/pkg/v1/pairing"

	v2c "github.com/drosocode/atvremote/pkg/v2/command"
	v2p "github.com/drosocode/atvremote/pkg/v2/pairing"
)

func Execute() {
	cfg := parseFlags()

	if cfg.Pair {
		if cfg.Version == 1 {
			p := v1p.New(cfg.IP, 6467, &cfg.Certificates)
			err := p.Connect()
			if err != nil {
				log.Fatal(err)
			}
			err = p.Secret(enterCode())
			if err != nil {
				log.Fatal(err)
			}

		} else {
			p := v2p.New(cfg.IP, 6467, &cfg.Certificates)
			err := p.Connect()
			if err != nil {
				log.Fatal(err)
			}
			err = p.Secret(enterCode())
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		if cfg.Version == 1 {
			cmd := v1c.New(cfg.IP, 6466, &cfg.Certificates)
			err := cmd.Connect()
			if err != nil {
				log.Fatal(err)
			}

			for _, k := range cfg.Command {
				cmd.SendKey(k)
			}

			if cfg.Open != "" {
				cmd.SendIntent(cfg.Open)
			}
		} else {
			cmd := v2c.New(cfg.IP, 6466, &cfg.Certificates)
			err := cmd.Connect()
			if err != nil {
				log.Fatal(err)
			}

			for _, k := range cfg.Command {
				cmd.SendKey(k)
			}

			if cfg.Link != "" {
				cmd.OpenLink(cfg.Link)
			}

			if cfg.Info {
				data, _ := cmd.GetData()
				bdata, _ := json.Marshal(data)
				fmt.Println(string(bdata))
			}
		}
	}
}

func enterCode() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter code: ")
	code, _ := reader.ReadString('\n')
	return code
}
