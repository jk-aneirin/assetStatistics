#!/bin/bash
SNMPWALK=/usr/bin/snmpwalk
$SNMPWALK -v 2c -c wumiiiRO 10.0.0.211 .1.3.6.1.2.1.3.1.1.2>ipvsmac.txt
$SNMPWALK -v 2c -c wumiiiRO 10.0.0.241 .1.3.6.1.2.1.3.1.1.2>>ipvsmac.txt
