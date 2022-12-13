/**
 * @param {number[]} nums
 * @return {number[]}
 */
var buildArray = function(nums) {
  return nums.map(i => nums[i]);
};

console.log(buildArray([0,2,1,5,3,4]));
