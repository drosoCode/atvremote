package pairing

import (
	"crypto/rsa"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"strconv"
)

func getSize(s int) []byte {
	return []byte{
		byte(s >> 24 & 0xFF),
		byte(s >> 16 & 0xFF),
		byte(s >> 8 & 0xFF),
		byte(s & 0xFF),
	}
}

func (p *Pairing) send(data *Message) error {
	raw, err := json.Marshal(data)
	if err != nil {
		return err
	}
	//conn.Write([]byte(strconv.Itoa(len(raw))))
	//_, err = conn.Write(raw)

	size := getSize(len(raw))

	_, err = p.Connection.Write(size)
	_, err = p.Connection.Write(raw)
	if err != nil {
		return err
	}
	return nil
}

func (p *Pairing) read() (*Message, error) {
	len, err := p.Connection.Read(p.Buffer)
	if err != nil {
		return nil, err
	}
	if len < 5 {
		return p.read()
	}

	resp := Message{}
	err = json.Unmarshal(p.Buffer[0:len], &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type Pairing struct {
	Buffer       []byte
	Connection   *tls.Conn
	Certificates *tls.Certificate
	Address      string
	Port         int
}

func New(addr string, port int, certs *tls.Certificate) Pairing {
	return Pairing{Buffer: make([]byte, 512), Address: addr, Port: port, Certificates: certs}
}

func (p *Pairing) Connect() error {
	config := &tls.Config{Certificates: []tls.Certificate{*p.Certificates}, InsecureSkipVerify: true}

	conn, err := tls.Dial("tcp", p.Address+":"+strconv.Itoa(p.Port), config)
	if err != nil {
		return errors.New("pair - connexion: " + err.Error())
	}
	p.Connection = conn

	state := conn.ConnectionState()
	for !state.HandshakeComplete {
	}

	// ------------- Pairing Request --------------------

	data := Message{ProtocolVersion: 1, Payload: PariringRequest{ServiceName: "androidtvremote", ClientName: "interface Web"}, Type: 10, Status: 200}

	if err = p.send(&data); err != nil {
		return errors.New("pair - pairing request - send: " + err.Error())
	}

	resp, err := p.read()
	if err != nil {
		return errors.New("pair - pairing request - receive : " + err.Error())
	}
	if resp == nil || resp.Type != 11 {
		return errors.New("pair - pairing request - receive: nil or invalid response")
	}

	// ------------- Options Request --------------------
	opt := Message{ProtocolVersion: 1, Payload: OptionsRequest{OutputEncodings: []OptionEncoding{OptionEncoding{SymbolLength: 4, Type: 3}}, InputEncodings: []OptionEncoding{OptionEncoding{SymbolLength: 4, Type: 3}}, PreferredRole: 1}, Type: 20, Status: 200}
	if err = p.send(&opt); err != nil {
		return errors.New("pair - options request - send: " + err.Error())
	}

	resp, err = p.read()
	if err != nil {
		return errors.New("pair - options request - receive: " + err.Error())
	}
	if resp == nil || resp.Type != 20 {
		return errors.New("pair - options request - receive: nil or invalid response")
	}

	// ------------- Configuration Request --------------------
	cfg := Message{ProtocolVersion: 1, Payload: ConfigurationRequest{Encoding: OptionEncoding{SymbolLength: 4, Type: 3}, ClientRole: 1}, Type: 30, Status: 200}
	if err = p.send(&cfg); err != nil {
		return errors.New("pair - configuration request - send: " + err.Error())
	}

	resp, err = p.read()
	if err != nil {
		return errors.New("pair - configuration request - receive: " + err.Error())
	}
	if resp == nil || resp.Type != 31 {
		return errors.New("pair - configuration request - receive: nil or invalid response")
	}

	return nil
}

func (p *Pairing) Secret(code string) error {
	defer p.Connection.Close()
	// ------------- Secret Request --------------------
	crt, _ := x509.ParseCertificate(p.Certificates.Certificate[0])
	serverPublicKey := p.Connection.ConnectionState().PeerCertificates[0].PublicKey.(*rsa.PublicKey)
	clientPublicKey := crt.PublicKey.(*rsa.PublicKey)

	hash := sha256.New()
	hash.Write(clientPublicKey.N.Bytes())
	//hash.Write([]byte(strconv.Itoa(clientPublicKey.E)))
	hash.Write([]byte{1, 0, 1})

	hash.Write(serverPublicKey.N.Bytes())
	//hash.Write([]byte(strconv.Itoa(serverPublicKey.E)))
	hash.Write([]byte{1, 0, 1})

	codeEnd, _ := hex.DecodeString(code[2:4])
	hash.Write(codeEnd)
	secret := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	sec := Message{ProtocolVersion: 1, Payload: SecretRequest{Secret: secret}, Type: 40, Status: 200}
	if err := p.send(&sec); err != nil {
		return errors.New("pair - secret request - send: " + err.Error())
	}

	resp, err := p.read()
	if err != nil {
		return errors.New("pair - secret request - receive: " + err.Error())
	}
	if resp == nil || resp.Type != 41 {
		return errors.New("pair - secret request - receive: nil or invalid response")
	}

	// secret := resp.Payload.(map[string]interface{})["secret"].(string)

	return nil
}
