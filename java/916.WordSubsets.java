class Solution {

    public List<String> wordSubsets(String[] A, String[] B) {
        int[] maxNums = new int[26];
        for (String s : B) {
            int[] nums = new int[26];
            s.chars().forEach(c -> nums[c - 'a']++);
            for (int i = 0; i < nums.length; i++) {
                maxNums[i] = Math.max(maxNums[i], nums[i]);
            }
        }

        return Arrays.stream(A).filter(s -> {
            int[] nums = new int[26];
            s.chars().forEach(c -> nums[c - 'a']++);
            for (int i = 0; i < nums.length; i++) {
                if (nums[i] < maxNums[i]) return false;
            }
            return true;
        }).collect(Collectors.toList());
    }
}

