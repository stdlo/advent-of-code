#! /usr/bin/env julia
using Test
include("../utils.jl")

# convert string to int or just return an int
to_int = (a) -> try parse(Int,a) catch; a end

# [test_data, expected_result]
test_cases = [[["+1", "-2", "+3", "+1"], 3],
              [[+1, +1, +1], 3],
              [[+1, +1, -2], 0],
              [[-1, -2, -3], -6]]
for case in test_cases
  @test reduce(+, map(to_int, case[1])) == case[2]
end

file_path = try ARGS[1] catch; "inputs/01-1" end
data = read_file_lines(file_path)
println(reduce(+, map(to_int, data)))
