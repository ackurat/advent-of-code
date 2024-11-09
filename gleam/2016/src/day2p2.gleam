import gleam/io
import gleam/list
import gleam/string
import input

pub fn main() {
  input.line_by_line("src/d2.txt")
  |> follow_instructions(#(0, 2), "")
  |> io.debug
}

fn follow_instructions(
  instructions: List(String),
  position: #(Int, Int),
  code: String,
) {
  case instructions {
    [] -> code
    [instruction, ..rest] -> {
      let new_position =
        string.to_graphemes(instruction)
        |> list.fold(position, fn(acc, instruction) {
          update_position(instruction, acc)
        })

      follow_instructions(
        rest,
        new_position,
        string.append(code, digit_from_position(new_position)),
      )
    }
  }
}

// 0 0 1 0 0
// 0 2 3 4 0
// 5 6 7 8 9
// 0 A B C 0
// 0 0 D 0 0
fn update_position(direction, position) -> #(Int, Int) {
  let #(x, y) = position
  let min = 0
  let max = 4

  let new_position = case direction {
    "U" if y - 1 >= min -> #(x, y - 1)
    "L" if x - 1 >= min -> #(x - 1, y)
    "D" if y + 1 <= max -> #(x, y + 1)
    "R" if x + 1 <= max -> #(x + 1, y)
    _ -> position
  }

  case new_position {
    #(2, 0) -> new_position
    #(1, 1) -> new_position
    #(2, 1) -> new_position
    #(3, 1) -> new_position
    #(0, 2) -> new_position
    #(1, 2) -> new_position
    #(2, 2) -> new_position
    #(3, 2) -> new_position
    #(4, 2) -> new_position
    #(1, 3) -> new_position
    #(2, 3) -> new_position
    #(3, 3) -> new_position
    #(2, 4) -> new_position
    _ -> position
  }
}

fn digit_from_position(position) -> String {
  case position {
    #(2, 0) -> "1"
    #(1, 1) -> "2"
    #(2, 1) -> "3"
    #(3, 1) -> "4"
    #(0, 2) -> "5"
    #(1, 2) -> "6"
    #(2, 2) -> "7"
    #(3, 2) -> "8"
    #(4, 2) -> "9"
    #(1, 3) -> "A"
    #(2, 3) -> "B"
    #(3, 3) -> "C"
    #(2, 4) -> "D"
    _ -> ""
  }
}
