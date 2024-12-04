import gleam/int
import gleam/io
import gleam/list
import gleam/regexp
import gleam/result
import gleam/string
import input

pub fn main() {
  let input = input.line("src/d3.txt")

  part1(input) |> io.debug
  part2(input) |> io.debug
}

fn part1(input) {
  input
  |> extract_muls
}

fn part2(input) {
  let explicit_input = "do()" <> input

  explicit_input
  |> string.split("do()")
  |> list.drop(1)
  |> list.map(fn(do) {
    do
    |> string.split("don't()")
    |> list.first
    |> result.unwrap("")
  })
  |> string.join("")
  |> extract_muls
}

fn extract_muls(input) {
  let opts = regexp.Options(True, False)
  let assert Ok(reg) = regexp.compile("mul\\([0-9]+,[0-9]+\\)", opts)

  regexp.scan(reg, input)
  |> list.map(fn(match) { match.content })
  |> calculate_products
}

fn calculate_products(mul) {
  let opts = regexp.Options(True, False)
  let assert Ok(reg) = regexp.compile("[0-9]+,[0-9]+", opts)

  mul
  |> list.fold(0, fn(acc, mu) {
    let p =
      regexp.scan(reg, mu)
      |> list.map(fn(match) { match.content })
      |> list.map(fn(digits) {
        string.split(digits, ",")
        |> list.map(fn(digit) {
          let assert Ok(parsed) = int.parse(digit)
          parsed
        })
      })
      |> list.flatten
      |> int.product

    acc + p
  })
}
