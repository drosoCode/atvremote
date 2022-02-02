package pairing

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"strconv"

	"github.com/drosocode/atvremote/internal/remote"
	pb "github.com/drosocode/atvremote/internal/v2/proto"
	"google.golang.org/protobuf/proto"
)

type Pairing struct {
	Buffer       []byte
	BufferSize   int
	Connection   *tls.Conn
	Certificates *tls.Certificate
	Address      string
	Port         int
}

func (p *Pairing) read(size int) ([]byte, error) {
	var back []byte
	if size != 0 {
		back = make([]byte, p.BufferSize)
		copy(back, p.Buffer[0:p.BufferSize])
	} else {
		p.BufferSize = 0
	}

	len, err := p.Connection.Read(p.Buffer)
	if err != nil {
		return nil, err
	}

	if size == 0 {
		if len == 1 {
			size = int(p.Buffer[0])
			return p.read(size)
		} else {
			size = int(p.Buffer[0])
			len = size
			p.Buffer = p.Buffer[1:]
		}
	}

	p.Buffer = append(back, p.Buffer[0:511-p.BufferSize]...)
	p.BufferSize += len
	if len < size {
		return p.read(size - len)
	}

	return p.Buffer[0:p.BufferSize], nil
}

func (p *Pairing) readMessage() (*pb.PairingMessage, error) {
	d, err := p.read(0)
	if err != nil {
		return nil, err
	}

	rsp := pb.PairingMessage{}
	err = proto.Unmarshal(d, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}

func (p *Pairing) write(data pb.PairingMessage) error {
	raw, err := proto.Marshal(&data)
	if err != nil {
		return errors.New("v2 - write message - marshal:" + err.Error())
	}

	_, err = p.Connection.Write([]byte{byte(len(raw))})
	if err != nil {
		return errors.New("v2 - write message - send 1:" + err.Error())
	}
	_, err = p.Connection.Write(raw)
	if err != nil {
		return errors.New("v2 - write message - send 2:" + err.Error())
	}

	return nil
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

	// ----------------- Pairing Message --------------------
	err = p.write(pb.PairingMessage{
		PairingRequest: &pb.PairingRequest{
			ServiceName: "com.droso.test",
			ClientName:  "test",
		},
		Status:          pb.PairingMessage_STATUS_OK,
		ProtocolVersion: 2,
	})
	if err != nil {
		return errors.New("pair - pairing - write: " + err.Error())
	}

	resp, err := p.readMessage()
	if err != nil {
		return errors.New("pair - pairing - read: " + err.Error())
	}
	if resp.Status != pb.PairingMessage_STATUS_OK {
		return errors.New("pair - pairing: invalid response")
	}

	// ----------------- Options Message --------------------
	err = p.write(pb.PairingMessage{
		PairingOption: &pb.PairingOption{
			PreferredRole: pb.RoleType_ROLE_TYPE_INPUT,
			InputEncodings: []*pb.PairingEncoding{{
				Type:         pb.PairingEncoding_ENCODING_TYPE_HEXADECIMAL,
				SymbolLength: 6,
			}},
		},
		Status:          pb.PairingMessage_STATUS_OK,
		ProtocolVersion: 2,
	})
	if err != nil {
		return errors.New("pair - options - write: " + err.Error())
	}

	resp, err = p.readMessage()
	if err != nil {
		return errors.New("pair - options - read: " + err.Error())
	}
	if resp.Status != pb.PairingMessage_STATUS_OK {
		return errors.New("pair - options: invalid response")
	}

	// ----------------- Configuration Message --------------------
	err = p.write(pb.PairingMessage{
		PairingConfiguration: &pb.PairingConfiguration{
			ClientRole: pb.RoleType_ROLE_TYPE_INPUT,
			Encoding: &pb.PairingEncoding{
				Type:         pb.PairingEncoding_ENCODING_TYPE_HEXADECIMAL,
				SymbolLength: 6,
			},
		},
		Status:          pb.PairingMessage_STATUS_OK,
		ProtocolVersion: 2,
	})
	if err != nil {
		return errors.New("pair - configuration - write: " + err.Error())
	}

	resp, err = p.readMessage()
	if err != nil {
		return errors.New("pair - configuration - read: " + err.Error())
	}
	if resp.Status != pb.PairingMessage_STATUS_OK {
		return errors.New("pair - configuration: invalid response")
	}

	return nil
}

func (p *Pairing) Secret(code string) error {
	defer p.Connection.Close()
	// ------------- Secret Request --------------------

	crt, _ := x509.ParseCertificate(p.Certificates.Certificate[0])
	serverPublicKey := p.Connection.ConnectionState().PeerCertificates[0].PublicKey.(*rsa.PublicKey)
	clientPublicKey := crt.PublicKey.(*rsa.PublicKey)
	secret := remote.GetHash(serverPublicKey, clientPublicKey, code[2:6])

	err := p.write(pb.PairingMessage{
		PairingSecret: &pb.PairingSecret{
			Secret: secret,
		},
		Status:          pb.PairingMessage_STATUS_OK,
		ProtocolVersion: 2,
	})
	if err != nil {
		return errors.New("pair - secret request - send: " + err.Error())
	}

	resp, err := p.readMessage()
	if err != nil {
		return errors.New("pair - secret request - receive: " + err.Error())
	}
	if resp.Status != pb.PairingMessage_STATUS_OK {
		return errors.New("pair - secret: invalid response")
	}

	return nil
}
