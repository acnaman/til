const http = require("http")
const fs = require("fs")

http.createServer((req, res) => {
  console.log(req.url)
  if(req.url === "/elm.js") {
    console.log("kotti?");
    const js = fs.readFileSync("./elm.js");
    res.writeHeader(200, {"Content-Type": "text/javascript"});
    res.write(js);
    res.end()
  } else {
    console.log("docchi?");
    const html = fs.readFileSync("./index4_7.html");
    res.writeHeader(200, {"Content-Type": "text/html"});
    res.write(html);
    res.end()
  }
})
.listen(3000)