class TweetCounts {
public:
    TweetCounts() {

    }

    void recordTweet(string tweetName, int time) {
        tweetRecord[tweetName].insert(time);
    }

    vector<int> getTweetCountsPerFrequency(string freq, string tweetName, int startTime, int endTime) {
        int factor = 1;
        if (freq == "minute") {
            factor = 60;
        } else if (freq == "hour") {
            factor = 3600;
        } else if (freq == "day") {
            factor = 86400;
        }
        int num = (endTime - startTime) / factor + 1;

        auto s = tweetRecord[tweetName];
        auto startPtr = s.lower_bound(startTime);
        auto endPtr = s.upper_bound(endTime);

        vector<int> result(num);
        while (startPtr != endPtr) {
            result[(*startPtr - startTime) / factor]++;
            startPtr++;
        }
        return result;
    }

private:
    unordered_map<string, multiset<int>> tweetRecord;
};

