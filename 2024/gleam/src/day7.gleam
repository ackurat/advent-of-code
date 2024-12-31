import gleam/int
import gleam/io
import gleam/list
import gleam/string

import input

pub type Equation {
  Equation(result: Int, operands: List(Int))
}

pub type Operators {
  Add
  Multiply
  Concat
}

fn apply_operator(op: Operators, a, b) {
  case op {
    Add -> a + b
    Multiply -> a * b
    Concat -> {
      let assert Ok(conced) = int.parse(int.to_string(a) <> int.to_string(b))
      conced
    }
  }
}

pub fn main() {
  let input = input.line_by_line("src/d7.txt") |> parse

  part1(input) |> io.debug
  part2(input) |> io.debug
}

fn parse(input) {
  use line <- list.map(input)

  let assert [result, operands] = string.split(line, ": ")
  let assert Ok(result) = int.parse(result)
  let operands = string.split(operands, " ") |> list.filter_map(int.parse)
  Equation(result: result, operands: operands)
}

fn part1(input: List(Equation)) {
  input
  |> list.filter(fn(equation) { calibrate(equation, [Add, Multiply]) })
  |> list.fold(0, fn(acc, equation) { acc + equation.result })
}

fn part2(input: List(Equation)) {
  input
  |> list.filter(fn(equation) { calibrate(equation, [Add, Multiply, Concat]) })
  |> list.fold(0, fn(acc, equation) { acc + equation.result })
}

fn calibrate(equation: Equation, operators: List(Operators)) -> Bool {
  case equation.operands {
    [] -> False
    [head, ..tail] ->
      do_calibrate(Equation(equation.result, tail), head, operators)
  }
}

fn do_calibrate(
  equation: Equation,
  acc: Int,
  operators: List(Operators),
) -> Bool {
  case equation.operands {
    [] -> acc == equation.result
    _ if acc > equation.result -> False
    [head, ..tail] -> {
      list.any(operators, fn(operator) {
        do_calibrate(
          Equation(equation.result, tail),
          apply_operator(operator, acc, head),
          operators,
        )
      })
    }
  }
}
