/**
 * @param {number} N
 * @param {number[][]} conections
 * @return {number}
 */
const minimumCost = function (N: number, conections: number[][]): number {
    if (N == 1) {
        return 0;
    }
    conections.sort((a, b) => {
        return a[2] - b[2];
    });
    let ufs: number[] = new Array(N + 1);
    for (let i = 1; i <= N; i++) {
        ufs[i] = i;
    }

    const find = (k: number): number => {
        if (ufs[k] == k) return k;
        ufs[k] = find(ufs[k]);
        return ufs[k];
    };

    const merge = (a: number, b: number) => {
        ufs[b] = a;
    };

    let n = 0;
    let ans = 0;
    for (let c of conections) {
        let a = find(c[0]),
            b = find(c[1]);
        if (a != b) {
            merge(a, b);
            n++;
            ans += c[2];
        }
        if (n == N - 1) {
            break;
        }
    }

    return n == N - 1 ? ans : -1;
};

