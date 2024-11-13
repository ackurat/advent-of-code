import gleam/bit_array
import gleam/crypto
import gleam/int
import gleam/io
import gleam/list
import gleam/option.{None, Some}
import gleam/string

import input

pub fn main() {
  let door_id = input.line("src/d5short.txt")

  iter(0, [None, None, None, None, None, None, None, None], door_id)
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
          let pos = string.slice(hash, 5, 1) |> int.parse
          let assert Ok(poss) = pos
          let max = list.length(collected)
          list.index_map(collected, fn(x, i) {
            case i {
              _j if poss == i && poss < max -> Some(string.slice(hash, 7, 1))
              _ -> x
            }
          })
          |> io.debug
        }
        False -> collected
      }

      iter(idx + 1, collected, door_id)
    }
    _ -> collected
  }
}
