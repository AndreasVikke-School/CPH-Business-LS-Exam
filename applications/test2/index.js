var http = require('http');

const options = {
  hostname: 'test1',
  port: 80,
  path: '/',
  method: 'GET'
}

//create a server object:
http.createServer(function (req, res) {
    const reqq = http.request(options, ress => {
        console.log(`statusCode: ${ress.statusCode}`)

        ress.on('data', d => {
            res.write(d);
        })
    })
    res.end(); //end the response
}).listen(8080); //the server object listens on port 8080