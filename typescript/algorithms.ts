// iterative binary search
const BinarySearchIterative = (arr: number[], whatToSearchFor: number) => {
  let left = 0;
  let right = arr.length - 1;
  while (left <= right) {
    let middle = Math.floor((right + left) / 2);
    if (arr[middle] === whatToSearchFor) return middle;
    if (arr[middle] > whatToSearchFor) {
      right = middle - 1;
      continue;
    }
    left = middle + 1;
  }
  return -1;
};

// console.log(BinarySearchIterative([1, 2, 3, 4, 5, 6, 7, 8, 9], 3)); //=> 2
// console.log(BinarySearchIterative([1, 2, 3, 4, 5, 6, 7, 8, 9], 10)); //=> -1

// recursive binary search
const BinarySearchRecursive = (
  arr: number[],
  whatToSearchFor: number,
  left: number,
  right: number
): number => {
  if (left > right) return -1;

  const middle = Math.floor((left + right) / 2);

  if (arr[middle] === whatToSearchFor) {
    return middle;
  }

  if (arr[middle] > whatToSearchFor) {
    return BinarySearchRecursive(arr, whatToSearchFor, left, middle - 1);
  }

  return BinarySearchRecursive(arr, whatToSearchFor, middle + 1, right);
};

// console.log(BinarySearchRecursive([1, 2, 3, 4, 5, 6, 7, 8, 9], 3, 0, 8)); //=> 2
// console.log(BinarySearchRecursive([1, 2, 3, 4, 5, 6, 7, 8, 9], 10, 0, 8)); //=> -1

// Meta binary search (one sided binary search)
const BinarySearchMeta = (arr: number[], whatToSearchFor: number) => {
  let bitsToHoldArrLength = 0;
  while (1 << bitsToHoldArrLength < arr.length - 1) {
    bitsToHoldArrLength++;
  }
  let pos = 0;

  for (let i = bitsToHoldArrLength; i >= 0; i--) {
    if (arr[pos] === whatToSearchFor) return pos;
    const newPos = pos | (1 << i);

    if (newPos < arr.length && arr[newPos] <= whatToSearchFor) {
      pos = newPos;
    }
  }
  return arr[pos] === whatToSearchFor ? pos : -1;
};

// console.log(BinarySearchMeta([1, 2, 3, 4, 5, 6, 7, 8, 9], 3)); //=> 2
// console.log(BinarySearchMeta([1, 2, 3, 4, 5, 6, 7, 8, 9], 10)); //=> -1

// ternary search iterative

const TernarySearchIterative = (arr: number[], whatToSearchFor: number) => {
  let left = 0;
  let right = arr.length - 1;

  while (left <= right) {
    const mid1 = left + Math.floor((right - left) / 3);
    const mid2 = right - Math.floor((right - left) / 3);

    if (arr[mid1] === whatToSearchFor) {
      return mid1;
    }
    if (arr[mid2] === whatToSearchFor) {
      return mid2;
    }
    if (arr[mid1] > whatToSearchFor) {
      right = mid1 - 1;
      continue;
    }
    if (arr[mid2] < whatToSearchFor) {
      left = mid2 + 1;
      continue;
    }
    left = mid1 + 1;
    right = mid2 - 1;
  }

  return -1;
};

// console.log(TernarySearchIterative([1, 2, 3, 4, 5, 6, 7, 8, 9], 3)); //=> 2
// console.log(TernarySearchIterative([1, 2, 3, 4, 5, 6, 7, 8, 9], 10)); //=> -1

// ternary search recursive

const TernarySearchRecursive = (
  arr: number[],
  whatToSearchFor: number,
  left: number,
  right: number
): number => {
  if (left > right) {
    return -1;
  }

  const mid1 = left + Math.floor((right - left) / 3);
  const mid2 = right - Math.floor((right - left) / 3);

  if (arr[mid1] === whatToSearchFor) {
    return mid1;
  }

  if (arr[mid2] === whatToSearchFor) {
    return mid2;
  }

  if (arr[mid1] > whatToSearchFor) {
    return TernarySearchRecursive(arr, whatToSearchFor, left, mid1 - 1);
  }

  if (arr[mid2] < whatToSearchFor) {
    return TernarySearchRecursive(arr, whatToSearchFor, mid2 + 1, right);
  }

  return TernarySearchRecursive(arr, whatToSearchFor, mid1 + 1, mid2 - 1);
};

// console.log(TernarySearchRecursive([1, 2, 3, 4, 5, 6, 7, 8, 9], 3, 0, 8)); //=> 2
// console.log(TernarySearchRecursive([1, 2, 3, 4, 5, 6, 7, 8, 9], 10, 0, 8)); //=> -1

// Jump search

const JumpSearch = (arr: number[], whatToSearchFor: number) => {
  const length = arr.length;
  const step = Math.floor(Math.sqrt(length));

  let start = 0;
  let end = step;

  while (arr[Math.min(length, end) - 1] < whatToSearchFor) {
    [start, end] = [end, end + step];
    if (start >= length) {
      return -1;
    }
  }

  while (arr[start] < whatToSearchFor) {
    start++;
    if (start === Math.min(length, end)) {
      return -1;
    }
  }

  return arr[start] === whatToSearchFor ? start : -1;
};

// console.log(JumpSearch([1, 2, 3, 4, 5, 6, 7, 8, 9], 3)); //=> 2
// console.log(JumpSearch([1, 2, 3, 4, 5, 6, 7, 8, 9], 10)); //=> -1

// ExponentialSearch
const ExponentialSearch = (arr: number[], whatToSearchFor: number) => {
  if (arr[0] === whatToSearchFor) return 0;
  let i = 1;
  while (i < arr.length && arr[i] <= whatToSearchFor) {
    i = i * 2;
  }
  return BinarySearchRecursive(
    arr,
    whatToSearchFor,
    i / 2,
    Math.min(arr.length - 1, i)
  );
};

// console.log(ExponentialSearch([1, 2, 3, 4, 5, 6, 7, 8, 9], 3)); //=> 2
// console.log(ExponentialSearch([1, 2, 3, 4, 5, 6, 7, 8, 9], 10)); //=> -1
