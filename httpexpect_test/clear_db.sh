#!/bin/bash
# set -v
db_name=""
if [ $# -eq 1 ];then
    db_name=$1
fi
echo $db_name
mysql -uroot -p123456 -N -s information_schema -e "SELECT CONCAT('TRUNCATE TABLE province_warehouse_test.',TABLE_NAME,';') FROM information_schema.TABLES WHERE TABLE_SCHEMA='$db_name';" | mysql -uroot -p123456 $db_name
