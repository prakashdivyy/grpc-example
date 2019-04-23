const grpc = require("grpc");
const protoLoader = require("@grpc/proto-loader");

//Load the protobuf
const proto = grpc.loadPackageDefinition(
  protoLoader.loadSync("../example.proto", {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
  })
);

//Create a new client instance that binds to the IP and port of the grpc server.
const client = new proto.example.ExampleService(
  "localhost:50050",
  grpc.credentials.createInsecure()
);

client.sayHello({ greeting: "Prakash" }, (error, response) => {
  if (!error) {
    console.log(response.reply);
  } else {
    console.log("Error:", error.message);
  }
});

client.doAddition({ first_number: 1, second_number: 2 }, (error, response) => {
  if (!error) {
    console.log("Addition: " + response.result);
  } else {
    console.log("Error:", error.message);
  }
});

client.doSubtraction(
  { first_number: 1, second_number: 2 },
  (error, response) => {
    if (!error) {
      console.log("Substraction: " + response.result);
    } else {
      console.log("Error:", error.message);
    }
  }
);
