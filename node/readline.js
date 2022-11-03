import net from "net"
import * as readline from "node:readline"

net
  .createServer(async (socket) => {
    let rl = readline.createInterface({ input: socket })
    for await (const line of rl) {
      socket.write(line + "\n")
    }
  })
  .listen(3333)
