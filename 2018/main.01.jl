#! /usr/bin/env julia
using Test

# return array of data split on newlines from file path
function load_data(file_path)
  return open(expanduser(file_path)) do f
    readlines(f)
  end
end

# convert string to int or just return an int
function to_int(a) 
  return try parse(Int,a) catch; a end
end

# recurse over an array and adjust freq
function calc(freq,changes)
  len = length(changes)
  return len == 0 ?
  freq :
  calc(freq + to_int(changes[1]), changes[2:len])
end

# [test_data, expected_result]
test_cases = [[["+1", "-2", "+3", "+1"], 3],
              [[+1, +1, +1], 3],
              [[+1, +1, -2], 0],
              [[-1, -2, -3], -6]]
for case in test_cases
  @test calc(0,case[1]) == case[2]
end

file_path = try ARGS[1] catch; "input.01.txt" end
data = load_data(file_path)

println(calc(0,data))
