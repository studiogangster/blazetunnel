const WebSocket = require('ws');
const requests = require('request');


requests.get('https://buzz.gamepind.com/buzz_server/?EIO=3&transport=polling&t=N26kHac',


    {
        headers:
        {
            'Cookie':

                '_ga=GA1.2.611139214.1580129948; _fbp=fb.1.1580129948503.1185898356; _gid=GA1.2.1178612749.1582783048; _ga=GA1.3.611139214.1580129948; _gid=GA1.3.1178612749.1582783048; _gat_UA-117948010-1=1; AWSALB=TA9XJyVeQu1gEcXI6qqUpPjYT06y3DU8Ar78haKaPNnr6owTbNWUrZPphyg1g0DzHmwOxtJKRE2Aibl5OYoXjtLkTwpOrqckzcAzm7TIgXlHVmGohXM7Sdh4KI7V; AWSALBCORS=TA9XJyVeQu1gEcXI6qqUpPjYT06y3DU8Ar78haKaPNnr6owTbNWUrZPphyg1g0DzHmwOxtJKRE2Aibl5OYoXjtLkTwpOrqckzcAzm7TIgXlHVmGohXM7Sdh4KI7V; io=oaZfyNtt9nzwRCnSAT1t'
        },
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
