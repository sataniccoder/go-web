package utilgo

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strconv"
	"time"
)

// genrate admin pin code for development
// used for accsesign the admin page
// eventually a user name and password system will be made but for now we won't use that
func Gen_code() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	a := r1.Intn(1000-100) + 100

	// get new time
	s1 = rand.NewSource(time.Now().UnixNano())
	r1 = rand.New(s1)
	b := r1.Intn(1000-100) + 100

	s1 = rand.NewSource(time.Now().UnixNano())
	r1 = rand.New(s1)
	c := r1.Intn(1000-100) + 100

	return strconv.Itoa(a) + "-" + strconv.Itoa(b) + "-" + strconv.Itoa(c)

}

// https enebler
func Gen_cert() {
	// gen the cert using openssl and save them too full-cert.crt, private-key.key files
	// genrate a conifg file for the certificats in certs/

	fmt.Println("[*] genrating certs (make sure to edit the cert.cnf file)")
	cmd := exec.Command("openssl", "genrsa", "-out", "certs/server.key", "2048")
	cmd.Run()
	cmd = exec.Command("openssl", "req", "-config", "certs/cert.cnf", "-new", "-x509", "-sha256", "-key", "certs/server.key", "-out", "certs/server.crt", "-days", "365")
	cmd.Run()
}
