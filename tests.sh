#!/bin/bash

echo "[38;5;83m[1mStarting Tests...[0m"

for i in classic shorty frenzy ghost sheriff stinger spectre bucky judge bulldog guardian phantom vandal marshall operator ares odin melee outlaw;
do
   status_code=$(curl -s -o /dev/null -w "%{http_code}" "http://127.0.0.1:8080/api/v1/$i")
   case $status_code in
       200)
           echo -e "http://127.0.0.1:8080/api/v1/$i: \e[0;32m$status_code\e[0m"
           ;;
       *)
           echo -e "http://127.0.0.1:8080/api/v1/$i: \e[0;31m$status_code\e[0m"
           ;;
   esac
done

