function lengthOfLongestSubstring(s: string): number {
    const myMap: Map<string, number> = new Map();
    let maxLen = 0;
    let left = 0;

    for (let right = 0; right < s.length; right++) {
        const char = s[right];
        if (myMap.has(char) && myMap.get(char)! >= left) {
            left = myMap.get(char)! + 1;
        }
        myMap.set(char, right);
        const len = right - left + 1;
        if (len > maxLen) {
            maxLen = len;
        }
    }

    return maxLen;
};

const s = "abceeabdcxbb";
console.log(lengthOfLongestSubstring(s)); // 3