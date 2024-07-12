
internal class Program
{
    private static void Main(string[] args)
    {
        int[] nums = [-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4];

        var result = ThreeSum(nums);

        foreach (var item in result)
        {
            foreach (var i in item)
            {
                Console.Write(i + " ");
            }
            Console.WriteLine();
        }
    }

    private static IList<IList<int>> ThreeSum(int[] nums)
    {
        IList<IList<int>> result = [];
        Array.Sort(nums);

        for (int i = 0; i < nums.Length - 2; i++)
        {
            if (i > 0 && nums[i] == nums[i - 1]) continue;

            int left = i + 1, right = nums.Length - 1;
            while (left < right)
            {
                int sum = nums[i] + nums[left] + nums[right];

                if (sum == 0)
                {
                    result.Add([nums[i], nums[left], nums[right]]);

                    while (left < right && nums[left] == nums[left + 1]) left++;
                    while (left < right && nums[right] == nums[right - 1]) right--;

                    left++;
                    right--;
                }
                else if (sum < 0)
                {
                    left++;
                }
                else
                {
                    right--;
                }
            }
        }

        return result;
    }

    // Not done
    private static IList<IList<int>> ThreeSumNotDone(int[] nums)
    {
        IList<IList<int>> result = [];
        int len = nums.Length;

        for (int i = 1; i < len - 1; i++)
        {
            int temp = nums[i];

            int left = 0;
            int right = len - 1;

            while (left < i)
            {
                while (right > i)
                {
                    int sum = nums[left] + nums[right];
                    if (sum == 0 - temp)
                    {

                        result.Add([nums[left], nums[i], nums[right]]);
                    }
                    right--;
                }
                right = len - 1;
                left++;
            }
        }

        return result;
    }
}