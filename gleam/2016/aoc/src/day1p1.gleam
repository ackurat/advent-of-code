import gleam/int
import gleam/io
import gleam/string
import input

pub fn main() {
  input.split_by_comma("src/day1.txt")
  |> follow_instructions(#(#(0, 0), "N"))
  |> manhattan_distance
  |> io.debug
}

fn follow_instructions(
  instructions: List(String),
  current: #(#(Int, Int), String),
) {
  case instructions {
    [] -> current.0
    [instruction, ..rest] -> {
      let #(current_position, current_heading) = current
      let #(turn, steps) = parse_instruction(instruction)
      let new_heading = update_direction(current_heading, turn)
      let new_position = update_position(current_position, new_heading, steps)
      follow_instructions(rest, #(new_position, new_heading))
    }
  }
}

fn parse_instruction(instruction) -> #(String, Int) {
  let turn = string.slice(instruction, 0, 1)
  let steps = string.slice(instruction, 1, string.length(instruction))
  case int.parse(steps) {
    Ok(parsed_steps) -> #(turn, parsed_steps)
    Error(_) -> #(turn, 0)
  }
}

fn update_direction(current_direction, turn) -> String {
  case current_direction {
    "N" ->
      case turn {
        "L" -> "W"
        _ -> "E"
      }
    "E" ->
      case turn {
        "L" -> "N"
        _ -> "S"
      }
    "S" ->
      case turn {
        "L" -> "E"
        _ -> "W"
      }
    "W" ->
      case turn {
        "L" -> "S"
        _ -> "N"
      }
    _ -> current_direction
  }
}

fn update_position(position, direction, steps) -> #(Int, Int) {
  let #(x, y) = position
  case direction {
    "N" -> #(x, y + steps)
    "E" -> #(x + steps, y)
    "S" -> #(x, y - steps)
    "W" -> #(x - steps, y)
    _ -> position
  }
}

fn manhattan_distance(position) -> Int {
  let #(x, y) = position
  int.absolute_value(x) + int.absolute_value(y)
}
