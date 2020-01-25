const int BEGIN = 1;
const int END = 2;

class Event {
public:
    int event;
    int position;

    Event(int event, int position) {
        this->event = event;
        this->position = position;
    }
};

bool cmp(Event a, Event b) {
    return a.position != b.position ? a.position < b.position : a.event < b.event;
}

class Solution {
public:
    vector<vector<int>> merge(vector<vector<int>> &intervals) {
        vector<Event> events;
        for (auto &interval : intervals) {
            events.emplace_back(BEGIN, interval[0]);
            events.emplace_back(END, interval[1]);
        }
        sort(events.begin(), events.end(), cmp);

        int cnt = 0;
        int start;
        vector<vector<int>> ans;
        for (auto &e : events) {
            if (e.event == BEGIN) {
                if (cnt == 0) {
                    start = e.position;
                }
                cnt++;
            } else {
                cnt--;
                if (cnt == 0) {
                    ans.push_back({start, e.position});
                }
            }
        }
        return ans;
    }
};

