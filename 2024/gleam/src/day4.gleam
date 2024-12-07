import gleam/dict
import gleam/io
import gleam/list
import gleam/result
import gleam/string
import grids

import input

pub fn main() {
  let input = input.line_by_line("src/d4.txt")
  let grid = parse(input)

  part1(grid) |> io.debug
  part2(grid) |> io.debug
}

fn parse(input) -> grids.Grid(String) {
  use grid, line, line_index <- list.index_fold(input, dict.new())
  use grid, char, col_index <- list.index_fold(string.to_graphemes(line), grid)
  dict.insert(grid, #(line_index, col_index), char)
}

fn part1(grid: grids.Grid(String)) {
  dict.fold(grid, 0, fn(acc, point, _letter) { acc + find_xmas(grid, point) })
}

fn find_xmas(grid: grids.Grid(String), origin: grids.Point) {
  use count, word <- list.fold(
    grids.get_adjacent(origin: origin, distance: 4),
    0,
  )
  case check_adjacent(grid, word) {
    Ok("XMAS") -> count + 1
    _ -> count
  }
}

fn part2(grid: grids.Grid(String)) {
  dict.fold(grid, 0, fn(acc, point, _letter) { acc + find_masmas(grid, point) })
}

fn find_masmas(grid: grids.Grid(String), origin: grids.Point) {
  let diagonal = grids.get_x_pattern(origin, 1)

  case check_adjacent(grid, diagonal.0), check_adjacent(grid, diagonal.1) {
    Ok("MAS"), Ok("MAS") -> 1
    Ok("MAS"), Ok("SAM") -> 1
    Ok("SAM"), Ok("MAS") -> 1
    Ok("SAM"), Ok("SAM") -> 1
    _, _ -> 0
  }
}

fn check_adjacent(grid, points) {
  use string, point <- list.try_fold(points, "")
  use letter <- result.try(dict.get(grid, point))
  Ok(string <> letter)
}
