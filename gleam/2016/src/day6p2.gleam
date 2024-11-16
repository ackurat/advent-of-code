import gleam/dict
import gleam/int
import gleam/io
import gleam/list
import gleam/order
import gleam/result
import gleam/string

import input

pub fn main() {
  input.line_by_line("src/d6.txt")
  |> list.map(fn(word) { string.to_graphemes(word) })
  |> list.transpose
  |> list.fold("", fn(acc, chars) {
    let char =
      chars
      |> list.fold(dict.new(), fn(acc_dict, char) {
        let new_val = case dict.get(acc_dict, char) {
          Ok(val) -> val + 1
          Error(_) -> 1
        }
        dict.insert(acc_dict, char, new_val)
      })
      |> find_least_common
    string.append(acc, char)
  })
  |> io.debug
}

fn find_least_common(freq: dict.Dict(String, Int)) -> String {
  dict.to_list(freq)
  |> list.sort(fn(a, b) { compare_tuple(a, b) })
  |> list.last
  |> result.map(fn(pair) { pair.0 })
  |> result.unwrap("")
}

fn compare_tuple(a: #(String, Int), b: #(String, Int)) -> order.Order {
  case int.compare(b.1, a.1) {
    order.Eq -> string.compare(a.0, b.0)
    other -> other
  }
}
