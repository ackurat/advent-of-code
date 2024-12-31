import gleam/int
import gleam/list
import gleam/string

pub fn parse_list_of_numbers(
  input: List(String),
  separator: String,
) -> List(List(Int)) {
  input
  |> list.map(fn(digits) {
    string.split(digits, separator)
    |> list.map(fn(digit) {
      let assert Ok(int) = int.parse(digit)
      int
    })
  })
}
