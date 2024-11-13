/*
The MIT License (MIT)

Copyright (c) 2014 Henry

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package gosam

import (
	"log"
	"net"
	"time"
)

// Conn Read data from the connection, writes data to te connection
// and logs the data in-between.
type Conn struct {
	RWC
	conn net.Conn
}

// WrapConn wraps a net.Conn in a Conn.
func WrapConn(c net.Conn) *Conn {
	wrap := Conn{
		conn: c,
	}
	wrap.Reader = NewReadLogger("<", c)
	wrap.Writer = NewWriteLogger(">", c)
	wrap.RWC.c = c
	return &wrap
}

// LocalAddr returns the local address of the connection.
func (c *Conn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

// RemoteAddr returns the remote address of the connection.
func (c *Conn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

// SetDeadline sets the read and write deadlines associated with the connection
func (c *Conn) SetDeadline(t time.Time) error {
	log.Println("WARNING: SetDeadline() not sure this works")
	return c.conn.SetDeadline(t)
}

// SetReadDeadline sets the read deadline associated with the connection
func (c *Conn) SetReadDeadline(t time.Time) error {
	log.Println("WARNING: SetReadDeadline() not sure this works")
	return c.conn.SetReadDeadline(t)
}

// SetWriteDeadline sets the write deadline associated with the connection
func (c *Conn) SetWriteDeadline(t time.Time) error {
	log.Println("WARNING: SetWriteDeadline() not sure this works")
	return c.conn.SetWriteDeadline(t)
}
