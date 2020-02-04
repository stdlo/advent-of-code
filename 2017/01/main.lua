for line in io.lines("input") do
  if line:len() >= 1 then
    input = f
    list = input .. input
  end
end

function solve(x)
  sum = 0
  i = 1
  for n in input:gmatch"." do
    if n == list:sub(i + x,i + x) then
      sum = sum + n
    end
    i = i + 1
  end
  return sum
end

print("part1 ->",solve(1))
print("part2 ->",solve(input:len() / 2))
