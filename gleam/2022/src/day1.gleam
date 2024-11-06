import gleam/int
import gleam/io
import gleam/list
import gleam/string
import simplifile

pub fn main() {
  let assert Ok(input) = simplifile.read("src/day1input.txt")

  let _ = io.debug(part_1(input))
  io.debug(part_2(input))
}

fn part_1(input) {
  totals(input)
  |> list.sort(by: int.compare)
  |> list.last
}

fn part_2(input) {
  totals(input)
  |> list.sort(by: int.compare)
  |> list.reverse
  |> list.take(3)
  |> list.fold(0, fn(acc, num) { acc + num })
}

fn totals(input) {
  input
  |> string.split(on: "\n\n")
  |> list.map(fn(elf) {
    elf
    |> string.split(on: "\n")
    |> list.fold(0, fn(acc, str) {
      let assert Ok(num) = int.parse(str)
      acc + num
    })
  })
}
