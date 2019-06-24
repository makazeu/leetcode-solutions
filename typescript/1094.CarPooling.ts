let carPooling = function (trips: number[][], capacity: number): boolean {
    let pairs: Array<number[]> = new Array<number[]>();
    trips.forEach(trip => {
        pairs.push([trip[1], trip[0]]);
        pairs.push([trip[2], -trip[0]]);
    });
    pairs.sort((a: number[], b: number[]) => {
        return a[0] != b[0] ? a[0] - b[0] : a[1] - b[1];
    });
    let passenger = 0;
    for (let pair of pairs) {
        passenger += pair[1];
        if (passenger > capacity) {
            return false;
        }
    }
    return true;
};

