class Solution {
    fun twoSum(numbers: IntArray, target: Int): IntArray {
        for (i in 0 until numbers.size) {
            // if (numbers[i] * 2 > target) break
            val j = numbers.binarySearch(target - numbers[i], i + 1)
            if (j < 0 || j >= numbers.size || numbers[j] != target - numbers[i]) continue
            return intArrayOf(i + 1, j + 1)
        }
        return intArrayOf(0, 0)
    }
}

