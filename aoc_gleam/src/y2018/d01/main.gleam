import gleam/io
import gleam/string
import gleam/list
import gleam/int
import simplifile

pub fn main() {
  let assert Ok(input) = simplifile.read("src/y2018/d01/input.txt")

  io.debug(part_1(input))
}

fn part_1(input) {
  let freq = totals(input)
}

fn totals(input) {
  input
  |> string.split(on: "\n")
  |> list.fold(0, fn(acc, freq) {
          let assert Ok(num) = int.parse(freq)
          acc + num
    })
}
