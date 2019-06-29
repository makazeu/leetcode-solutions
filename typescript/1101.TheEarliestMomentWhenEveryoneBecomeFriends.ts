const earliestAcq = function (logs: number[][], N: number): number {
    logs.sort((a, b) => a[0] - b[0]);
    let ufs: number[] = new Array(N);
    const find = (k: number) => {
        return ufs[k] == k ? k : ufs[k] = find(ufs[k]);
    };
    const union = (x: number, y: number) => {
        let a = find(x);
        let b = find(y);
        if (a != b) {
            ufs[b] = a;
        }
    };
    for (let i = 0; i < N; i++) {
        ufs[i] = i;
    }
    let left = N;
    let ans = logs[0][0];
    for (let log of logs) {
        let x = find(ufs[log[1]]);
        let y = find(ufs[log[2]]);
        if (x == y) continue;
        union(x, y);
        ans = log[0];
        if (left-- <= 1) break;
    }
    return left <= 1 ? ans : -1;
};

