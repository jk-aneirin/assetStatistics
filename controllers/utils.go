package controllers

import (
	"regexp"
	"strconv"
	"strings"
	"net"
	"time"
	"fmt"

	g "github.com/soniah/gosnmp"
)

func IsMac(key string) bool {
	if m, _ := regexp.MatchString("^[0-9a-fA-F]{2}([:-][0-9a-fA-F]{2}){5}$", key); !m {
		return false
					
	}
		return true
}

func IsIp(key string) bool {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", key); !m {
				return false
					
	}
		return true
}

func IptoInt(ip net.IP) int64 {
	bits := strings.Split(ip.String(), ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}

func GetMac(key string) string {
	g.Default.Target = "192.168.1.1" //backbone switch ip address
	g.Default.Community = "public"  //SNMP community string
	g.Default.Timeout = time.Duration(10 * time.Second)
	g.Default.Connect()

	defer g.Default.Conn.Close()
	
	IpInt := IptoInt(net.ParseIP(key))
	oid := make([]string,1)

	switch  {
		// 10.1.80.1-10.1.80.255
		case 167858177 < IpInt && IpInt < 167858431:
			s := []string{"1.3.6.1.2.1.3.1.1.2.56.1",key}
			oid[0] = strings.Join(s,".")
		// 10.1.96.1-10.1.96.255
		case 167862273 < IpInt && IpInt < 167862527:
			s := []string{"1.3.6.1.2.1.3.1.1.2.71.1",key}
			oid[0] = strings.Join(s,".")
		default:
			oid[0] = ""
	}

	result,err := g.Default.Get(oid)
	if err != nil {
		fmt.Println(err)
	}

	variable := result.Variables[0]
	switch variable.Type {
	case g.OctetString:
		var mac net.HardwareAddr
		mac = variable.Value.([]byte)
		return mac.String()
	default:
		return ""
	} 
}
