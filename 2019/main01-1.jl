#! /usr/bin/env julia
using Test
include("../utils.jl")

# convert string to int or just return an int
to_int = a -> try parse(Int,a) catch; a end

# [test_data, expected_result]
test_cases = [["12", 2],
              [14, 2],
              [1969, 654],
              [100756, 33583]]

fuel_required(mass) = (to_int(mass) / 3 |> floor |> Int) - 2
for case in test_cases
  @test fuel_required(case[1]) == case[2]
end

file_path = try ARGS[1] catch; "inputs/01-1" end
data = read_file_lines(file_path)
println(reduce(+, map(fuel_required, data)))
