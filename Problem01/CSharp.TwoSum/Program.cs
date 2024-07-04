internal class Program
{
    private static void Main(string[] args)
    {
        int[] nums = [1,1,1,1,1,4,1,1,1,1,1,7,1,1,1,1,1];
        int target = 11;

        var result = TwoSumDict(nums, target);

        System.Console.WriteLine($"{result[0]} & {result[1]}");


    }

    public static int[] TwoSum(int[] nums, int target)
    {
        for (int i = 0; i < nums.Length; i++)
        {
            for (int j = 0; j < nums.Length; j++)
            {
                if (i != j && nums[i] + nums[j] == target)
                {
                    return [i, j];
                }
            }
        }
        return [-1, -1];
    }

    public static int[] TwoSumDict(int[] nums, int target)
    {
        Dictionary<int, int> dict = new Dictionary<int, int>();
        for (int i = 0; i < nums.Length; i++)
        {
            if(dict.ContainsKey(target - nums[i])){
                return [i, dict[target - nums[i]]];
            }
            dict[nums[i]] = i;
        }
        return [];
    }
}