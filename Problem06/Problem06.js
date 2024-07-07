function convert(s, numRows) {
    if (numRows === 1 || numRows >= s.length) {
        return s;
    }
    let result = "";
    // The term is the number of characters in one zigzag cycle
    const term = 2 * numRows - 2;
    for (let i = 0; i < numRows; i++) {
        for (let j = i; j < s.length; j += term) {
            // Always add the current character in the row
            result += s[j];
            // For non-first and non-last rows, add the character that appears diagonally in the zigzag
            if (i !== 0 && i !== numRows - 1) {
                const diagIndex = j + term - 2 * i;
                if (diagIndex < s.length) {
                    result += s[diagIndex];
                }
            }
        }
    }
    return result;
}
const s = "PAYPALISHIRING";
const numRows = 3;
console.log(convert(s, numRows));
