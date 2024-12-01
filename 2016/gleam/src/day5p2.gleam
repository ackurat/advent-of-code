import gleam/bit_array
import gleam/crypto
import gleam/int
import gleam/io
import gleam/list
import gleam/option.{None, Some}
import gleam/string

import input

pub fn main() {
  let door_id = input.line("src/d5.txt")

  iter(0, [None, None, None, None, None, None, None, None], door_id)
  |> option.values
  |> string.join("")
  |> string.lowercase
  |> io.debug
}

fn iter(idx, collected, door_id) {
  case option.all(collected) {
    None -> {
      let door_id_idx = door_id <> int.to_string(idx)
      let hash =
        crypto.hash(crypto.Md5, <<door_id_idx:utf8>>)
        |> bit_array.base16_encode

      let res =
        hash
        |> string.starts_with("00000")

      case res {
        True -> {
          let pos = string.slice(hash, 5, 1) |> int.base_parse(16)
          let assert Ok(pos_int) = pos

          let max = list.length(collected)
          let new_collect =
            list.index_map(collected, fn(x, i) {
              case i {
                _j if pos_int == i && pos_int < max ->
                  case x {
                    None -> Some(string.slice(hash, 6, 1))
                    some -> some
                  }
                _ -> x
              }
            })
            |> io.debug
          iter(idx + 1, new_collect, door_id)
        }
        False -> iter(idx + 1, collected, door_id)
      }
    }
    _ -> collected
  }
}
