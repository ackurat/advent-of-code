import gleam/int
import gleam/io
import gleam/list
import gleam/string
import input

pub fn main() {
  let input =
    input.line_by_line("src/d1.txt")
    |> list.map(fn(pairs) {
      string.split(pairs, "   ")
      |> list.map(fn(digit) {
        let assert Ok(int) = int.parse(digit)
        int
      })
    })
    |> list.transpose

  part1(input) |> io.debug
  part2(input) |> io.debug
}

fn part1(input) {
  let sorted = list.map(input, fn(row) { list.sort(row, int.compare) })

  sorted
  |> list.interleave
  |> list.sized_chunk(2)
  |> list.fold(0, fn(acc, pairs) {
    case pairs {
      [a, b] -> acc + int.absolute_value(a - b)
      _ -> acc
    }
  })
}

fn part2(input) {
  case input {
    [first_list, second_list] -> {
      list.map(first_list, fn(elem) {
        elem * list.count(second_list, fn(x) { x == elem })
      })
      |> int.sum
    }
    _ -> 0
  }
}
