let map = fn(arr, f) {
    let iter = fn(arr, accumulated) {
        if (len(arr) == 0) {
            return accumulated;
        } else {
            return iter(rest(arr), 
              push(accumulated, 
                f(first(arr))
              )
            );
        }
    };
    return iter(arr, []);
};

let double = fn(x) { return x * 2 };
map([1, 2, 3, 4], double);