(dap-register-debug-template
 "Go: Benchmark Fib"
 (list :type "go"
       :request "launch"
       :name "Go: Fib"
       :mode "test"
       :program nil
       :args "-test.bench=BenchmarkFibRecursion25"
       :env nil))
