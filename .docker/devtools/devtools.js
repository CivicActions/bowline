var http = require('http');
http.createServer(function (req, res) {
    res.writeHead(200, {'Content-Type': 'text/plain'});
      res.end('Dev Tools with nodejs.\n');
}).listen(80, '0.0.0.0');
console.log('Server running at http://127.0.0.1:80/');
