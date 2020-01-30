const int MOST_RECENT_TWEET_NUM = 10;

class Tweet {
public:
    int tweetId;
    int timestamp;

    Tweet(int tweetId, int timestamp) : tweetId(tweetId), timestamp(timestamp) {}
};

bool tweet_cmp(Tweet *a, Tweet *b) {
    return a->timestamp < b->timestamp;
}

class User {
public:
    int id;
    vector<Tweet *> tweets;
    set<User *> following;

    User(int id) : id(id) {}
};

class Twitter {
public:
    /** Initialize your data structure here. */
    Twitter() {
    }

    /** Compose a new tweet. */
    void postTweet(int userId, int tweetId) {
        auto user = getUser(userId);
        user->tweets.emplace_back(new Tweet(tweetId, timestamp++));
    }

    /** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
    vector<int> getNewsFeed(int userId) {
        set<Tweet *, decltype(tweet_cmp) *> newsFeed(tweet_cmp);
        auto user = getUser(userId);

        for (auto followee:user->following) {
            int tweetNum = followee->tweets.size();
            for (int i = tweetNum - 1; i >= 0 && i >= tweetNum - MOST_RECENT_TWEET_NUM; i--) {
                auto currentTweet = followee->tweets[i];
                if (newsFeed.size() < MOST_RECENT_TWEET_NUM) {
                    newsFeed.insert(currentTweet);
                    continue;
                }

                auto firstTweet = *newsFeed.begin();
                if (firstTweet->timestamp < currentTweet->timestamp) {
                    newsFeed.erase(newsFeed.begin());
                    newsFeed.insert(currentTweet);
                }
            }
        }

        int num = newsFeed.size();
        vector<int> ans(num);
        while (!newsFeed.empty()) {
            ans[--num] = (*newsFeed.begin())->tweetId;
            newsFeed.erase(newsFeed.begin());
        }
        return ans;
    }

    /** Follower follows a followee. If the operation is invalid, it should be a no-op. */
    void follow(int followerId, int followeeId) {
        auto follower = getUser(followerId);
        auto followee = getUser(followeeId);
        follower->following.insert(followee);
    }

    /** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
    void unfollow(int followerId, int followeeId) {
        if (followerId == followeeId) return;
        if (!checkUserExists(followerId) || !checkUserExists(followeeId)) return;
        auto follower = getUser(followerId);
        auto followee = getUser(followeeId);
        follower->following.erase(followee);
    }

private:
    int timestamp = 0;
    map<int, User *> users;

    User *getUser(int userId) {
        auto ptr = users.find(userId);
        if (ptr == users.end()) {
            auto user = new User(userId);
            user->following.insert(user);
            users.insert({userId, user});
            return user;
        }
        return ptr->second;
    }

    bool checkUserExists(int userId) {
        return users.find(userId) != users.end();
    }
};

