class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        List<List<Integer>> answers = new ArrayList<>();
        Arrays.sort(nums);

        // build a map
        Map<Integer, Integer> map = new HashMap<>();
        for (int num : nums) {
            map.put(num, map.getOrDefault(num, 0) + 1);
        }

        for (int i = 0; i < nums.length - 2; i++) {
            if (i > 0 && nums[i] == nums[i - 1]) {
                continue;
            }
            map.put(nums[i], map.get(nums[i]) - 1);
            for (int j = i + 1; j < nums.length - 1; j++) {
                if (j > i + 1 && nums[j] == nums[j - 1]) {
                    continue;
                }
                map.put(nums[j], map.get(nums[j]) - 1);
                int ret = 0 - nums[i] - nums[j];
                if (ret >= nums[j] && map.getOrDefault(ret, 0) > 0) {
                    answers.add(Arrays.asList(nums[i], nums[j], ret));
                }
                map.put(nums[j], map.get(nums[j]) + 1);
            }
            map.put(nums[i], map.get(nums[i]) + 1);
        }

        return answers;
    }
}
