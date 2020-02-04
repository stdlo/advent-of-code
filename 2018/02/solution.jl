#! /usr/bin/env julia
using Test
include("../utils.jl")

pick2and3 = partial(filter, x-> !(x<2 || x>3))
function inc_occr(list, d = Dict()) foreach(a->haskey(d,a) ? d[a] += 1 : d[a] = 1, list); return d end

function solve_v1(arr, d = Dict())
  len = length(arr)
  return  len == 0 ?
  reduce(*, values(d)) :
  solve_v1(arr[2:len],
          inc_occr(pick2and3(unique(values(inc_occr(arr[1])))), d))
end

solve_v2 = list -> (map(pick2and3 ∘ unique ∘ values ∘ inc_occr, list) 
                 |> flatten 
                 |> inc_occr 
                 |> values 
                 |> partial(reduce, *)
                )

test_case = [["abcdef","bababc","abbcde","abcccd","aabcdd","abcdee","ababab"], 12]
@test solve_v1(test_case[1]) == test_case[2]
@test solve_v2(test_case[1]) == test_case[2]

file_path = try ARGS[1] catch; "inputs/02-1" end
data = read_file_lines(file_path)
println(solve_v2(data))

