import net from "net"

net
  .createServer(function (socket) {
    socket.on("data", function (data) {
      socket.write(data.toString())
    })
  })
  .listen(3333)
