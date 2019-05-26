let shortestToChar = function (S: string, C: string): number[] {
    let list: number[] = [];
    for (let i = 0; i < S.length; i++) {
        if (S[i] === C) {
            list.push(i);
        }
    }

    let ans: number[] = new Array(S.length);
    let i = 0, j = 0;
    for (; i < S.length && j < list.length; i++) {
        if (list[j] == i) {
            ans[i] = 0;
            j++;
        } else if (j > 0) {
            ans[i] = Math.min(list[j] - i, i - list[j - 1]);
        } else {
            ans[i] = list[j] - i;
        }
    }
    for (; i < S.length; i++) {
        ans[i] = ans[i - 1] + 1;
    }
    return ans;
};

