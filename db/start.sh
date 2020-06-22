#!/bin/bash
DB_USER='root'
DB_PASS='secret123'
# DB='database_name'
mysql --user="$DB_USER" --password="$DB_PASS" -Bse"use mysql;select host, user from user; ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'secret123'; FLUSH PRIVILEGES;"

# mysql -u root -p
# mysql > use mysql 
# mysql > select host, user from user;
# mysql > ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'secret123';
# mysql > FLUSH PRIVILEGES;

