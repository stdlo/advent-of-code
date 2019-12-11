#! /usr/bin/env julia
using Test
include("../utils.jl")

get_indexes = (p,i) -> [p[i+1]+1,p[i+2]+1,p[i+3]+1]

function apply_opcode(p, i) 
  a,b,o = get_indexes(p,i)
  if p[i] == 1
    p[o] = p[a] + p[b]
  elseif p[i] == 2
    p[o] = p[a] * p[b]
  else 
    throw("opcode not 1 or 2")
  end
  return p
end

function solve(p, i = 1)
 return p[i] == 99 ? p : solve(apply_opcode(p,i), i + 4)
end

gen_permutations=(r,l)->map(a->map(b->(a,b),r),r)
# rand_range = (lo,hi) -> Int(floor(clamp(rand()*100,lo,hi)))

function get_output(data,noun,verb) 
  data[2]=noun
  data[3]=verb
  return solve(data)[1]
end

test_cases = [
  [[1,0,0,0,99], [2,0,0,0,99]],
  [[2,3,0,3,99], [2,3,0,6,99]],
  [[2,4,4,5,99,0], [2,4,4,5,99,9801]],
  [[1,1,1,4,99,5,6,0,99], [30,1,1,4,2,5,6,0,99]],
  ]

for case in test_cases
  @test solve(case[1]) == case[2]
end

file_path = try ARGS[1] catch; "inputs/02-1" end
data = map(a->parse(Int,a), split(read_file_slurp(file_path), ","))

global attempts = 0
for (noun,verb) in flatten(gen_permutations(0:99,5))
  global attempts +=1
  result = try get_output(deepcopy(data),noun,verb) catch; 0 end
  if result == 19690720
    println("noun=$(noun) verb=$(verb) result=$(result) attempts=$(attempts)")
    println(muladd(100,noun,verb))
    break
  end
end
