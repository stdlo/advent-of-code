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
function fuel_required_inclusive(mass, acc=0, depth=0)
  fr = fuel_required(mass)
  addtl = max(fuel_required(fr), 0)
  result = fr + addtl
  acc += result
  println("$(depth)$(join(fill(" ",depth))) $(acc); $(mass); $(fr)")
  if addtl == 0 return acc end
  return fuel_required_inclusive(addtl, acc, depth += 1)
end

for case in test_cases
  @test fuel_required_inclusive(case[1]) == case[2]
end

file_path = try ARGS[1] catch; "inputs/01-1" end
data = read_file_lines(file_path)

#foo = a -> (fuel_required(a) + a)
#println(reduce(+, map( fuel_required_inclusive, data),1))
