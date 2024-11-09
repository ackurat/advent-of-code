import gleam/io
import gleam/list
import gleam/string
import input

const max = 3

const min = 0

pub fn main() {
  input.line_by_line("src/d2.txt")
  |> follow_instructions(#(1, 1), [])
  |> list.reverse
  |> io.debug
}

fn follow_instructions(
  instructions: List(String),
  position: #(Int, Int),
  code: List(Int),
) {
  case instructions {
    [] -> code
    [instruction, ..rest] -> {
      let new_position =
        string.to_graphemes(instruction)
        |> list.fold(position, fn(acc, instruction) {
          update_position(instruction, acc)
          |> io.debug
        })

      follow_instructions(rest, new_position, [
        digit_from_position(new_position),
        ..code
      ])
    }
  }
}

fn update_position(direction, position) -> #(Int, Int) {
  let #(x, y) = position
  case direction {
    "U" if y - 1 >= min -> #(x, y - 1)
    "L" if x - 1 >= min -> #(x - 1, y)
    "D" if y + 1 < max -> #(x, y + 1)
    "R" if x + 1 < max -> #(x + 1, y)
    _ -> position
  }
}

fn digit_from_position(position) -> Int {
  case position {
    #(0, 0) -> 1
    #(1, 0) -> 2
    #(2, 0) -> 3
    #(0, 1) -> 4
    #(1, 1) -> 5
    #(2, 1) -> 6
    #(0, 2) -> 7
    #(1, 2) -> 8
    #(2, 2) -> 9
    _ -> 0
  }
}
