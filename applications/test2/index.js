var http = require('http');
var request = require("request");

//create a server object:
http.createServer(function (req, res) {
    request("test1.test", function (error, response, body) {
        if (!error) {
            res.write(body);
        } else {
            console.log(error);
        }
    });
    res.end(); //end the response
}).listen(8080); //the server object listens on port 8080