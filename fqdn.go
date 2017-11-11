package bonzoline

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"sync"
)

// Return the local fully qualified hostname
//
// If the FQDN can't be found, a (possibly) relative hostname is
// returned.
func LocalHostname() (hostname string, err error) {
	fqdnOnce.Do(initializeFqdn)
	hostname = fqdn
	err = fqdnError
	return
}

var fqdn string
var fqdnError error
var fqdnOnce sync.Once

func initializeFqdn() {
	var path string
	if path, fqdnError = exec.LookPath("hostname"); fqdnError == nil {
		// hostname exists, try hostname -f
		var b []byte
		if b, fqdnError = Capture(path, "-f"); fqdnError == nil {
			b = bytes.TrimSuffix(b, []byte{'\n'})
			if bytes.Contains(b, []byte(" \n\t\r\f")) {
				log.Printf("%s output malformed\n", path)
			} else {
				fqdn = string(b)
				return
			}
		} else {
			log.Printf("%s failed: %s\n", path, fqdnError)
		}
	}
	if fqdn, fqdnError = os.Hostname(); fqdnError != nil {
		log.Printf("os.Hostname failed: %s\n", fqdnError)
		return
	}
	return

}
