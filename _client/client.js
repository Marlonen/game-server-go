var reader = new FileReader();
reader.onload = function (e) {
    var js_data = JSON.parse(e.target.result)
    console.log(js_data)
}

var ws = new WebSocket('ws://127.0.0.1:7777')
ws.onopen = function (event) {
    console.log("Send Text WS was opened.");
    ws.send(JSON.stringify({
        C_GameJoin: {
            GameID: 1001,
            RoomID: 1
        }
    }))
};
ws.onmessage = function (event) {
    reader.readAsBinaryString(event.data)
};
ws.onerror = function (event) {
    console.log("Send Text fired an error");
};
ws.onclose = function (event) {
    console.log("WebSocket instance closed.");
};