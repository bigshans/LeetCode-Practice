function min(a, b) {
  return a > b ? b : a;
}
function max(a, b) {
  return a > b ? a : b;
}
function maxArea(height) {
  let maxArea = min(height[0], height[1]);
  if (height.length === 2) {
    return maxArea;
  }
  let i = 0, j = height.length - 1;
  while (i < j) {
    maxArea = max(maxArea, min(height[i], height[j]) * Math.abs(i - j));
    if (height[i] < height[j]) {
      i++;
    } else {
      j--;
    }
  }
  return maxArea;
}

console.log(maxArea([1,8,6,2,5,4,8,3,7]) === 49);
