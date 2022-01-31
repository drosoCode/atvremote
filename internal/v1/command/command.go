package command

import (
	"crypto/tls"
	"errors"
	"strconv"
	"time"

	"github.com/drosocode/atvremote/internal/remote"
)

type Command struct {
	Buffer       []byte
	Connection   *tls.Conn
	Certificates *tls.Certificate
	Address      string
	Port         int
}

func New(addr string, port int, certs *tls.Certificate) Command {
	return Command{Buffer: make([]byte, 512), Address: addr, Port: port, Certificates: certs}
}

func (c *Command) Connect() error {
	config := &tls.Config{Certificates: []tls.Certificate{*c.Certificates}, InsecureSkipVerify: true}

	conn, err := tls.Dial("tcp", c.Address+":"+strconv.Itoa(c.Port), config)
	if err != nil {
		return errors.New("command - connexion: " + err.Error())
	}
	c.Connection = conn

	state := conn.ConnectionState()
	for !state.HandshakeComplete {
	}

	return nil
}

func (c *Command) read() ([]byte, error) {
	len, err := c.Connection.Read(c.Buffer)
	if err != nil {
		return nil, err
	}
	if len < 5 {
		return c.read()
	}

	return c.Buffer[0:len], nil
}

func (c *Command) sendConfiguration() error {
	data := []byte{1, 0, 0, 21, 0, 0, 0, 1, 0, 0, 0, 1, 32, 3, 0, 0, 0, 0, 0, 0, 4, 116, 101, 115, 116}
	c.Connection.Write(data)

	data, err := c.read()
	if err != nil {
		return err
	}
	if data[0] != 1 || data[1] != 7 || data[2] != 0 {
		return errors.New("command - send: invalid ack")
	}
	_, err = c.read()
	if err != nil {
		return err
	}
	_, err = c.read()
	if err != nil {
		return err
	}
	return nil
}

func (c *Command) SendKey(keycode remote.RemoteKeyCode) error {
	err := c.sendConfiguration()
	if err != nil {
		return err
	}

	code := byte(keycode)
	press := []byte{1, 2, 0, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, code}
	_, err = c.Connection.Write(press)
	if err != nil {
		return err
	}

	// wait 100ms between press and released
	time.Sleep(100 * time.Millisecond)

	release := []byte{1, 2, 0, 16, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, code}
	_, err = c.Connection.Write(release)
	if err != nil {
		return err
	}

	return nil
}

func text2byte(text string) []byte {
	arr := []byte{}
	for _, c := range text {
		arr = append(arr, byte(int(c)))
	}
	return arr
}

func (c *Command) SendIntent(intent string) error {
	err := c.sendConfiguration()
	if err != nil {
		return err
	}

	//intent:#Intent;component=
	iStart := []byte{105, 110, 116, 101, 110, 116, 58, 35, 73, 110, 116, 101, 110, 116, 59, 99, 111, 109, 112, 111, 110, 101, 110, 116, 61}
	// ex: com.netflix.ninja/.MainActivity
	intentMsg := append(iStart, text2byte(intent)...)
	// append  ;end
	intentMsg = append(intentMsg, []byte{59, 101, 110, 100}...)

	message := append([]byte{1, 16, 0, byte(len(intentMsg))}, intentMsg...)

	_, err = c.Connection.Write(message)
	if err != nil {
		return err
	}

	return nil
}
