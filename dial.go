package goSam

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
)

// DialContext implements the net.DialContext function and can be used for http.Transport
func (c *Client) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	c.oml.Lock()
	defer c.oml.Unlock()
	errCh := make(chan error, 1)
	connCh := make(chan net.Conn, 1)
	go func() {
		if conn, err := c.Dial(network, addr); err != nil {
			errCh <- err
		} else if ctx.Err() != nil {
			log.Println(ctx)
			errCh <- ctx.Err()
		} else {
			connCh <- conn
		}
	}()
	select {
	case err := <-errCh:
		return c.SamConn, err
	case conn := <-connCh:
		return conn, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (c *Client) dialCheck(addr string) (int32, bool) {
	if c.lastaddr == "invalid" {
		fmt.Println("Preparing to dial new address.")
		return c.NewID(), true
	} else if c.lastaddr != addr {
		fmt.Println("Preparing to dial next new address.")
		return c.NewID(), true
	}
	return c.id, false
}

// Dial implements the net.Dial function and can be used for http.Transport
func (c *Client) Dial(network, addr string) (net.Conn, error) {
	c.ml.Lock()
	defer c.ml.Unlock()
	portIdx := strings.Index(addr, ":")
	if portIdx >= 0 {
		addr = addr[:portIdx]
	}
	addr, err := c.Lookup(addr)
	if err != nil {
		return nil, err
	}

	c.id, _ = c.dialCheck(addr)
	c.destination, err = c.CreateStreamSession(c.id, c.destination)
	if err != nil {
		c.id += 1
		c, err = c.NewClient()
		if err != nil {
			return nil, err
		}
		c.destination, err = c.CreateStreamSession(c.id, c.destination)
		if err != nil {
			return nil, err
		}
	}
	c, err = c.NewClient()
	if err != nil {
		return nil, err
	}
	c.lastaddr = addr

	err = c.StreamConnect(c.id, addr)
	if err != nil {
		return nil, err
	}
	return c.SamConn, nil
}
