package Problem01;

import java.util.HashMap;

/**
 * TwoSum
 */
public class TwoSum {

    public static void main(String[] args) {
        int[] nums = { 1, 1, 1, 1, 1, 4, 1, 1, 1, 1, 1, 7, 1, 1, 1, 1, 1 };
        int target = 11;

        int[] result = twoSumDict(nums, target);
        System.out.printf("%d & %d\n", result[0], result[1]);
    }

    public static int[] twoSum(int[] nums, int target) {
        for (int i = 0; i < nums.length; i++) {
            for (int j = 0; j < nums.length; j++) {
                if (i != j && nums[i] + nums[j] == target) {
                    return new int[] { i, j };
                }
            }
        }
        return null;
    }

    private static int[] twoSumDict(int[] nums, int target) {
        HashMap<Integer, Integer> map = new HashMap<Integer, Integer>();

        for (int i = 0; i < nums.length; i++) {
            int temp = target - nums[i];
            if (map.containsKey(temp)) {
                return new int[] { map.get(temp), i };
            }
            map.put(nums[i], i);
        }
        return null;
    }
}