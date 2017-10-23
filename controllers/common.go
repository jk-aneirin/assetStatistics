package controllers

import (
	"regexp"
	"io"
	"os"
	"strings"
	"fmt"
	"bufio"
//	"log"
//	"github.com/jtblin/go-ldap-client"
//	"github.com/astaxie/beego/context"
)

//var FilterUser = func(ctx *context.Context) {
//	client := &ldap.LDAPClient{
//	Base:         "dc=example,dc=com",
//	Host:         "ldap.example.com",
//	Port:         636,
//	UseSSL:       true,
//	BindDN:       "uid=xl,cn=users,cn=accounts,dc=example,dc=com",
//	BindPassword: "password",
//	UserFilter:   "(uid=%s)",
//	GroupFilter: "(memberUid=%s)",
//	Attributes:   []string{"givenName", "sn", "mail", "uid"},
//	}
//
//	client.ServerName = "ldap.example.com"
//
//	defer client.Close()
//
//	ok,_,_ := client.Authenticate("xl","password")
//	if !ok {
//		ctx.Redirect(302,"/")
//	}
//
//}

func IsMacAddr(mac string) bool {
	if m ,_ := regexp.MatchString("^[0-9a-fA-F]{2}([:-][0-9a-fA-F]{2}){5}$",mac); !m {
		return false
	} 
	return true
}

func IsIpAddr(ip string) bool {
	if m ,_ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$",ip); !m {
		return false
	} 
	return true
}

func GetMacFromIp(ip string) string {
	inputFile,inputError := os.Open("/home/xuliang/assetStatistics/src/ipvsmac.txt")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
		         "Does the file exist?\n" +  "Have you got acces to it?\n")
		return ""// exit the function on error
	}

	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString,readerError := inputReader.ReadString('\n')
		match,_ := regexp.MatchString(ip,inputString)
		if match {
			targetlens := len(inputString)
			tgt := []string{inputString[targetlens-19:targetlens-17],inputString[targetlens-16:targetlens-14],
			inputString[targetlens-13:targetlens-11],inputString[targetlens-10:targetlens-8],
			inputString[targetlens-7:targetlens-5],inputString[targetlens-4:targetlens-2]}
			fmt.Printf(strings.Join(tgt,"-"))
			return strings.Join(tgt,"-")
		}
		if readerError == io.EOF {
			return ""
		}
	}

}
