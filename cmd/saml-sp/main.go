package main

import (
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/url"

	"github.com/crewjam/saml/samlsp"
)

func hello(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello, %w!", samlsp.AttributeFromContext(r.Context(), "displayName"))
	s := samlsp.SessionFromContext(r.Context())
	if s == nil {
		return
	}
	sa, ok := s.(samlsp.SessionWithAttributes)
	if !ok {
		return
	}
	fmt.Fprintf(w, "UserInfo\n")
	as := sa.GetAttributes()
	for k, v := range as {
		fmt.Printf("key: %s value: %s\n", k, v)
		fmt.Fprintf(w, "key: %s value: %s\n", k, v)
	}
}

func main() {
	keyPair, err := tls.LoadX509KeyPair("sp.cert", "sp.key")
	if err != nil {
		panic(err)
	}
	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		panic(err)
	}
	idpMetadataURL, err := url.Parse("http://localhost:8080/realms/sample/protocol/saml/descriptor")
	if err != nil {
		panic(err)
	}
	idpMetadata, err := samlsp.FetchMetadata(context.Background(), http.DefaultClient, *idpMetadataURL)
	if err != nil {
		panic(err)
	}
	rootURL, err := url.Parse("http://localhost:8000")
	if err != nil {
		panic(err)
	}
	samlSP, _ := samlsp.New(samlsp.Options{
		EntityID:    "sample-client",
		URL:         *rootURL,
		Key:         keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate: keyPair.Leaf,
		IDPMetadata: idpMetadata,
		SignRequest: true,
	})

	app := http.HandlerFunc(hello)
	http.Handle("/hello", samlSP.RequireAccount(app))
	http.Handle("/saml/", samlSP)
	http.ListenAndServe(":8000", nil)
}
