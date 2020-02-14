using System;
namespace AdventOfCode {
    class Solution
    {
        static void Main(string[] args)
        {
            char[] input = System.IO.File.ReadAllText(@"input").ToCharArray();
            Console.WriteLine("part1: {0}", SolveP1(input));
            Console.WriteLine("part2: {0}", SolveP2(input));
        }

        static int ParenValue(char c)
        {
            if (c == '(') {
                return 1;
            }
            return -1;
        }

        static int SolveP1(char[] input)
        {
            int result = 0;
            for (var i = 0; i < input.Length; i++) {
                result += ParenValue(input[i]);
            }
            return result;
        }

        static int SolveP2(char[] input)
        {
            int result = 0;
            for (var i = 0; i < input.Length; i++) {
                result += ParenValue(input[i]);
                if (result == -1) {
                    // c# indexed at 0, but the problem wants us to index at 1.
                    return i + 1; 
                }
            }
            return 0;
        }
    }
}