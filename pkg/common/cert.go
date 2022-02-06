package common

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"math/big"
	"net"
	"os"
	"time"
)

func CreateCertificate(name string, hosts []string, certPath string, keyPath string) error {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return errors.New("Failed to generate private key: " + err.Error())
	}

	keyUsage := x509.KeyUsageDigitalSignature
	keyUsage |= x509.KeyUsageKeyEncipherment

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return errors.New("Failed to generate serial number: " + err.Error())
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization:       []string{"O"},
			OrganizationalUnit: []string{"OU"},
			Country:            []string{"CNT"},
			Locality:           []string{"LOC"},
			Province:           []string{"ST"},
			CommonName:         name,
		},

		NotBefore: time.Now(),
		NotAfter:  time.Date(2099, 0, 0, 0, 0, 0, 0, time.UTC),

		KeyUsage:              keyUsage,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, h)
		}
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return errors.New("Failed to create certificate: " + err.Error())
	}

	certOut, err := os.Create(certPath)
	if err != nil {
		return errors.New("Failed to open cert.pem for writing: " + err.Error())
	}
	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		return errors.New("Failed to write data to cert.pem: " + err.Error())
	}
	if err := certOut.Close(); err != nil {
		return errors.New("Error closing cert.pem: " + err.Error())
	}

	keyOut, err := os.OpenFile(keyPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return errors.New("Failed to open key.pem for writing: " + err.Error())
	}
	privBytes, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return errors.New("Unable to marshal private key: " + err.Error())
	}
	if err := pem.Encode(keyOut, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}); err != nil {
		return errors.New("Failed to write data to key.pem: " + err.Error())
	}
	if err := keyOut.Close(); err != nil {
		return errors.New("Error closing key.pem: " + err.Error())
	}

	return nil
}
