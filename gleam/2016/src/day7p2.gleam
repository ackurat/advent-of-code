import gleam/io
import gleam/list
import gleam/pair
import gleam/regex
import gleam/string
import input

pub fn main() {
  input.line_by_line("src/d7.txt")
  |> list.filter(supports_ssl)
  |> list.length
  |> io.debug
}

fn supports_ssl(input) {
  let opts = regex.Options(True, False)
  let assert Ok(reg) = regex.compile("\\[(.*?)\\]", opts)

  let bracketed =
    regex.scan(reg, input)
    |> list.map(fn(match) { match.content })
    |> list.map(string.to_graphemes)
    |> list.map(find_aba)

  let outside_brackets =
    regex.split(reg, input)
    |> list.index_map(fn(x, i) { #(i, x) })
    |> list.filter(fn(t) { pair.first(t) % 2 == 0 })
    |> list.map(fn(t) { pair.second(t) })
    |> list.map(string.to_graphemes)
    |> list.map(find_aba)

  case has_aba(outside_brackets) {
    True -> has_matching_bab(outside_brackets, bracketed)
    False -> False
  }
}

fn has_matching_bab(outside, inside) {
  let babs =
    outside
    |> list.flatten
    |> list.map(fn(aba) {
      case string.to_graphemes(aba) {
        [a, b, _] -> b <> a <> b
        _ -> ""
      }
    })

  list.any(babs, fn(bab) { list.contains(list.flatten(inside), bab) })
}

fn has_aba(input) {
  list.any(input, fn(l) { !list.is_empty(l) })
}

fn find_aba(input) {
  input
  |> list.window(3)
  |> list.filter(fn(triplet) {
    case triplet {
      [a, b, c] -> a == c && a != b
      _ -> False
    }
  })
  |> list.map(fn(triplet) {
    case triplet {
      [a, b, c] -> a <> b <> c
      _ -> ""
    }
  })
}
