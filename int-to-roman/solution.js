const numMap = {
  1: 'I',
  4: 'IV',
  5: 'V',
  9: 'IX',
  10: 'X',
  40: 'XL',
  50: 'L',
  90: 'XC',
  100: 'C',
  400: 'CD',
  500: 'D',
  900: 'CM',
  1000: 'M'
};
function mul(s, times) {
  if (times <= 1) {
    return s;
  }
  return s + mul(s, times - 1);
}
/**
 * @param {number} num
 * @return {string}
 */
var intToRoman = function(num) {
  if (numMap[num]) return numMap[num];
  let base = 10;
  let _num = num;
  let roman = '';
  let t = 1;
  while (_num) {
    const r = _num % base;
    const times = r * t;
    if (r > 5) {
      if (r !== 9) {
        const k = r - 5;
        roman = numMap[5 * t] + mul(numMap[t], k) + roman;
      } else {
        roman = numMap[times] + roman;
      }
    } else if (r === 5) {
      roman = numMap[times] + roman;
    } else if (r > 0) {
      if (r !== 4) {
        roman = mul(numMap[t], r) + roman;
      } else {
        roman = numMap[times] + roman;
      }
    }
    _num = Math.floor(_num / base);
    t *= 10;
  }
  return roman;
};

console.log(intToRoman(1994) === "MCMXCIV");
console.log(intToRoman(58) === "LVIII", intToRoman(58));
console.log(intToRoman(60) === "LX", intToRoman(60));
