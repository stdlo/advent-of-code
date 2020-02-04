#! /usr/bin/env julia
using Test
include("../utils.jl")

function step(vertex, axis, distance, direction)
  vertices=[]
  for i in 1:distance
    vertex[axis] = direction(vertex[axis],1)
    vertices = vcat(vertices,[copy(vertex)])
  end
  return vertices
end

function plot(path)
  current=[0,0]
  coords = [copy(current)]
  for a in path
    num = parse(Int,a[2:length(a)])
    dir = a[1]
    # X Axis
    if dir == 'L'
      coords = vcat(coords, step(current,1,num,-))
    elseif dir == 'R'
      coords = vcat(coords, step(current,1,num,+))
    # Y Axis
    elseif dir == 'U'
      coords = vcat(coords, step(current,2,num,+))
    elseif dir == 'D'
      coords = vcat(coords, step(current,2,num,-))
    end
  end
  return coords
end

function count_steps(path,target)
  for i in 1:length(path)
    if path[i]==target
      return i - 1
    end
  end
end

function transform(wa,wb)
  return function(vertex)
    steps = count_steps(wa,vertex)+count_steps(wb,vertex)
    return (vertex, steps)
  end
end

function solve(wires)
  wa = plot(split(wires[1],","))
  wb = plot(split(wires[2],","))
  intersections = map(transform(wa,wb), intersect(wa,wb)[2:end])

  result=intersections[1]
  for intersection in intersections
    # if result steps > intersection steps
    if result[2] > intersection[2]
      result=intersection
    end
  end
  return result[2]
end

test_cases = [
  [["R8,U5,L5,D3",                                "U7,R6,D4,L4"],                         30],
  [["R75,D30,R83,U83,L12,D49,R71,U7,L72",         "U62,R66,U55,R34,D71,R55,D58,R83"],     610],
  [["R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51","U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"],410],
  ]

for case in test_cases
  @test solve(case[1]) == case[2]
end

file_path = try ARGS[1] catch; "inputs/03-1" end
data = read_file_lines(file_path)
result = solve(data)
@test result == 15622
println(result)