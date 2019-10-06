package xmpp

import (
	"crypto/tls"
	"io"
	"os"
)

type Config struct {
	Address        string
	Jid            string
	parsedJid      *Jid // For easier manipulation
	Credential     Credential
	StreamLogger   *os.File // Used for debugging
	Lang           string   // TODO: should default to 'en'
	ConnectTimeout int      // Client timeout in seconds. Default to 15
	// tls.Config must not be modified after having been passed to NewClient. Any
	// changes made after connecting are ignored.
	TLSConfig *tls.Config
	// Insecure can be set to true to allow to open a session without TLS. If TLS
	// is supported on the server, we will still try to use it.
	Insecure      bool
	CharsetReader func(charset string, input io.Reader) (io.Reader, error) // passed to xml decoder
}
