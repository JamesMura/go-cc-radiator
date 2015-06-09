var sock = null;
var loc = window.location,
  wsuri;
if (loc.protocol === "https:") {
  wsuri = "wss:";
} else {
  wsuri = "ws:";
}
wsuri += "//" + loc.host;
wsuri += loc.pathname + "check"
window.onload = function() {
  console.log("onload");
  sock = new WebSocket(wsuri);
  sock.onopen = function() {
    console.log("connected to " + wsuri);
    sock.send("START")
  }
  sock.onclose = function(e) {
    console.log("connection closed (" + e.code + ")");
  }
  sock.onmessage = function(e) {
    console.log("message received: " + e.data);
  }

};