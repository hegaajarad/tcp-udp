#!/bin/bash
Portfile=( `cat ports`)
for prot in "${Portfile[@]}";do
sudo masscan -iL SCAN_IP.txt --rate=50000 -p$prot,U:$prot --excludefile No --interface eth0 --router-mac 00-c0-1d-c0-ff-ee| grep Discovered | awk {'print $6":"$4'} | awk -F/ {'print $1'} >> port_ress.txt
done
sleep 3