var http = require('http');
var request = require("request");

const options = {
  hostname: 'test1.test',
  port: 80,
  path: '/',
  method: 'GET'
}

//create a server object:
http.createServer(function (req, res) {
    request("http://www.myawesomepage.com/", function (error, response, body) {
        if (!error) {
            res.write(body);
        } else {
            console.log(error);
        }
    });
    res.end(); //end the response
}).listen(8080); //the server object listens on port 8080