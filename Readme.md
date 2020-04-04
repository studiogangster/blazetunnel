# blazetunnel

Blazetunnel creates a secured tunnel(s) between your local machine(s) & a dedicated server.
This tunnel can be used expose your local service to the public on the internet.
   
######   Status
Currently in in beta stage. 

### Problems:

1. NAT restrictions imposed by your ISP
2. Double NAT imposed by your ISP
3. No option to upgrade to Static IP
4. Demanding a premium fee to get a static IP

### Use-cases:

1. Collaborative development & testing
2. Showcase your portfolio to the client
3. You own a high-end machine, and want to utilise it's CPU , GPU remotely (Video Rendering, ML, AI, etc.)
4. Dont't want to pay for expesive spec server, and create your own cloud
5. Bypass firewall rules (imposed by ISP, organization, etc.) to expose your webservice
6. Create your own on-the-go-lab
7. Have your own intranet available on the internet

### Why Blazetunnel:

1. Seamsless to create & deploy new apps on the go. 
2. Side-car for your docker container to directly expose the service on the internet
3. Registration, authentication, & end-to-end encryption
4. Blazetunnel uses UDP rather than TCP for tunneling (QUIC). 
So, it can bypass TCP firewall restrictions, is fast, scalable, supports multiplexing, & no head-of-line blocking.
5. Assembled with nginx, & cert-bot to make it easy to deploy standalone server, with TLS, and be scale ready
6. Custom error-page in case the exposed service is down
7. Written in GoLang (Faster,) 
8. Uses zero-copy syscall (Faster & less cpu-cycles) *if possible*.
9. Low latency ( Thanks to UDP & QUIC )
10. Multiplexing allows to serve parallel requests coming to the same service.
11. Built-in SSL Plugin (certbot)
12. End-to-end security

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

