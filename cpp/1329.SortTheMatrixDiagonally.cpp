class Solution {
public:
    vector<vector<int>> diagonalSort(vector<vector<int>>& mat) {
        int M = mat.size();
        int N = mat[0].size();

        vector<int> tmp;
        for (int j = 1 - M; j < N; j++) {
            tmp.clear();
            for (int i = 0; i < M; i++) {
               if (i + j < 0 || i + j >= N) continue;
               tmp.emplace_back(mat[i][i + j]);
            }

            sort(tmp.begin(), tmp.end());
            int pos = 0;
            for (int i = 0; i < M; i++) {
                if (i + j < 0 || i + j >= N) continue;
                mat[i][i+j] = tmp[pos++];
            }
        }
        return mat;
    }
};

