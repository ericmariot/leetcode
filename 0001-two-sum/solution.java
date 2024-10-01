class Solution {
    public int[] twoSum(int[] nums, int target) {
        HashMap<Integer, Integer> m = new HashMap<>();
        int[] res = new int[2];
        int n = nums.length;

        for (int i = 0; i < n; i++) {
            int t = target - nums[i];
            if (m.containsKey(t)) {
                res[0] = i;
                res[1] = m.get(t);
                return res;
            }
            m.put(nums[i], i);
        }

        return res;
    }
}
