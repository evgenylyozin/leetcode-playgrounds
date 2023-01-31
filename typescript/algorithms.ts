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
