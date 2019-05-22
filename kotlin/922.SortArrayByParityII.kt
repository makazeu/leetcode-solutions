import java.util.*

class Solution {
    fun sortArrayByParityII(A: IntArray): IntArray {
        val oddStack = Stack<Int>()
        val evenStack = Stack<Int>()
        for (i in A.indices) {
            if (i % 2 == 0 && A[i] % 2 == 1) {
                evenStack.push(i)
            } else if (i % 2 == 1 && A[i] % 2 == 0) {
                oddStack.push(i)
            }
        }
        
        while (oddStack.isNotEmpty() && evenStack.isNotEmpty()) {
            val x = oddStack.pop()
            val y = evenStack.pop()
            val t = A[x]
            A[x] = A[y]
            A[y] = t
        }
        
        return A
    }
}

