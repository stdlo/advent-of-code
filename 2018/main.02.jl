#! /usr/bin/env julia
using Test

# return array of data split on newlines from file path
load_data = (file_path) -> open(expanduser(file_path)) do f readlines(f) end
partial = (f, a...) -> (b...) -> f(a..., b...)
function tap(fn, x) fn(x); return x end
logtap = partial(tap, println)

case = [[
"abcdef",# contains no letters that appear exactly two or three times.
"bababc",# contains two a and three b, so it counts for both.
"abbcde",# contains two b, but no letter appears exactly three times.
"abcccd",# contains three c, but no letter appears exactly two times.
"aabcdd",# contains two a and two d, but it only counts once.
"abcdee",# contains two e.
"ababab",# contains three a and three b, but it only counts once.
], 12]


function inc_occr(list, d = Dict()) foreach(a->haskey(d,a) ? d[a] += 1 : d[a] = 1, list); return d end
function recurse(arr, d = Dict())
  len = length(arr)
  return  len == 0 ?
  reduce(*, values(d)) :
  recurse(arr[2:len],
          inc_occr(filter(x-> !(x<2 || x>3), unique(values(inc_occr(arr[1])))), d))
end
my_filter = partial(filter, x-> !(x<2 || x>3))
flatten = a -> collect(Iterators.flatten(a))
solve = list -> (map(my_filter ∘ unique ∘ values ∘ inc_occr, list) 
                 |> flatten 
                 |> inc_occr 
                 |> values 
                 |> partial(reduce, *)
                )
@test recurse(case[1]) == case[2]
@test solve(case[1]) == case[2]

file_path = try ARGS[1] catch; "input.02.txt" end
data = load_data(file_path)

println(solve(data))

