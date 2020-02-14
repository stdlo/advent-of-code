(require '[clojure.test :as test])

(def cases_p1 [
    [0 "(())"]
    [0 "()()"]
    [3 "((("]
    [3 "(()(()("]
    [3 "))((((("]
    [-1 "())"]
    [-1 "))("]
    [-3 ")))"]
    [-3 ")())())"]
    ])
(def cases_p2 [
    [1 ")"]
    [5 "()())"]
    ])

; Get the value of a parenthesis '(' = 1; ')' = -1
(defn paren_value [a] (if (= a \( ) 1 -1))

; Solve Part 1
(defn solve_p1 [input] (reduce + (map paren_value input)))

; reducer that counts iterations and stops, returning the index, when we hit target
(defn floor_index_reducer [target]
    (fn [acc, val]
        (let [floor (+ val (acc :floor)) index (+ 1(acc :index))]
            ; (println floor index val acc)
            (if (= floor target) ; if floor = target
                (reduced index) ; return index that matches target
                {:floor floor :index index} ; contunue reducing with new acc
            )
        )
    )
)

(defn solve_p2 [input] (reduce (floor_index_reducer -1) {:index 0 :floor 0} (map paren_value input)))

; Define tests from given cases
(test/deftest part1 (doseq [case cases_p1] (test/is (= (nth case 0) (solve_p1 (nth case 1))))))
(test/deftest part2 (doseq [case cases_p2] (test/is (= (nth case 0) (solve_p2 (nth case 1))))))

; Redirect test output to stderr so you can directly copy the results of a successful run
(binding [test/*test-out* *err*] (def summary (test/run-tests)))

; Check test results, if it was good run the main input file and spit the result to stdout
(if (test/successful? summary) (println "part1: " (solve_p1 (slurp "input"))))
(if (test/successful? summary) (println "part2: " (solve_p2 (slurp "input"))))