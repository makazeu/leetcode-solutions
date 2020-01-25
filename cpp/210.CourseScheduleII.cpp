class Node {
public:
    int courseId;
    int prerequisite;

    Node(int courseId, int prerequisite) {
        this->courseId = courseId;
        this->prerequisite = prerequisite;
    }

    bool operator<(const Node &rhs) const {
        if (prerequisite != rhs.prerequisite) {
            return prerequisite < rhs.prerequisite;
        }
        return courseId < rhs.courseId;
    }
};

class Solution {
public:
    vector<int> findOrder(int numCourses, vector<vector<int>> &prerequisites) {
        set<Node> s;
        vector<int> num(numCourses);
        vector<int> out[numCourses];

        for (auto &prerequisite : prerequisites) {
            num[prerequisite[0]]++;
            out[prerequisite[1]].emplace_back(prerequisite[0]);
        }
        for (int i = 0; i < numCourses; i++) {
            s.emplace(i, num[i]);
        }

        vector<int> ans;
        while (ans.size() < numCourses) {
            auto ptr = s.begin();
            if (ptr->prerequisite > 0) {
                return *new vector<int>();
            }
            int course = ptr->courseId;
            s.erase(ptr);

            ans.emplace_back(course);
            for (auto &nxt : out[course]) {
                s.erase(Node(nxt, num[nxt]));
                num[nxt]--;
                s.emplace(nxt, num[nxt]);
            }
        }

        return ans;
    }
};

