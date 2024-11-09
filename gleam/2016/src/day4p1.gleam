import gleam/dict
import gleam/int
import gleam/io
import gleam/list
import gleam/order
import gleam/string

import input

pub fn main() {
  input.line_by_line("src/d4short.txt")
  |> list.map(parse_input)
  |> io.debug
  |> list.filter(check_room)
  |> list.map(fn(room) { room |> list.take(1) })
  |> list.flatten
  |> list.map(fn(room) {
    room
    |> case string.slice(-3, 3) |> int.parse {
      Ok(parsed_id) -> parsed_id
      Error(_) -> 0
    }
  })
  |> io.debug
}

fn parse_input(line) {
  line
  |> string.replace(each: "-", with: "")
  |> string.drop_right(1)
  |> string.split("[")
}

fn check_room(room) {
  case room {
    [a, b] -> {
      let occs =
        a
        |> string.drop_right(3)
        |> string.split("")
        |> list.fold(dict.new(), fn(acc, letter) {
          let new_val = case dict.get(acc, letter) {
            Ok(val) -> val + 1
            Error(_) -> 1
          }
          dict.insert(acc, letter, new_val)
        })
        |> dict.to_list
        |> list.sort(fn(a, b) {
          case int.compare(a.1, b.1) {
            order.Eq -> string.compare(a.0, b.0)
            other -> order.negate(other)
          }
        })
        |> list.take(5)
        |> list.map(fn(tuple) { tuple.0 })
        |> string.concat

      occs == b
    }
    _ -> False
  }
}
