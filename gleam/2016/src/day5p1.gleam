import gleam/bit_array
import gleam/crypto
import gleam/int
import gleam/io
import gleam/iterator
import gleam/string

import input

pub fn main() {
  let in = input.line("src/d5short.txt")
  // iterator.iterate(0, fn(i) {
  //   let ind = in <> int.to_string(i)
  //   crypto.hash(crypto.Md5, <<ind:utf8>>)
  //   |> bit_array.base16_encode
  // })
  // |> iterator.take(1)
  // |> iterator.to_list
  // |> io.debug
  iter(5_017_308, "a", in)
  |> io.debug
}

fn iter(idx, _fun, in) {
  let ind = in <> int.to_string(idx)
  crypto.hash(crypto.Md5, <<ind:utf8>>)
  |> bit_array.base16_encode
}
