class Solution {
public:
    vector<vector<int>> criticalConnections(int n, vector<vector<int>> &connections) {
        vector<int> map[n];
        for (auto &c : connections) {
            map[c[0]].emplace_back(c[1]);
            map[c[1]].emplace_back(c[0]);
        }

        int tot = 0;
        int dfn[n], low[n], pnt[n];
        pnt[0] = 0;
        for (int i = 0; i < n; ++i) dfn[i] = 0;

        function<void(int)> dfs;
        dfs = [&](int u) {
            dfn[u] = ++tot;
            low[u] = dfn[u];
            for (auto &v : map[u]) {
                if (v == pnt[u]) continue;
                if (dfn[v] == 0) {
                    pnt[v] = u;
                    dfs(v);
                    low[u] = min(low[u], low[v]);
                } else {
                    low[u] = min(low[u], dfn[v]);
                }
            }
        };

        dfs(0);
        vector<vector<int>> ans;
        for (auto &c : connections) {
            if (dfn[c[0]] < dfn[c[1]] && low[c[1]] > dfn[c[0]]) {
                ans.emplace_back(vector<int>{c[0], c[1]});
            }
            if (dfn[c[1]] < dfn[c[0]] && low[c[0]] > dfn[c[1]]) {
                ans.emplace_back(vector<int>{c[0], c[1]});
            }
        }

        return ans;
    }
};

