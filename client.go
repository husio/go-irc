package irc

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// Client represents single server connection client.
type Client struct {
	conn net.Conn
	rd   *bufio.Reader
}

// Connect return client connected to given address.
func Connect(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	c := &Client{
		conn: conn,
		rd:   bufio.NewReader(conn),
	}
	return c, nil
}

// Send writes message to connected server.
//
// Because every message has to end with \r\n, if given format string do not
// end this way, those two character are written additionally.
func (c *Client) Send(format string, args ...interface{}) error {
	if _, err := fmt.Fprintf(c.conn, format, args...); err != nil {
		return err
	}
	if !strings.HasSuffix(format, "\r\n") {
		if _, err := fmt.Fprint(c.conn, "\r\n"); err != nil {
			return err
		}
	}
	return nil
}

// ReadMessage reads line from server and return parsing result.
func (c *Client) ReadMessage() (*Message, error) {
	line, err := c.rd.ReadString('\n')
	if err != nil {
		return nil, err
	}
	return ParseLine(line)
}
