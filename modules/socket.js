var socket = require('socket.io-client')('wss://spin.gamepind.com/stws/?EIO=3&transport=websocket&sid=b626wATjchL2iMLpBNIW');
socket.on('connect', function(){

    console.log('conne')
});
socket.on('event', function(data){


    console.log('event')
});
socket.on('disconnect', function(){


    console.log('disconnect')
});