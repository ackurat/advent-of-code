import gleam/int
import gleam/io
import gleam/list
import gleam/string

import input

pub fn main() {
  input.line_by_line("src/d3.txt")
  |> list.map(parse_input)
  |> list.transpose
  |> list.map(chunk_list)
  |> list.flatten
  |> list.filter(fn(triangle) {
    case triangle {
      [a, b, c] -> a + b > c && a + c > b && b + c > a
      _ -> False
    }
  })
  |> list.length
  |> io.debug
}

fn chunk_list(l) {
  case l {
    [] -> []
    [a, b, c, ..rest] -> [[a, b, c], ..chunk_list(rest)]
    _ -> []
  }
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
