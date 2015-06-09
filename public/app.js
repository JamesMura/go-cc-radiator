'use strict';
// var sock = null;
// window.onload = function() {
//   console.log("onload");
//   sock = new WebSocket(wsuri);
//   sock.onopen = function() {
//     console.log("connected to " + wsuri);
//     sock.send("START")
//   }
//   sock.onclose = function(e) {
//     console.log("connection closed (" + e.code + ")");
//   }
//   sock.onmessage = function(e) {
//     console.log("message received: " + e.data);
//   }

// };

var app = angular.module('radiator', ['ngWebsocket']);
app.service('RadiatorService', [

  function() {
    this.wsUrl = function(path) {
      var loc = window.location,
        wsuri;
      if (loc.protocol === "https:") {
        wsuri = "wss:";
      } else {
        wsuri = "ws:";
      }
      wsuri += "//" + loc.host;
      wsuri += loc.pathname + path;
      return wsuri;
    }
  }
]);
app.controller('MainController', ['RadiatorService', '$websocket', '$scope',
  function(RadiatorService, $websocket, $scope) {
    var ws = $websocket.$new({
      url: RadiatorService.wsUrl('check'),
      reconnect: true
    });
    ws.$on('$open', function() {
      console.log('Socket is open');
      ws.$emit('START');
    })
      .$on('$close', function() {
        console.log('Got close, damn you silly wifi!');
      })
      .$on('UPDATE', function(data) {
        console.log('got an update', data);
        $scope.data = data;
        $scope.$apply();
      })
  }
]);