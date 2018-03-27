# go-deliver

Go-deliver is a payload delivery tool coded in Go. This is the first version and other features will be added in the future.

## Installation
To use go-deliver without installing Go and the required dependencies you can download the precompiled binaries. 

If you want to compilefrom source:

1) Clone this repository.
2) Install the required dependecies.
3) Compile and run.

## Dependecies
* https://github.com/chzyer/readline
* https://github.com/gorilla/mux
* https://github.com/mattn/go-sqlite3
* https://github.com/olekukonko/tablewriter
* http://gopkg.in/gcfg.v1

## Configuration
Currently go-deliver supports only http server for payload delivery. More server types will be added later.
To change the port of the http server you can modify the ```config.conf```.

Sample configuration :

```
[http]
port = 8000
template403 = template/403.html
template404 = template/403.html

[https]
port = 8443
template403 = template/403.html
template404 = template/403.html
```
The only configuration that can be changed here is the ```port``` under http. The other options are for features that will be added later.
There is a lot of code that has been created for future versions so you can ignore them.

## Usage
The main logic behind go-deliver is to deliver different type of payloads to specific IP/Subnet address or block access for specific IP/Subnet.

Currently you have two types of objects in go-deliver.
* Payload - Used to define an object to deliver.
* Host - Used to define an object to combine with Payload object.

### Payload object commands
* Add - Add a new payload to database.
* Delete - Delete a payload from the database.
* List - List all the payloads on the database.

### Payload types
```mshta  regsrv32  powershell  javascript  html  text  exe```

### Payload Options
* Id - ID of the payload on the database. (Unchangeable)
* Name - Name of the payload.
* Content Type - Content Type that will be sent as a header.
* Host Blacklist - The name of a Host object to be used as a blacklist.
* Host Whitelist - The name of a Host object to be used as a whitelist.
* Data File - Location of a file to be delivered.
* Data B64 - B64 encoded data to be delivered.
* Ptype - Payload type.
* Guid - Unique identifier for every payload.

Note : ``` If no whitelist or blacklist is specified the payload will get delivered to anyone with the correct URL.```

### Host object commands
* Add - Add a new host object to the database.
* Delete - Delete a host object from the database.
* List - List all the host objects on the database.

### Host options
* Id - ID of the host object on the database. (Unchangeable)
* Name - Name of the host object.
* Htype - Host object type. It can be ```IP``` or ```Subnet```

## TODO
* Add more types of servers.
* Add templates and the ability to generate payloads.
* Add one-liner for every payload type.

Suggestions ???

## Screenshots
### Payload Creation

![alt text](https://raw.githubusercontent.com/0x09AL/go-deliver/master/screenshot/payload_creation.png "Payload Creation")

### Payload List

![alt text](https://raw.githubusercontent.com/0x09AL/go-deliver/master/screenshot/payload_list.png "Payload List")

### Payload Deliver

![alt text](https://raw.githubusercontent.com/0x09AL/go-deliver/master/screenshot/payload_deliver.png "Payload Deliver")

### Host Creation

![alt text](https://raw.githubusercontent.com/0x09AL/go-deliver/master/screenshot/host_creation.png "Payload List")

### Payload with Black List

![alt text](https://raw.githubusercontent.com/0x09AL/go-deliver/master/screenshot/blacklist_example.png "Black List Example")

