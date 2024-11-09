import gleam/int
import gleam/io
import gleam/set
import gleam/string
import input

pub fn main() {
  input.split_by_comma("src/day1.txt")
  |> follow_instructions(#(#(0, 0), "N"), set.new() |> set.insert(#(0, 0)))
  |> manhattan_distance
  |> io.debug
}

fn follow_instructions(
  instructions: List(String),
  current: #(#(Int, Int), String),
  visited: set.Set(#(Int, Int)),
) {
  case instructions {
    [] -> #(0, 0)
    [instruction, ..rest] -> {
      let #(current_position, current_heading) = current
      let #(heading, steps) = parse_instruction(instruction)
      let new_heading = update_direction(current_heading, heading)
      let #(new_position, visited, done) =
        walk(steps, current_position, new_heading, visited)
      case done {
        True -> new_position
        _ -> follow_instructions(rest, #(new_position, new_heading), visited)
      }
    }
  }
}

fn walk(steps, location, heading, visited) {
  case steps {
    0 -> #(location, visited, False)
    _ -> {
      let new_position = update_position(location, heading)
      case visited |> set.contains(new_position) {
        True -> #(new_position, visited, True)
        False ->
          walk(
            steps - 1,
            new_position,
            heading,
            visited |> set.insert(new_position),
          )
      }
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

fn update_position(position, direction) -> #(Int, Int) {
  let #(x, y) = position
  case direction {
    "N" -> #(x, y + 1)
    "E" -> #(x + 1, y)
    "S" -> #(x, y - 1)
    "W" -> #(x - 1, y)
    _ -> position
  }
}

fn manhattan_distance(position) -> Int {
  let #(x, y) = position
  int.absolute_value(x) + int.absolute_value(y)
}
