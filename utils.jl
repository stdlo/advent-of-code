export partial, read_file_lines, read_file_slurp, tap, logtap, flatten

# return array of data split on newlines from file path
read_file_lines = (file_path) -> open(expanduser(file_path)) do f readlines(f) end
read_file_slurp = (file_path) -> open(expanduser(file_path)) do f read(f, String) end

partial = (f, a...) -> (b...) -> f(a..., b...)

function tap(fn, x) fn(x); return x end
logtap = partial(tap, println)

flatten = a -> collect(Iterators.flatten(a))
