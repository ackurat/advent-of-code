import gleam/string
import gleam/int
import gleam/io
import gleam/list
import input

pub fn main() {
  let input =
    input.line_by_line("src/d2.txt")
      |> list.map(fn(digits) {
        string.split(digits, " ")
        |> list.map(fn(digit) {
            let assert Ok(int) = int.parse(digit)
        int
        })
      })

    part1(input) |> io.debug
    part2(input) |> io.debug
}

fn part1(input) {
    input
    |> list.filter(fn(row) {
        let values = 
        list.window_by_2(row)
        
        let decreasing = list.all(values, check_decreasing)
        let increasing = list.all(values, check_increasing)
        case decreasing, increasing {
            True, _ -> True
            _, True -> True
            _, _ -> False
        }
    })
    |> list.length
}

fn check_decreasing(pair: #(Int, Int)) {
    case pair.0, pair.1 {
        a, b if a - b >= 1 && a - b <= 3 -> True
        _ ,_-> False
    }
}

fn check_increasing(pair: #(Int, Int)) {
    case pair.0, pair.1 {
        a, b if b - a >= 1 && b - a <= 3 -> True
        _ ,_-> False
    }
}

fn part2(input) {
    input
    |> list.filter(fn(row) {
        let sublists = list.index_map(row, fn(_char, i) {
            let sublist = list.append(list.take(row, i), list.drop(row, i + 1))
            let sublist_pairs = list.window_by_2(sublist)
            list.all(sublist_pairs, check_decreasing) || list.all(sublist_pairs, check_increasing)
        })
        
        list.any(sublists, fn(res) { res })
    })
    |> list.length
}
