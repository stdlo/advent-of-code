#! /usr/bin/env julia
using Test

# return array of data split on newlines from file path
load_data = (file_path) -> open(expanduser(file_path)) do f readlines(f) end

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

file_path = try ARGS[1] catch; "input.01.txt" end
data = load_data(file_path)
println(reduce(+, map(to_int, data)))
