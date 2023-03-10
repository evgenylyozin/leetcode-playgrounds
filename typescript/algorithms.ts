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

// FibonacciSearch - search in ranges defined by fibonacci sequence
const FibonacciSearch = (arr: number[], whatToSearchFor: number) => {
  let fibMm2 = 0;
  let fibMm1 = 1;
  let fibM = fibMm2 + fibMm1;

  while (fibM < arr.length) {
    [fibMm2, fibMm1] = [fibMm1, fibM];
    fibM = fibMm2 + fibMm1;
  }

  let offset = -1;

  while (fibM > 1) {
    let i = Math.min(fibMm2 + offset, arr.length - 1);

    if (arr[i] < whatToSearchFor) {
      fibM = fibMm1;
      fibMm1 = fibMm2;
      fibMm2 = fibM - fibMm1;
      offset = i;
      continue;
    }
    if (arr[i] > whatToSearchFor) {
      fibM = fibMm2;
      fibMm1 = fibMm1 - fibMm2;
      fibMm2 = fibM - fibMm1;
      continue;
    }
    return i;
  }
  if (fibMm1 && arr[arr.length - 1] === whatToSearchFor) {
    return arr.length - 1;
  }
  return -1;
};

// console.log(FibonacciSearch([1, 2, 3, 4, 5, 6, 7, 8, 9], 3)); //=> 2
// console.log(FibonacciSearch([1, 2, 3, 4, 5, 6, 7, 8, 9], 10)); //=> -1

// BubbleSort
const BubbleSort = (arr: number[]) => {
  for (let i = 0; i < arr.length - 1; i++) {
    for (let j = 0; j < arr.length - i - 1; j++) {
      if (arr[j] > arr[j + 1]) {
        [arr[j], arr[j + 1]] = [arr[j + 1], arr[j]];
      }
    }
  }
  return arr;
};

// console.log(BubbleSort([3, 2, 4, 1, 0])); // => [0,1,2,3,4]

//SelectionSort
const SelectionSort = (arr: number[]) => {
  for (let i = 0; i < arr.length - 1; i++) {
    let min = i;
    for (let j = i + 1; j < arr.length; j++) {
      if (arr[j] < arr[min]) {
        min = j;
      }
    }
    [arr[i], arr[min]] = [arr[min], arr[i]];
  }
  return arr;
};

// console.log(SelectionSort([3, 2, 4, 1, 0])); // => [0,1,2,3,4]

//InsertionSort
const InsertionSort = (arr: number[]) => {
  for (let i = 1; i < arr.length; i++) {
    let key = arr[i];
    let j = i - 1;
    while (j >= 0 && arr[j] > key) {
      arr[j + 1] = arr[j];
      j = j - 1;
    }
    arr[j + 1] = key;
  }
  return arr;
};

//console.log(InsertionSort([3, 2, 4, 1, 0])); // => [0,1,2,3,4]

const partition = (arr: number[], left: number, right: number) => {
  const pivot = arr[right];
  let index = left - 1;

  for (let i = left; i <= right - 1; i++) {
    if (arr[i] < pivot) {
      index++;
      [arr[index], arr[i]] = [arr[i], arr[index]];
    }
  }
  [arr[index + 1], arr[right]] = [arr[right], arr[index + 1]];
  return index + 1;
};

// QuickSort ...
const QuickSort = (arr: number[], left: number, right: number) => {
  if (left < right) {
    const pivotIndex = partition(arr, left, right);
    QuickSort(arr, left, pivotIndex - 1);
    QuickSort(arr, pivotIndex + 1, right);
  }
};

// const arr = [3, 2, 4, 1, 0];
// QuickSort(arr, 0, 4);
// console.log(arr); // => [0,1,2,3,4]

// MergeSort ...
const merge = (arr: number[], left: number, middle: number, right: number) => {
  const n1 = middle - left + 1;
  const n2 = right - middle;

  const L = new Array(n1);
  const R = new Array(n2);

  for (let i = 0; i < n1; i++) L[i] = arr[left + i];
  for (let j = 0; j < n2; j++) R[j] = arr[middle + 1 + j];

  let i = 0;
  let j = 0;
  let k = left;
  while (i < n1 && j < n2) {
    if (L[i] <= R[j]) {
      arr[k] = L[i];
      i++;
    } else {
      arr[k] = R[j];
      j++;
    }
    k++;
  }

  while (i < n1) {
    arr[k] = L[i];
    i++;
    k++;
  }

  while (j < n2) {
    arr[k] = R[j];
    j++;
    k++;
  }
};

const MergeSort = (arr: number[], left: number, right: number) => {
  if (left >= right) {
    return;
  }
  const middle = left + Math.floor((right - left) / 2);
  MergeSort(arr, left, middle);
  MergeSort(arr, middle + 1, right);
  merge(arr, left, middle, right);
};

// const arr = [3, 2, 4, 1, 0];
// MergeSort(arr, 0, 4);
// console.log(arr); // => [0,1,2,3,4]
