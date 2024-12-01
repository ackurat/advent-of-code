import gleam/bit_array
import gleam/crypto
import gleam/int
import gleam/io
import gleam/string

import input

pub fn main() {
  let door_id = input.line("src/d5.txt")

  iter(0, "", door_id)
  |> io.debug
}

fn iter(idx, collected, door_id) {
  case string.length(collected) {
    8 -> string.lowercase(collected)
    _ -> {
      let door_id_idx = door_id <> int.to_string(idx)
      let hash =
        crypto.hash(crypto.Md5, <<door_id_idx:utf8>>)
        |> bit_array.base16_encode

      let res =
        hash
        |> string.starts_with("00000")

      case res {
        True -> iter(idx + 1, collected <> string.slice(hash, 5, 1), door_id)
        False -> iter(idx + 1, collected, door_id)
      }
    }
  }
}
