# What doing project

this project copy json from server
and write to rabbitmq
then copy from rabbitmq to PostgreSQL

# get started

write to .env file data:

### Info about copy

serverPort = Port server from copy
serverHost = Host server from copy

### Info about paste

dbname - DB name to copy 
host - Host of DB 
port - Port of DB
user - Username DB
password - Password for Username

tablename - Table name to copy