(println "こんにちは Clojure!")
(def cases [
"(())" "()()" ; both result in floor 0.
"(((" "(()(()(" ;both result in floor 3.
"))(((((" ; also results in floor 3.
"())" "))(" ; both result in floor -1 (the first basement level).
")))" ")())())" ; both result in floor -3.
])
(defn solve [input] (println (reduce + (map #(if (= % \( ) 1 -1) input))))
(doseq [case cases] (solve case))
; (def input (slurp "input"))