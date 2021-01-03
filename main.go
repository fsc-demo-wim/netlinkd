package main

import (
	"bytes"
	"encoding/json"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/fsc-demo-wim/netlinkd/netlinktypes"

	log "github.com/sirupsen/logrus"
	"github.com/vishvananda/netlink"
)

func netlinkServer(c net.Conn) {
	for {
		buf := make([]byte, 2048)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		msg := buf[0:nr]
		log.Info("Server got:", string(msg))

		nll, err := netlink.LinkList()
		if err != nil {
			log.Error(err)
		}

		newLinkData := make([]netlinktypes.Link, 0)

		for _, l := range nll {
			newLink := netlinktypes.Link{}
			newLink.Index = l.Attrs().Index
			newLink.Name = l.Attrs().Name
			newLink.HardwareAddr = l.Attrs().HardwareAddr.String()
			newLink.MTU = l.Attrs().MTU
			newLink.OperState = l.Attrs().OperState.String()
			newLink.ParentIndex = l.Attrs().ParentIndex
			newLink.MasterIndex = l.Attrs().MasterIndex
			for _, vf := range l.Attrs().Vfs {
				newVf := netlinktypes.VfInfo{}
				newVf.ID = vf.ID
				newVf.Vlan = vf.Vlan
				newLink.Vfs = append(newLink.Vfs, newVf)
			}
			newLinkData = append(newLinkData, newLink)
		}

		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(newLinkData)

		log.Infof(string(reqBodyBytes.Bytes()))

		_, err = c.Write(reqBodyBytes.Bytes())
		if err != nil {
			log.Fatal("Write: ", err)
		}
	}
}

func main() {
	log.Infof("Starting netlink-server...")
	syscall.Unlink("/tmp/netlink.sock")
	l, err := net.Listen("unix", "/tmp/netlink.sock")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	go func(ln net.Listener, c chan os.Signal) {
		sig := <-c
		log.Printf("Caught signal %s: shutting down.", sig)
		ln.Close()
		os.Exit(0)
	}(l, sigc)

	for {
		fd, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go netlinkServer(fd)
	}
}
