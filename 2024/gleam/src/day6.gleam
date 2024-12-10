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

  part1(clean_grid, guard) |> io.debug
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

fn walk(grid: grids.Grid(String), guard: grids.Cursor, visited: set.Set(grids.Point)) -> #(#(grids.Grid(String), grids.Cursor), set.Set(grids.Point)) {
  let adjacent_coord = grids.get_adjacent_in_direction(origin: guard.0, direction: guard.1)
  let adjacent_char = dict.get(grid, adjacent_coord)
  case adjacent_char {
    Ok(".") ->  walk(grid, #(adjacent_coord, guard.1), set.insert(visited, guard.0))
    Ok("#") ->  walk(grid, #(guard.0, grids.turn(guard.1, grids.Right)), visited)
    _ -> #(#(grid, guard), set.insert(visited, guard.0))
  }
}

fn part1(grid, guard) {
  let final_destination = walk(grid, guard, set.new())
  io.debug(final_destination.1)
  #(final_destination.0.1, set.size(final_destination.1))
}