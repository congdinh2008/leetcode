# LeetCode - Problem 13 - Roman to Integer

## Description
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9. 
X can be placed before L (50) and C (100) to make 40 and 90. 
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer.

 

Example 1:

Input: s = "III"
Output: 3
Explanation: III = 3.
Example 2:

Input: s = "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.
Example 3:

Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
 

Constraints:

1 <= s.length <= 15
s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
It is guaranteed that s is a valid roman numeral in the range [1, 3999].

## Complexity Analysis
The time complexity for this approach is O(n) where n is the length of the input string. The space complexity is O(1).

## My Approach - RomanToInt
I created a dictionary to store the values of the roman numerals.
```c#
Dictionary<string, int> romans = new Dictionary<string, int>() {
            {"I", 1},
            {"II", 2},
            {"III", 3},
            {"IV", 4},
            {"V", 5},
            {"VI", 6},
            {"VII", 7},
            {"VIII", 8},
            {"IX", 9},
            {"X", 10},
            {"XL", 40},
            {"L", 50},
            {"XC", 90},
            {"C", 100},
            {"CD", 400},
            {"D", 500},
            {"CM", 900},
            {"M", 1000}
        };
```
I will iterate through the string beginning from the end. While iterating, I will check if the current character and the next character form a valid roman numeral. If they do, I will add the value of the roman numeral to the result and skip the next character. If they do not, I will add the value of the current character to the result.

## Internet Approach - RomanToInt02
The internet approach iterates through the string from the beginning. It checks if the current character is less than the next character. If it is, it subtracts the value of the current character from the result. If it is not, it adds the value of the current character to the result.

## Summary
The internet approach is more efficient than my approach. I need to improve my solution to be more efficient.