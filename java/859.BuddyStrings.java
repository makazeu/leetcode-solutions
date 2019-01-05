class Solution {
    public boolean buddyStrings(String A, String B) {
        if (A.length() != B.length()) {
            return false;
        }

        if (A.equals(B)) {
            return IntStream.range(0, A.length())
                    .mapToObj(A::charAt)
                    .collect(Collectors.groupingBy(Function.identity(), Collectors.counting()))
                    .values()
                    .stream()
                    .anyMatch(a -> a > 1);
        }

        List<Integer> list = new ArrayList<>();
        for (int i = 0; i < A.length(); i++) {
            if (A.charAt(i) != B.charAt(i)) {
                list.add(i);
            }
        }

        if (list.size() != 2) {
            return false;
        }

        return A.charAt(list.get(0)) == B.charAt(list.get(1))
                && A.charAt(list.get(1)) == B.charAt(list.get(0));
    }
}
