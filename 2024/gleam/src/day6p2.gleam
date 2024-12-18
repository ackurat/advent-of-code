import gleam/result
import gleam/set
import gleam/dict
import gleam/io
import gleam/list
import gleam/string
import grids
import input

pub fn main() {
  let input = input.line_by_line("src/d6.txt")
  let grid = parse(input)
  let guard = find_guard(grid)
  let clean_grid = dict.insert(grid, guard.0, ".")
  let width = list.first(input) |> result.unwrap("")|> string.length
  let height = list.length(input)

  part2(clean_grid, guard, width, height) |> io.debug
}

fn parse(input) -> grids.Grid(String) {
  use grid, line, line_index <- list.index_fold(input, dict.new())
  use grid, char, col_index <- list.index_fold(string.to_graphemes(line), grid)
  dict.insert(grid, #(line_index, col_index), char)
}

fn find_guard(grid) -> grids.Cursor {
  let assert Ok(position) = dict.filter(grid, fn(_pos, char) {
    case char {
      "^" -> True
      _ -> False
    }
  }) |> dict.keys |> list.first
  #(position, grids.Up)
}

fn walk_until_cycle(grid: grids.Grid(String), guard: grids.Cursor, visited: set.Set(#(grids.Point, grids.Direction))) -> Bool {
  let adjacent_coord = grids.get_adjacent_in_direction(origin: guard.0, direction: guard.1)
  let adjacent_char = dict.get(grid, adjacent_coord)
  let previously_visited = set.contains(visited, guard)
  case adjacent_char, previously_visited {
    Ok("."), False ->  walk_until_cycle(grid, #(adjacent_coord, guard.1), set.insert(visited, guard))
    Ok("#"), False ->  walk_until_cycle(grid, #(guard.0, grids.turn(guard.1, grids.Right)), visited)
    _, True -> True
    _, False -> False
  }
}

fn part2(grid, guard, width, height) {
  list.range(0, width * height - 1)
  |> list.filter(fn(idx) {
    let row = idx / width
    let col = idx % width
    case dict.get(grid, #(row, col)) {
      Ok(".") -> {
        let new_grid = dict.insert(grid, #(row, col), "#")
        walk_until_cycle(new_grid, guard, set.new())
      }
      _ -> False
    }
  })
  |> list.length
}
