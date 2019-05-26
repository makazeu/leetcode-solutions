let increasingTriplet = function (nums: number[]): boolean {
    if (nums.length < 3) return false;
    let a = nums[0];
    let b = Number.MIN_SAFE_INTEGER;
    for (let num of nums) {
        if (num < a) {
            a = num;
            continue;
        }
        if (num > a) {
            if (b == Number.MIN_SAFE_INTEGER) {
                b = num;
            } else {
                if (num > b) {
                    return true;
                }
                if (num < b) {
                    b = num;
                }
            }
        }
    }
    return false;
};

