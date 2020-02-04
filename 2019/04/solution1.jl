#! /usr/bin/env julia
using Test
include("../utils.jl")

function double_predicate(n)
  s = string(n)
  last_c = '0'
  for c in s
    if c == last_c
      return true
    end
    last_c=c
  end
  return false
end
function never_decrease_predicate(n)
  s = string(n)
  last_c = '0'
  for c in s
    if parse(Int,c) < parse(Int,last_c)
      return false
    end
    last_c=c
  end
  return true
end
function solve(range)
  filter!(never_decrease_predicate,range)
  filter!(double_predicate,range)
  return length(range)
end

data = collect(254032:789860)
result = solve(data)
@test result == 1033
println(result)