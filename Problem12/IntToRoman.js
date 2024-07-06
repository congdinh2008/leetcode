const num = 3999;
console.log(intToRoman(num));
function intToRoman(num) {
    const romanMap = {
        1000: 'M',
        900: 'CM',
        500: 'D',
        400: 'CD',
        100: 'C',
        90: 'XC',
        50: 'L',
        40: 'XL',
        10: 'X',
        9: 'IX',
        5: 'V',
        4: 'IV',
        1: 'I',
    };
    let result = '';
    for (const key of Object.keys(romanMap).reverse()) {
        const value = parseInt(key);
        const roman = romanMap[value];
        while (num >= value) {
            result += roman;
            num -= value;
        }
    }
    return result;
}
