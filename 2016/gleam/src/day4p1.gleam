import gleam/dict
import gleam/int
import gleam/io
import gleam/list
import gleam/order
import gleam/string

import input

pub fn main() {
  input.line_by_line("src/d4.txt")
  |> list.map(parse_input)
  |> list.filter(check_room)
  |> list.map(fn(room) { room |> list.take(1) })
  |> list.flatten
  |> list.fold([], fn(acc, room) {
    let id = string.slice(room, -3, 3)
    case int.parse(id) {
      Ok(parsed_id) -> list.append(acc, [parsed_id])
      Error(_) -> list.append(acc, [0])
    }
  })
  |> list.fold(0, fn(acc, id) { acc + id })
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
