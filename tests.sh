#!/bin/bash

echo "───────┬────────────────────────────────────────"
echo "       │ [38;5;83m[1mStarting Tests...[0m"
echo "───────┼────────────────────────────────────────"

o=0
for i in classic shorty frenzy ghost sheriff stinger spectre bucky judge bulldog guardian phantom vandal marshal operator ares odin melee outlaw;
do
   status_code=$(curl -s -o /dev/null -w "%{http_code}" "http://127.0.0.1:8080/api/v1/skin/$i")
   if test $o -le 9; then
        case $status_code in
            200)
                echo -e "   $o   │ /api/v1/skin/$i: \e[0;32m$status_code\e[0m"
                ;;
            *)
                echo -e "   $o   │ /api/v1/skin/$i: \e[0;31m$status_code\e[0m"
                ;;
            esac
            ((o++))
    else
        case $status_code in
            200)
                echo -e "   $o  │ /api/v1/skin/$i: \e[0;32m$status_code\e[0m"
                ;;
            *)
                echo -e "   $o  │ /api/v1/skin/$i: \e[0;31m$status_code\e[0m"
                ;;
            esac
            ((o++))
    fi
done

