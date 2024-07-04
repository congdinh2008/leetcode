function twoSum(nums: number[], target: number): number[] {
    const map = new Map<number, number>();
    for (let i = 0; i < nums.length; i++) {
        const complement = target - nums[i];
        if (map.has(complement)) {
            return [map.get(complement) as any, i];
        }
        map.set(nums[i], i);
    }
    return [];
};

const numbers: number[] = [1, 1, 1, 1, 1, 4, 1, 1, 1, 1, 1, 7, 1, 1, 1, 1, 1];
const target: number = 11

const result = twoSum(numbers, target);

console.log(`${result[0]} & ${result[1]}`);
