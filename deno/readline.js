import { TextLineStream } from "https://deno.land/std@0.129.0/streams/mod.ts"

const listener = Deno.listen({ port: 3333 })

for await (const conn of listener) {
  ;(async () => {
    const read = conn.readable
      .pipeThrough(new TextDecoderStream())
      .pipeThrough(new TextLineStream())

    for await (const line of read) {
      conn.write(new TextEncoder().encode(line + "\n"))
    }
  })()
}
