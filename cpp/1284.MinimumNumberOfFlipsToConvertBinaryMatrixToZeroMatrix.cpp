class Solution {
private:
    int M, N;
    int ans;
    int flag[1024];

public:
    int minFlips(vector<vector<int>> &mat) {
        M = mat.size();
        N = mat[0].size();
        memset(flag, 0, sizeof(flag));
        ans = -1;
        int now = 0;
        for (int i = 1; i <= M; i++) {
            for (int j = 1; j <= N; j++) {
                if (mat[i - 1][j - 1] == 1) {
                    now |= (1 << (M * N - (i - 1) * N - j));
                }
            }
        }
        if (now == 0) {
            return 0;
        }
        dfs(now, 0, 0);

        return ans;
    }

    void dfs(int now, int deep, int last) {
        if (now == 0) {
            if (ans == -1 || deep < ans) {
                ans = deep;
            }
            return;
        }

        int next;
        for (int i = last + 1; i <= M * N; i++) {
            next = mod(now, i);
            if (flag[next] && flag[next] <= flag[now] + 1) continue;
            flag[next] = flag[now] + 1;
            dfs(next, deep + 1, i);
        }
    }

    inline int myxor(int pre, int pos) {
        return pre ^ (1 << (M * N - pos));
    }

    inline int mod(int pre, int pos) {
        pre = myxor(pre, pos);
        //Left
        if (N != 1 && pos % N != 1) pre = myxor(pre, pos - 1);
        //Right
        if (N != 1 && pos % N != 0) pre = myxor(pre, pos + 1);
        //Up
        if (pos > N) pre = myxor(pre, pos - N);
        //Down
        if (pos + N <= M * N) pre = myxor(pre, pos + N);
        return pre;
    }
};

