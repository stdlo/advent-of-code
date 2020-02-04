#! /usr/bin/env julia
using Test
include("../utils.jl")

# convert string to int or just return an int
to_int = a -> try parse(Int,a) catch; a end

# [test_data, expected_result]
test_cases = [[14, 2],
              [1969, 966],
              ["100756", 50346]]

fuel_required(mass) = (to_int(mass) / 3 |> floor |> Int) - 2
mathify(acc) = "$(join(acc, " + ")) = $(sum(acc))"
function fuel_required_inclusive(mass, acc=[])
  f = max(fuel_required(mass), 0)
  return f == 0 ?
  reduce(+, acc) : # to log do: (println(mathify(acc)); reduce(+, acc))
  fuel_required_inclusive(f, vcat(acc,f))
end

for case in test_cases
  @test fuel_required_inclusive(case[1]) == case[2]
end

file_path = try ARGS[1] catch; "inputs/01-1" end
data = read_file_lines(file_path)
println(reduce(+, map(fuel_required_inclusive, data))) # == 5180690
