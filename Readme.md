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

###  Customize docker-compose.yamL
#### Services: 
##### server (blazetunnel server): Needs to be run on exit node (internet)
##### client (blazetunnel client): Node may or may not be behind NAT
##### mockserver (A sample server that would be exposed to the internet)

`*TODO:* Add instructions to use blazetunnel with docker`




##  Start server
```
docker-compose up server
```


##  Start client

```
docker-compose up client
```




Go build first.

*Server*

```
# run qxpose as server mode with the configured sub domain as poniesareaweso.me
# and the idle time out for QUIC sessions as an hour (default is 1/2 hour)
qxpose server --domain poniesareaweso.me -i 3600
```

```
# run qxpose as client mode with the following options
#  1. Tunnel server: to the locally running one
#  2. Local: Which local server/TCP address to proxy to public.
#  3. Idle Timeout: idle time out for QUIC sessions as an hour (default is 1/2 hour)
qxpose client --tunnel "localhost:2723" --local "localhost:8100" -i 3600
```

The client spits out a new hostname for the tunnel. (something like fb6b5b1749f59e70.poniesareaweso.me)
For locally testing, edit the /etc/hosts to point
the host to 127.0.0.1. 

Something like this.
```
127.0.0.1   fb6b5b1749f59e70.poniesareaweso.me
```

Now try the address in the browser (insecure) or cURL (with `-k` flag).