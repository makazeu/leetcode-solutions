/**
 * // This is the HtmlParser's API interface.
 * // You should not implement it, or speculate about its implementation
 * class HtmlParser {
 *   public:
 *     vector<string> getUrls(string url);
 * };
 */
class Solution {
public:
    vector<string> crawl(string startUrl, HtmlParser htmlParser) {
        string originalHostname = parseUrlHost(startUrl);
        queue<string> q;
        set<string> s;
        q.emplace(startUrl);
        s.emplace(startUrl);

        while (!q.empty()) {
            string now = q.front();
            q.pop();

            auto links = htmlParser.getUrls(now);
            for (auto &link : links) {
                if (parseUrlHost(link) != originalHostname) continue;
                if (s.find(link) != s.end()) continue;
                s.emplace(link);
                q.emplace(link);
            }
        }

        return vector<string>(s.begin(), s.end());
    }

private:
    static string parseUrlHost(string url) {
        url = url.substr(strlen("http://"));
        url = url.substr(0, url.find('/'));
        return url;
    }
};

