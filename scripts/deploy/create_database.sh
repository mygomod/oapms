#!/usr/bin/env bash
mysql="/home/usr/mysql/bin/mysql -h127.0.0.1 -P3306 -uroot -proot"
sql="CREATE DATABASE oapms DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;"
$mysql -e "$sql"

