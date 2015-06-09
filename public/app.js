'use strict';
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
        console.log('Got close, damn you silly connection!');
      })
      .$on('UPDATE', function(data) {
        $scope.data = data;
        $scope.$apply();
      })
  }
]);