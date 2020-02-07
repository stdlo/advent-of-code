(require '[clojure.test :as test])

(def cases [
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

; Solve function
(defn solve_p1 [input] (reduce + (map #(if (= % \( ) 1 -1) input)))
; reducer that counts iterations and stops when we hit -1 calculations
; (defn reducer [acc, val] (if (= -1 (acc :floor)) (reduced acc) {:floor (+ val (acc :f)) :index (+ 1 (acc :index))}))
(defn reducer [acc, val] {:floor (+ val (acc :f)) :index (+ 1 (acc :index))})
; (reduce (fn [acc, val] {:f (+ val (acc :f)) :i (+ 1(acc :i))}) {:i -1 :f 0} [1 1 1 -1 -1 -1 -1])
    ; (let [sum (+ (nth 1 a) b) i (nth 0 a)]
        ; (println a b)
        ; (if (> sum -1) (reduced sum) [sum])))

(defn solve_p2 [input] (reduce reducer {:index -1 :floor 0} (map #(if (= % \( ) 1 -1) input)))

; Define tests from given cases
; (test/deftest test-cases (doseq [case cases] (test/is (= (nth case 0) (solve_p1 (nth case 1))))))

; Redirect test output to stderr so you can directly copy the results of a successful run
; (binding [test/*test-out* *err*] (def summary (test/run-tests)))

; Check test results, if it was good run the main input file and spit the result to stdout
; (if (test/successful? summary) (println (solve_p2 (slurp "input"))))
(println (solve_p2 (slurp "input")))