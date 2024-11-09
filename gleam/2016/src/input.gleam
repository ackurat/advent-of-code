import gleam/string
import simplifile

pub fn open(file) {
  let assert Ok(input) = simplifile.read(file)
  input
}

pub fn line_by_line(file) -> List(String) {
  open(file)
  |> string.trim
  |> string.split("\n")
}

pub fn line(file) -> String {
  open(file)
  |> string.trim
}

pub fn split_by_comma(file) -> List(String) {
  line(file)
  |> string.split(", ")
}
