const MostRecentTweetNum = 10

type Tweet struct {
	userId    int
	tweetId   int
	timestamp int
}

type User struct {
	id        int
	tweets    []*Tweet
	following map[*User]bool
}

type Twitter struct {
	timestamp int
	users     map[int]*User
}

type NewsFeed []*Tweet

func (newsFeed NewsFeed) Len() int {
	return len(newsFeed)
}

func (newsFeed NewsFeed) Less(i, j int) bool {
	return newsFeed[i].timestamp < newsFeed[j].timestamp
}

func (newsFeed NewsFeed) Swap(i, j int) {
	newsFeed[i], newsFeed[j] = newsFeed[j], newsFeed[i]
}

func (newsFeed *NewsFeed) Push(x interface{}) {
	*newsFeed = append(*newsFeed, x.(*Tweet))
}

func (newsFeed *NewsFeed) Pop() interface{} {
	n := len(*newsFeed)
	item := (*newsFeed)[n-1]
	(*newsFeed)[n-1] = nil
	*newsFeed = (*newsFeed)[0 : n-1]
	return item
}

func (this *Twitter) getUser(userId int) *User {
	if user := this.users[userId]; user != nil {
		return user
	}

	var user User
	user.id = userId
	user.following = make(map[*User]bool)
    user.following[&user] = true
	this.users[userId] = &user
	return &user
}

func (user *User) addTweet(tweet *Tweet) {
	user.tweets = append(user.tweets, tweet)
}

func (user *User) follow(followee *User) {
	user.following[followee] = true
}

func (user *User) unfollow(followee *User) {
	delete(user.following, followee)
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	var twitter Twitter
	twitter.users = make(map[int]*User, 0)
	return twitter
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	tweet := &Tweet{userId, tweetId, this.timestamp}
	this.timestamp++

	this.getUser(userId).addTweet(tweet)
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	var user *User
	if user = this.users[userId]; user == nil {
		return []int{}
	}

	var newsFeed NewsFeed
	heap.Init(&newsFeed)
	for followee := range user.following {
		for i := len(followee.tweets) - 1; i >= 0 && i >= len(followee.tweets)-MostRecentTweetNum; i-- {
			currentTweet := followee.tweets[i]
			if newsFeed.Len() < MostRecentTweetNum {
				heap.Push(&newsFeed, currentTweet)
				continue
			}
			leastRecentTweet := newsFeed[0]
			if currentTweet.timestamp > leastRecentTweet.timestamp {
				heap.Pop(&newsFeed)
				heap.Push(&newsFeed, currentTweet)
			}
		}
	}

	n := newsFeed.Len()
	result := make([]int, n)
	for newsFeed.Len() > 0 {
        n--
        result[n] = heap.Pop(&newsFeed).(*Tweet).tweetId
	}
	return result
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	follower := this.getUser(followerId)
	followee := this.getUser(followeeId)
	follower.follow(followee)
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
    if followerId == followeeId {
        return
    }
	if this.users[followerId] == nil || this.users[followeeId] == nil {
		return
	}
	follower := this.getUser(followerId)
	followee := this.getUser(followeeId)
	follower.unfollow(followee)
}

