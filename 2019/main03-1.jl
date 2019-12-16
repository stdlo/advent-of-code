#! /usr/bin/env julia
using Test
include("../utils.jl")
function plot(path)
  path = map( a -> split(a,""), path)
  current=[0,0]
  coords = [copy(current)]
  for a in path
    # X Axis
    if a[1] == "R"
      current[1] += parse(Int,a[2])
    elseif a[1] == "L"
      current[1] -= parse(Int,a[2])
    # Y Axis
    elseif a[1] == "U"
      current[2] += parse(Int,a[2])
    elseif a[1] == "D"
      current[2] -= parse(Int,a[2])
    end
    coords=vcat(coords,[copy(current)])
  end
  return coords
end
# Take current and last coords and determine if an intersection occurred
function determine_intersection(current_a, current_b, prev_a, prev_b)
    # if last coord flip flopped high/low then an intersection occured at the new high
    # println("---")
    # println("current_a,b  : ", current_a, current_b)
    # println("prev_a,b     : ", prev_a, prev_b)


    # println("current_b  : ", current_b)
    # println("prev_b     : ", prev_b)
    current_lesser_a = current_a[1] >= current_a[2]  ? 2 : 1
    current_lesser_b = current_b[1] >= current_b[2]  ? 2 : 1
    prev_lesser_a = prev_a[1] >= prev_a[2]  ? 2 : 1
    prev_lesser_b = prev_b[1] >= prev_b[2]  ? 2 : 1
    # println("a => ", current_greater_a)
    # println("b => ", current_greater_b)
    println(current_a,current_b, current_lesser_a == prev_lesser_a, current_lesser_b == prev_lesser_b)
    if current_lesser_a == prev_lesser_a && current_lesser_b == prev_lesser_b
      println("INTERSECT ", current_lesser_a => current_a[current_lesser_a], current_lesser_b => current_b[current_lesser_b])
    end
    println("---")


end

function solve(wires)
  wa = plot(split(wires[1],","))
  wb = plot(split(wires[2],","))
  larger_length = max(length(wa),length(wb))
  for i in 1:larger_length
    println(wa[i],wb[i])
    if wa[i] == [0,0]; continue end
    determine_intersection(wa[i], wb[i], wa[i-1], wb[i-1])
  end
end
test_cases = [
  [["R8,U5,L5,D3", "U7,R6,D4,L4"], 6],
  ]

for case in test_cases
  @test solve(case[1]) == case[2]
end

# file_path = try ARGS[1] catch; "inputs/03-1" end
# data = map(a->parse(Int,a), split(read_file_slurp(file_path), ","))
