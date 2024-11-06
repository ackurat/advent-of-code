import gleam/int
import gleam/io
import gleam/list
import gleam/string
import simplifile

pub fn main() {
  let assert Ok(input) = simplifile.read("src/day1.txt")
  let _ = io.debug(part_1(input))
  let _ = io.debug(part_2(input))
}

fn part_1(input) {
  let instructions = parse_instructions(input)
  let final_position = follow_instructions(instructions)
  manhattan_distance(final_position)
}

fn part_2(input) {
  let instructions = parse_instructions(input)
  let final_position = follow_instructions(instructions)
  manhattan_distance(final_position)
}

fn parse_instructions(input) -> List(String) {
  input
  |> string.trim
  |> string.split(", ")
}

fn follow_instructions(instructions) -> #(Int, Int) {
  let start_position = #(0, 0)
  let start_direction = "N"
  list.fold(instructions, #(start_position, start_direction), move)
  |> fn(tup) { tup.0 }
}

fn move(acc, instruction) -> #(#(Int, Int), String) {
  let #(position, direction) = acc
  let #(turn, steps) = parse_instruction(instruction)
  let new_direction = update_direction(direction, turn)
  let new_position = update_position(position, new_direction, steps)
  #(new_position, new_direction)
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
