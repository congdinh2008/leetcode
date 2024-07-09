function longestPalindromeDP02(s: string): string {
    const n = s.length;
    let startIndex = 0;
    let maxlen = 0;
    const table: boolean[][] = Array.from({ length: n }, () => Array(n).fill(false));

    for (let g = 0; g < n; g++) {
        for (let i = 0, j = g; j < n; i++, j++) {
            solve(table, i, j, s);
            if (table[i][j]) {
                if (j - i + 1 > maxlen) {
                    startIndex = i;
                    maxlen = j - i + 1;
                }
            }
        }
    }
    return s.substring(startIndex, startIndex + maxlen);
}

function solve(table: boolean[][], i: number, j: number, s: string): boolean {
    if (i === j) {
        table[i][j] = true;
        return true;
    }
    if (j - i === 1) {
        return table[i][j] = s[i] === s[j];
    }
    if (s[i] === s[j] && table[i + 1][j - 1]) {
        table[i][j] = true;
        return true;
    } else {
        table[i][j] = false;
        return false;
    }
}

function longestPalindromeDP(s: string): string {
    const length = s.length;
    const table: boolean[][] = Array.from({ length }, () => Array(length).fill(false));

    let maxLength = 1;
    let start = 0;

    for (let i = 0; i < length; i++) {
        table[i][i] = true;
    }

    for (let i = 2; i <= length; i++) {
        for (let j = 0; j < length - i + 1; j++) {
            const k = i + j - 1;
            if (table[j + 1][k - 1] && s[j] === s[k]) {
                table[j][k] = true;
                if (i > maxLength) {
                    start = j;
                    maxLength = i;
                }
            }
        }
    }
    return s.substring(start, start + maxLength);
}

function longestPalindrome02(s: string): string {
    let result = "";

    for (let i = 0; i < s.length; i++) {
        const oddPalindrome = expandFromCenter(s, i, i);
        const evenPalindrome = expandFromCenter(s, i, i + 1);

        if (oddPalindrome.length > result.length) {
            result = oddPalindrome;
        }

        if (evenPalindrome.length > result.length) {
            result = evenPalindrome;
        }
    }
    return result;
}

function expandFromCenter(s: string, left: number, right: number): string {
    while (left >= 0 && right < s.length && s[left] === s[right]) {
        left--;
        right++;
    }
    return s.substring(left + 1, right);
}

function longestPalindrome(s: string): string {
    let result = "";

    for (let i = 0; i < s.length; i++) {
        let temp = s.substring(i);
        let tempLen = temp.length;
        while (tempLen > 0) {
            if (isPalindrome(temp)) {
                if (temp.length > result.length) {
                    result = temp;
                }
                break;
            } else {
                tempLen--;
                temp = temp.substring(0, tempLen);
            }
        }
    }
    return result;
}

function isPalindrome(s: string): boolean {
    for (let i = 0; i < s.length / 2; i++) {
        if (s[i] !== s[s.length - 1 - i]) {
            return false;
        }
    }
    return true;
}

console.log(longestPalindromeDP02("abadba"));