package util

import (
	"fmt"
	"log"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	intFace        = "eth0"
	devFound       = false
	snapshotLength = int32(1600)
	timeout        = pcap.BlockForever
	promisc        = false //capture only certain packets
	filter         = "tcp and port 80"
)

func PacketSniffer() {
	devices, err := pcap.FindAllDevs()

	if err != nil {
		log.Fatal(err)
	}
	for _, dev := range devices {
		// fmt.Println(i, " ", dev.Name)
		// for _, addr := range dev.Addresses {
		// 	fmt.Println("\t IP: ", addr.IP)
		// 	fmt.Println("\t NetMask: ", addr.Netmask)
		// }
		if dev.Name == intFace {
			devFound = true
		}

		if !devFound {
			log.Fatal("no device found")
		}

		handle, err := pcap.OpenLive(intFace, snapshotLength, promisc, timeout)

		if err != nil {
			log.Fatal(err)
		}
		defer handle.Close()

		err = handle.SetBPFFilter(filter)

		if err != nil {
			log.Fatal(err)
		}

		source := gopacket.NewPacketSource(handle, handle.LinkType())

		for pkt := range source.Packets() {
			appLayer := pkt.ApplicationLayer()

			if appLayer == nil {
				continue
			}
			payload := appLayer.Payload()
			searchArr := []string{"name", "username", "pass"}

			for _, s := range searchArr {
				index := strings.Index(string(payload), s)

				if index != -1 {
					fmt.Println(string(payload[index : index+100]))
				}
			}
		}
	}
}
