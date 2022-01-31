package v1

import (
	"bufio"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
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

func send[T any](conn *tls.Conn, data *T) error {
	raw, err := json.Marshal(data)
	if err != nil {
		return err
	}
	//conn.Write([]byte(strconv.Itoa(len(raw))))
	//_, err = conn.Write(raw)

	size := getSize(len(raw))

	_, err = conn.Write(size)
	_, err = conn.Write(raw)
	if err != nil {
		return err
	}
	return nil
}

func read(conn *tls.Conn, buf *[]byte) (*Message, error) {
	len, err := conn.Read(*buf)
	if err != nil {
		return nil, err
	}
	if len < 5 {
		return read(conn, buf)
	}

	resp := Message{}
	err = json.Unmarshal((*buf)[0:len], &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func Pair(addr string, port int, clientCerts *tls.Certificate) (string, error) {
	config := &tls.Config{Certificates: []tls.Certificate{*clientCerts}, InsecureSkipVerify: true}

	conn, err := tls.Dial("tcp", addr+":"+strconv.Itoa(port), config)
	if err != nil {
		return "", errors.New("pair - connexion: " + err.Error())
	}
	defer conn.Close()
	log.Println("client: connected to: ", conn.RemoteAddr())

	state := conn.ConnectionState()
	for !state.HandshakeComplete {
	}

	buf := make([]byte, 512)
	// ------------- Pairing Request --------------------

	data := Message{ProtocolVersion: 1, Payload: PariringRequest{ServiceName: "androidtvremote", ClientName: "interface Web"}, Type: 10, Status: 200}

	if err = send(conn, &data); err != nil {
		return "", errors.New("pair - pairing request - send: " + err.Error())
	}

	resp, err := read(conn, &buf)
	if err != nil {
		return "", errors.New("pair - pairing request - receive : " + err.Error())
	}
	if resp == nil || resp.Type != 11 {
		fmt.Println(resp.Type)
		return "", errors.New("pair - pairing request - receive: nil or invalid response")
	}

	// ------------- Options Request --------------------
	opt := Message{ProtocolVersion: 1, Payload: OptionsRequest{OutputEncodings: []OptionEncoding{OptionEncoding{SymbolLength: 4, Type: 3}}, InputEncodings: []OptionEncoding{OptionEncoding{SymbolLength: 4, Type: 3}}, PreferredRole: 1}, Type: 20, Status: 200}
	if err = send(conn, &opt); err != nil {
		return "", errors.New("pair - options request - send: " + err.Error())
	}

	resp, err = read(conn, &buf)
	if err != nil {
		return "", errors.New("pair - options request - receive: " + err.Error())
	}
	if resp == nil || resp.Type != 20 {
		return "", errors.New("pair - options request - receive: nil or invalid response")
	}

	// ------------- Configuration Request --------------------
	cfg := Message{ProtocolVersion: 1, Payload: ConfigurationRequest{Encoding: OptionEncoding{SymbolLength: 4, Type: 3}, ClientRole: 1}, Type: 30, Status: 200}
	if err = send(conn, &cfg); err != nil {
		return "", errors.New("pair - configuration request - send: " + err.Error())
	}

	resp, err = read(conn, &buf)
	if err != nil {
		return "", errors.New("pair - configuration request - receive: " + err.Error())
	}
	if resp == nil || resp.Type != 31 {
		fmt.Println(resp.Type)
		return "", errors.New("pair - configuration request - receive: nil or invalid response")
	}

	// ------------- Secret Request --------------------
	crt, _ := x509.ParseCertificate(clientCerts.Certificate[0])
	serverPublicKey := state.PeerCertificates[0].PublicKey.(*rsa.PublicKey)
	clientPublicKey := crt.PublicKey.(*rsa.PublicKey)

	// ---------------------------------------
	//code := "1234"
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter code: ")
	code, _ := reader.ReadString('\n')
	// ---------------------------------------

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
	if err = send(conn, &sec); err != nil {
		return "", errors.New("pair - secret request - send: " + err.Error())
	}

	resp, err = read(conn, &buf)
	if err != nil {
		return "", errors.New("pair - secret request - receive: " + err.Error())
	}
	fmt.Println(resp)
	if resp == nil || resp.Type != 41 {
		return "", errors.New("pair - secret request - receive: nil or invalid response")
	}

	return resp.Payload.(map[string]interface{})["secret"].(string), nil
}

type Message struct {
	ProtocolVersion int         `json:"protocol_version"`
	Payload         interface{} `json:"payload"`
	Type            int         `json:"type"`
	Status          int         `json:"status"`
}

type PariringRequest struct {
	ServiceName string `json:"service_name"`
	ClientName  string `json:"client_name"`
}

type OptionEncoding struct {
	SymbolLength int `json:"symbol_length"`
	Type         int `json:"type"`
}

type OptionsRequest struct {
	OutputEncodings []OptionEncoding `json:"output_encodings"`
	InputEncodings  []OptionEncoding `json:"input_encodings"`
	PreferredRole   int              `json:"preferred_role"`
}

type ConfigurationRequest struct {
	Encoding   OptionEncoding `json:"encoding"`
	ClientRole int            `json:"client_role"`
}

type SecretRequest struct {
	Secret string `json:"secret"`
}
