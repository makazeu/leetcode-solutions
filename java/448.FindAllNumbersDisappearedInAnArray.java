class Solution {
    public List<Integer> findDisappearedNumbers(int[] nums) {
        Set<Integer> set = IntStream.rangeClosed(1, nums.length).boxed()
                .collect(Collectors.toSet());
        set.removeAll(Arrays.stream(nums).boxed().collect(Collectors.toSet()));
        return new ArrayList<>(set);
    }
}

