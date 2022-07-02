/* GenX509Cert
 */

package crypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/gob"
	"encoding/pem"
	"log"
	"math/big"
	"os"
	"time"
)

var certsFP = "certs/"

//X509Config data
type X509Config struct {
	CommonName string
	OrgName    string
	DomainName string
	KeyName    string
	CertPath   string
}

func priKeyName(kn string) string {
	return kn + "-private.key"
}

func pubKeyName(kn string) string {
	return kn + "-public.key"
}

func priPemName(kn string) string {
	return kn + "-private.pem"
}

func pubPemName(kn string) string {
	return kn + "-public.pem"
}

//GenerateX509Certificate generates certificate for secure tls connection
//Have to generate key pair first before using this
func GenerateX509Certificate(config X509Config) error {
	err := PKeys(config)
	if err != nil {
		return err
	}
	random := rand.Reader

	var key rsa.PrivateKey
	loadKey(certsFP+priKeyName(config.KeyName), &key)

	now := time.Now()
	then := now.Add(60 * 60 * 24 * 365 * 1000 * 1000 * 1000) // one year
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName:   config.CommonName,
			Organization: []string{config.OrgName},
		},
		//    NotBefore: time.Unix(now, 0).UTC(),
		//    NotAfter:  time.Unix(now+60*60*24*365, 0).UTC(),
		NotBefore: now,
		NotAfter:  then,

		SubjectKeyId: []byte{1, 2, 3, 4},
		KeyUsage:     x509.KeyUsageCertSign | x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,

		BasicConstraintsValid: true,
		IsCA:                  true,
		DNSNames:              []string{config.CommonName, config.DomainName},
	}
	derBytes, err := x509.CreateCertificate(random, &template,
		&template, &key.PublicKey, &key)
	if err != nil {
		return err
	}

	certCerFile, err := os.Create(certsFP + config.CommonName + ".cer")
	if err != nil {
		return err
	}
	certCerFile.Write(derBytes)
	certCerFile.Close()

	certPEMFile, err := os.Create(certsFP + config.CommonName + ".pem")
	if err != nil {
		return err
	}
	pem.Encode(certPEMFile, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certPEMFile.Close()

	keyPEMFile, err := os.Create(certsFP + config.KeyName + "-private.pem")
	if err != nil {
		return err
	}
	pem.Encode(keyPEMFile, &pem.Block{Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(&key)})
	keyPEMFile.Close()
	return err
}

func loadKey(fileName string, key interface{}) {
	inFile, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
		return
	}
	decoder := gob.NewDecoder(inFile)
	err = decoder.Decode(key)
	if err != nil {
		log.Println(err)
		return
	}
	inFile.Close()
}

//PKeys generate keys for certi
func PKeys(config X509Config) error {
	keyName := config.KeyName
	reader := rand.Reader
	bitSize := 2048
	key, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		return err
	}

	//fmt.Println("Private key primes", key.Primes[0].String(), key.Primes[1].String())
	//fmt.Println("Private key exponent", key.D.String())

	publicKey := key.PublicKey
	//fmt.Println("Public key modulus", publicKey.N.String())
	//fmt.Println("Public key exponent", publicKey.E)

	err = saveCertGobKey(priKeyName(keyName), key)
	err = saveCertGobKey(pubKeyName(keyName), publicKey)
	err = saveCertPEMKey(priPemName(keyName), key)
	return err
}

func saveCertGobKey(fileName string, key interface{}) error {
	outFile, err := os.Create(certsFP + fileName)
	if err != nil {
		return err
	}
	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	if err != nil {
		return err
	}
	outFile.Close()
	return err
}

func saveCertPEMKey(fileName string, key *rsa.PrivateKey) error {

	outFile, err := os.Create(certsFP + fileName)
	if err != nil {
		return err
	}

	var privateKey = &pem.Block{Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key)}

	pem.Encode(outFile, privateKey)

	outFile.Close()
	return err
}
