// package: 
// file: helloWorld.proto

var helloWorld_pb = require("./helloWorld_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var HelloWorld = (function () {
  function HelloWorld() {}
  HelloWorld.serviceName = "HelloWorld";
  return HelloWorld;
}());

HelloWorld.Greeting = {
  methodName: "Greeting",
  service: HelloWorld,
  requestStream: false,
  responseStream: false,
  requestType: helloWorld_pb.Request,
  responseType: helloWorld_pb.Response
};

exports.HelloWorld = HelloWorld;

function HelloWorldClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

HelloWorldClient.prototype.greeting = function greeting(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(HelloWorld.Greeting, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.HelloWorldClient = HelloWorldClient;

