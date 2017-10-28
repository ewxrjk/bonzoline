package bonzoline

import (
	"log"
	"net"
	"time"
)

// Repeatedly attempt to connect to an address and return on success.
//
// Does not discriminate between different kinds of errors (but it
// probably should).
func AwaitDialable(network, address string, deadline time.Time) (err error) {
	var conn net.Conn
	for deadline.After(time.Now()) {
		if conn, err = net.DialTimeout(network, address, time.Until(deadline)); err == nil {
			if err = conn.Close(); err != nil {
				log.Printf("closing throwaway connection: %s", err)
			}
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	return
}
