import gleam/int
import gleam/io
import gleam/list
import gleam/string

import input

pub fn main() {
  input.line_by_line("src/d3.txt")
  |> list.map(parse_input)
  |> list.filter(fn(triangle) {
    case triangle {
      [a, b, c] -> a + b > c && a + c > b && b + c > a
      _ -> False
    }
  })
  |> list.length
  |> io.debug
}

fn parse_input(line) {
  line
  |> string.trim
  |> string.split(" ")
  |> list.map(fn(entry) {
    case int.parse(entry) {
      Ok(parsed_entry) -> parsed_entry
      Error(_) -> 0
    }
  })
  |> list.filter(fn(x) { x != 0 })
}
