import gleam/string
import gleam/result
import gleam/list
import matrix
import gleam/io
import input

pub fn main() {
  let input = input.line_by_line("src/d4short.txt")
  let width = list.first(input) |> result.unwrap("") |> string.to_graphemes |> list.length
  let height = list.length(input)
  let matrix = matrix.new_matrix(width, height, ".")

  matrix.to_string(matrix, fn(x) { x }) |> io.println
  part1(input) |> construct_matrix(matrix) |>  matrix.to_string(fn(x) { x }) |> io.println

}

fn construct_matrix(input, matrix) {
  list.fold(input, matrix, fn(acc, _line) {
    list.range(0, 1)
      |> list.fold(acc, fn(matrix, x) {
        list.range(0, 1)
        |> list.fold(matrix, fn(matrix, y) { matrix.set(matrix, x, y, "X") })
      })
  })
}

fn part1(input) {
  input
}
