let twoCitySchedCost = function (costs: number[][]): number {
    let N = costs.length / 2;
    costs.sort((c1, c2) =>
        Math.abs(c2[0] - c2[1]) - Math.abs(c1[0] - c1[1])
    );

    let cost = 0;
    let num = [0, 0];
    costs.forEach(c => {
        let dest: number;
        if (c[0] < c[1]) {
            dest = num[0] < N ? 0 : 1;
        } else {
            dest = num[1] < N ? 1 : 0;
        }

        num[dest]++;
        cost += c[dest];
    });
    return cost;
};

