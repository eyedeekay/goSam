//go:build nettest
// +build nettest

package gosam

import "testing"

import (
	"fmt"
	//	"math"
	//	"math/rand"
	//	"time"
	//"log"
	"net/http"
	//"github.com/go-i2p/sam3/helper"
	//"github.com/go-i2p/i2pkeys"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

/*func TestCompositeClient(t *testing.T) {
	listener, err := sam.I2PListener("testservice"+fmt.Sprintf("%d", rand.Int31n(math.MaxInt32)), "127.0.0.1:7656", "testkeys")
	if err != nil {
		t.Fatalf("Listener() Error: %q\n", err)
	}
	defer listener.Close()
	http.HandleFunc("/", HelloServer)
	go http.Serve(listener, nil)

	listener2, err := sam.I2PListener("testservice"+fmt.Sprintf("%d", rand.Int31n(math.MaxInt32)), "127.0.0.1:7656", "testkeys2")
	if err != nil {
		t.Fatalf("Listener() Error: %q\n", err)
	}
	defer listener2.Close()
	//	http.HandleFunc("/", HelloServer)
	go http.Serve(listener2, nil)

	listener3, err := sam.I2PListener("testservice"+fmt.Sprintf("%d", rand.Int31n(math.MaxInt32)), "127.0.0.1:7656", "testkeys3")
	if err != nil {
		t.Fatalf("Listener() Error: %q\n", err)
	}
	defer listener3.Close()
	//	http.HandleFunc("/", HelloServer)
	go http.Serve(listener3, nil)

	sam, err := NewClientFromOptions(SetDebug(false))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	tr := &http.Transport{
		Dial: sam.Dial,
	}
	client := &http.Client{Transport: tr}
	defer sam.Close()
	x := 0
	for x < 15 {
		time.Sleep(time.Second * 2)
		t.Log("waiting a little while for services to register", (30 - (x * 2)))
		x++
	}
	go func() {
		resp, err := client.Get("http://" + listener.Addr().(i2pkeys.I2PAddr).Base32())
		if err != nil {
			t.Fatalf("Get Error test 1: %q\n", err)
		}
		defer resp.Body.Close()
	}()
	//time.Sleep(time.Second * 15)
	go func() {
		resp, err := client.Get("http://" + listener2.Addr().(i2pkeys.I2PAddr).Base32())
		if err != nil {
			t.Fatalf("Get Error test 2: %q\n", err)
		}
		defer resp.Body.Close()
	}()
	//time.Sleep(time.Second * 15)
	go func() {
		resp, err := client.Get("http://" + listener3.Addr().(i2pkeys.I2PAddr).Base32())
		if err != nil {
			t.Fatalf("Get Error test 3: %q\n", err)
		}
		defer resp.Body.Close()
	}()

	time.Sleep(time.Second * 45)
}*/

func TestClientHello(t *testing.T) {
	client, err := NewClientFromOptions(SetDebug(false))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	t.Log(client.Base32())
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestNewDestination(t *testing.T) {
	client, err := NewClientFromOptions(SetDebug(false))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	t.Log(client.Base32())
	if s, p, err := client.NewDestination(SAMsigTypes[3]); err != nil {
		t.Error(err)
	} else {
		t.Log(s, p)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}
