import gleam/int
import gleam/io
import gleam/list
import gleam/result
import gleam/string

import input

pub fn main() {
  input.line_by_line("src/d4.txt")
  |> list.map(parse_input)
  |> list.map(find_north_pole)
  |> list.filter(fn(str) { string.contains(str, "north") })
  |> io.debug
}

fn parse_input(line) {
  line
  |> string.replace(each: "-", with: " ")
  |> string.drop_right(7)
  |> string.split(" ")
}

fn find_north_pole(room: List(String)) {
  let id = result.try(list.last(room), int.parse) |> result.unwrap(0)

  let words = list.reverse(room) |> list.drop(1) |> list.reverse

  words
  |> list.map(fn(word) {
    word
    |> shift_letters_in_word(id)
  })
  |> string.join(" ")
  |> string.append(" - " <> int.to_string(id))
}

fn shift_letters_in_word(word: String, offset: Int) -> String {
  word
  |> string.to_utf_codepoints
  |> list.fold([], fn(acc, a) {
    let assert Ok(b) =
      a
      |> string.utf_codepoint_to_int
      |> fn(a) {
        let normalized = a - 97
        let shifted = { normalized + offset } % 26
        shifted + 97
      }
      |> string.utf_codepoint

    list.append(acc, [b])
  })
  |> string.from_utf_codepoints
}
