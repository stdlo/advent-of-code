#! /usr/bin/env julia
using Test
include("../utils.jl")

# Make a predicate function that filters based on duplicate characters at specified length
function duplicate_predicate(l)
    return function (n)
        s = string(n)
        b = []
        for c in s
            if length(b) == 0
                b = [c]
            elseif b[1] == c
                b = vcat(c, b)
            elseif length(b) == l
                return true
            else
                b = [c]
            end
        end
        return length(b) == l ? true : false
    end
end

function never_decrease_predicate(n)
    s = string(n)
    last_c = '0'
    for c in s
        if parse(Int, c) < parse(Int, last_c)
            return false
        end
        last_c = c
    end
    return true
end
function solve(range)
    filter!(never_decrease_predicate, range)
    filter!(duplicate_predicate(2), range)
    return length(range)
end

data = collect(254032:789860)
result = solve(data)
@test result == 670
println(result)
