# blazetunnel

Blazetunnel creates a secured tunnel(s) between your local machine(s) & a dedicated server.
This tunnel can be used expose your local service to the public on the internet.
   
######   Status
Currently in in beta stage. 

### Problems:

NAT restrictions imposed by your ISP
Double NAT imposed by your ISP
No option to upgrade to Static IP
Demanding a premium fee to get a static IP

### Use-cases:

Collaborative development & testing
Showcase your portfolio to the client
You own a high-end machine, and want to utilise it's CPU , GPU remotely (Video Rendering, ML, AI, etc.)
Dont't want to pay for expesive spec server, and create your own cloud
Bypass firewall rules (imposed by ISP, organization, etc.) to expose your webservice
Create your own on-the-go-lab
Have your own intranet available on the internet

### Why Blazetunnel:

Seamsless to create & deploy new apps on the go. 
Side-car for your docker container to directly expose the service on the internet
Registration, authentication, & end-to-end encryption
Blazetunnel uses UDP rather than TCP for tunneling (QUIC). 
So, it can bypass TCP firewall restrictions, is fast, scalable, supports multiplexing, & no head-of-line blocking.
Assembled with nginx, & cert-bot to make it easy to deploy standalone server, with TLS, and be scale ready
Custom error-page in case the exposed service is down
Written in GoLang (Faster,) 
*Uses zero-copy syscall (Faster & less cpu-cycles) if possible.
Low latency ( Thanks to UDP & QUIC )
Multiplexing allows to serve parallel requests coming to the same service.
Built-in SSL Plugin (certbot)
End-to-end security

### Installation:

1. Docker container
2. Sidecar
3. Build from source
4. npm i blazetunnel




### Architechture

```
                            Blazetunnel architecture
 _______                       _________________            ________
|       | Behind NAT          |     Exposed     |          |        |
|       |                     |                 |          |        |        (HTTP/HTTPS)
| local |    <--- QUIC --->   |     Tunnel      | <--TLS-->|  Nginx |    <--Browser/cURL/openssl-->     
|       |                     |                 |          |        |
|_______|                     |_________________|          |________|



         __________              __________             
        |          |            |          |    
        |          |            |          |    
        |  docker  |            |  dokcer  |   
        |    I     |            |    II    |        
        |__________|            |__________|     

        Service I (service)     Service II (quic)
        Expose: 8080            Port: 80 : service:8080      

         _________ 
        |         |             REQUEST
        | CLIENT  |   ----> quic:80 ---> service:8080|
        |         |   <------------------------------|
        |         |               RESPONSE
        |_________|



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

