class Record {
public:
    int playerId;
    int score;

    explicit Record(int a, int b) : playerId(a), score(b) {}

    bool operator<(const Record &rhs) const {
        return score != rhs.score ? score > rhs.score : playerId < rhs.playerId;
    }
};

class Leaderboard {
private:
    set<Record> *s;
    map<int, int> *m;

public:
    Leaderboard() {
        s = new set<Record>();
        m = new map<int, int>();
    }

    void addScore(int playerId, int score) {
        auto p = m->find(playerId);
        if (p == m->end()) {
            m->insert({playerId, score});
            s->emplace(playerId, score);
        } else {
            s->erase(Record(playerId, p->second));
            p->second += score;
            s->emplace(playerId, p->second);
        }
    }

    int top(int K) {
        int k = 0;
        int sum = 0;
        for (auto &ele : *s) {
            sum += ele.score;
            k++;
            if (k == K) {
                break;
            }
        }
        return sum;
    }

    void reset(int playerId) {
        auto p = m->find(playerId);
        if (p != m->end()) {
            s->erase(Record(playerId, p->second));
            p->second = 0;
            s->emplace(playerId, 0);
        }
    }
};

/**
 * Your Leaderboard object will be instantiated and called as such:
 * Leaderboard* obj = new Leaderboard();
 * obj->addScore(playerId,score);
 * int param_2 = obj->top(K);
 * obj->reset(playerId);
 */

