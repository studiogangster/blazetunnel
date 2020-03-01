const WebSocket = require('ws');
const requests = require('request');


requests.get('https://buzz.gamepind.com/buzz_server/?EIO=3&transport=polling&t=N2GEOet',


    {
        headers:
        {
            'Cookie':

   '_ga=GA1.2.611139214.1580129948; _fbp=fb.1.1580129948503.1185898356; _ga=GA1.3.611139214.1580129948; _gid=GA1.2.1633700889.1582928890; _gid=GA1.3.1633700889.1582928890; _gat_UA-117948010-1=1; AWSALB=nsN+bgI7OcIrZt+frrr4GmHZ/nFpnhm40OD9HIzLUvzd9VjMoLHKjTbXF7pSzj5AYBl84R4m/d9IMYOUhM3I6qLuOrqBSxbbqDe+vJ0nqRMTQQXIuT++TLt8nwkK; AWSALBCORS=nsN+bgI7OcIrZt+frrr4GmHZ/nFpnhm40OD9HIzLUvzd9VjMoLHKjTbXF7pSzj5AYBl84R4m/d9IMYOUhM3I6qLuOrqBSxbbqDe+vJ0nqRMTQQXIuT++TLt8nwkK; io=pmrdgrNwja2Eu945BNc5'     },
    },
    function (error, response, body) {
        let sid = body.indexOf("sid") + 4 + 2
        let _sid = body.indexOf('","upg')
        let SID = body.substring(sid, _sid)
        console.log('body', SID);

        let url = `wss://buzz.gamepind.com/buzz_server/?EIO=3&transport=websocket&sid=${SID}`

        connect(url)


    });






function connect(url) {



    const client = new WebSocket(url);

    client.on('open', (connected) => {


        let msg = '422["getMyInfo",{"game_type":"cash"}]'
        console.log('connection is open now', connected)

        client.send('2probe')
        client.send('5')
        client.send(msg)
    });
    client.on('ping', console.log);
    client.on('message', console.log);
    client.on('error', console.log);


    client.on('close', function clear() {
        console.log('closde')
        clearTimeout(this.pingTimeout);
    });

}
