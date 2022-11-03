import { copy } from "https://deno.land/std@0.159.0/streams/conversion.ts"
const listener = Deno.listen({ port: 3333 })
for await (const conn of listener) {
  copy(conn, conn).finally(() => conn.close())
}
