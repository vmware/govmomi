package simulator

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	kmip "github.com/smira/go-kmip"
)

type client struct {
	kclient kmip.Client
}

var kmipclient *client

func initClient() (*client, error) {
	c := &client{}
	/*
		// Gets the key vault
		keyVaultName, err := common.GetEnvVar("CLIENT_KV_URL")
		if err != nil {
			log.Println("No client keyvault in env")
		}
		certname, err := common.GetEnvVar("CERT_NAME")
		if err != nil {
			log.Println("missing certname in env")
		}
	*/
	// GetCertFromVault(keyVaultName, certname)
	c.kclient.Endpoint = "127.0.0.1:5696"
	c.kclient.TLSConfig = &tls.Config{}
	kmip.DefaultClientTLSConfig(c.kclient.TLSConfig)
	// Skip server cert verification
	c.kclient.TLSConfig.InsecureSkipVerify = true
	// c.kclient.TLSConfig.RootCAs = certCAPool
	/*
		c.kclient.TLSConfig.Certificates = []tlCertificate{
			{
				Certificate: [][]byte{certClientCert.Raw},
				PrivateKey:  certClientKey,
			},
		}
	*/

	c.kclient.ReadTimeout = time.Minute
	c.kclient.WriteTimeout = time.Minute

	return c, nil

}

func sendCreateKey(w http.ResponseWriter, req *http.Request) {
	// TODO: debug later, initial connection fails because of no command
	fmt.Fprintf(w, "sendCreateKey\n")
	err := kmipclient.kclient.Connect()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Client connected\n")
	resp, err := kmipclient.CreateKey()
	fmt.Fprintf(w, "sendCreateKey: resp %v, error %v\n", resp, err)
}

func sendGetKey(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "sendGetKey\n")
	err := kmipclient.kclient.Connect()
	if err != nil {
		panic(err)
	}
	id := req.URL.Query().Get("key")
	fmt.Printf("Client connected\n")
	fmt.Println("Key ID : " + id)
	resp, err := kmipclient.GetKey(id)
	fmt.Fprintf(w, "sendGetKey: resp %v, error %v\n", resp, err)
}

func sendGetAttributes(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "sendGetAttributes\n")
	err := kmipclient.kclient.Connect()
	if err != nil {
		panic(err)
	}
	id := req.URL.Query().Get("key")
	fmt.Printf("Client connected\n")
	fmt.Println("Key ID : " + id)
	resp, err := kmipclient.GetKeyAttributes(id)
	fmt.Fprintf(w, "sendGetAttributes: resp %v, error %v\n", resp, err)
}

func sendDeleteKey(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "sendDeleteKey\n")
	err := kmipclient.kclient.Connect()
	if err != nil {
		panic(err)
	}
	id := req.URL.Query().Get("key")
	fmt.Printf("Client connected\n")
	fmt.Println("Key ID : " + id)
	resp, err := kmipclient.DeleteKey(id)
	fmt.Fprintf(w, "sendDeleteKey: resp %v, error %v\n", resp, err)
}

func main() {

	kmipclient, _ = initClient()
	fmt.Printf("Client init \n")
	// connect to server.
	err := kmipclient.kclient.Connect()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Client connected\n")

	mux := http.NewServeMux()
	mux.HandleFunc("/create", sendCreateKey)
	mux.HandleFunc("/get", sendGetKey)
	mux.HandleFunc("/getattributes", sendGetAttributes)
	mux.HandleFunc("/delete", sendDeleteKey)

	server := &http.Server{
		Addr:           ":8090",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
