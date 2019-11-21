#! /usr/bin/env julia

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

file_path = try ARGS[1] catch; "input.01.txt" end
data = load_data(file_path)
test_data = [["+1", "-2", "+3", "+1"],[+1, +1, +1],[+1, +1, -2],[-1, -2, -3]]

println("Test data:")
for data in test_data
  println(data, " => ", calc(0,data))
end

println(file_path, " => ", calc(0,data))
