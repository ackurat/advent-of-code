import gleam/int
import gleam/io
import gleam/list
import gleam/result
import gleam/string

import input

pub fn main() {
  input.line_by_line("src/d4short.txt")
  |> list.map(parse_input)
  |> list.filter(find_north_pole)
}

fn parse_input(line) {
  line
  |> string.replace(each: "-", with: " ")
  |> string.drop_right(7)
  |> string.split(" ")
}

fn find_north_pole(room: List(String)) {
  let _id = result.try(list.last(room), int.parse) |> result.unwrap(0)

  let _words = list.reverse(room) |> list.drop(1) |> list.reverse

  shift_letters_in_word("abc", 2)
  True
}

fn shift_letters_in_word(word: String, offset: Int) -> String {
  let shifted_word =
    word
    |> string.to_utf_codepoints
    |> io.debug
    |> list.fold([], fn(acc, a) {
      list.append(acc, [string.utf_codepoint_to_int(a) + { offset % 24 }])
    })
    |> io.debug

  ""
  // |> string.utf_codepoint

  // case shifted_word {
  //   Ok(char) -> [char] |> string.from_utf_codepoints
  //   Error(_) -> "error"
  // }
  // |> io.debug
}
