# blazetunnel

Experimenting QUIC protocol to create P2P tunnel between private intranets and expose all services within interconnected network to internet without Static IP.
   
######   Status
Currently in experimental stage. 


### Architechture

```
                      _________________
 Behind NAT          |     Exposed     |
                     |                 |
Local <--- QUIC ---> |     Tunnel      | <--- TLS ---> Browser/cURL/openssl
                     |                 |
                     |_________________|
```

### Why QUIC?

Utilise single socket to serve parallel requests, using QUIC's multilexing. 
Eliminate Head of line blocking in TCP
Elimiate / Reduce Round-trip delay time

### Details

##### docker & docker-compose are required
```
Customize docker-compose.yamL
    Services: 
        server (blazetunnel server): Needs to be run on exit node (internet)
        client (blazetunnel client): Node may or may not be behind NAT
        mockserver (A sample server that would be exposed to the internet)
```
`*TODO:* Add instructions to use blazetunnel with docker`




##  Instructions
```

Run 'Server' on internet facing node.
Run 'Client' on local machine
Create A record for *.domainname & domainname pointing to 'Server' node.

Modify .env file in the root directory as follows:
DOMAIN_NAME => Domain name of the internet facing server / exit node 
SERVICE_NAME => Service name is the subdomain, that would be used to access the local server 

```

##  Start server
```
docker-compose up server
```


##  Start client

```
docker-compose up client
```

#### Go to https://{service_name}.{domain_name} to access local server! Ex: https://quic.meddler.xyz

