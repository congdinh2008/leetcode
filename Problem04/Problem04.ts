function findMedianSortedArrays(nums1: number[], nums2: number[]): number {
    const len1 = nums1.length;
    const len2 = nums2.length;
    const resultLen = len1 + len2;
    const result: number[] = new Array(resultLen);

    let i = 0, j = 0, k = 0;
    while (i < len1 || j < len2) {
        if (i < len1 && (j >= len2 || nums1[i] <= nums2[j])) {
            result[k++] = nums1[i++];
        } else if (j < len2) {
            result[k++] = nums2[j++];
        }
    }

    if (resultLen % 2 === 1) {
        return result[Math.floor(resultLen / 2)];
    } else {
        return (result[resultLen / 2] + result[resultLen / 2 - 1]) / 2;
    }
}

function findMedianSortedArraysTwoPointer(nums1: number[], nums2: number[]): number {
    const len1 = nums1.length;
    const len2 = nums2.length;
    const total = len1 + len2;

    let i = 0, j = 0, median1 = 0, median2 = 0;
    for (let k = 0; k <= total / 2; k++) {
        median2 = median1;
        if (i === len1 || (j < len2 && nums1[i] > nums2[j])) {
            median1 = nums2[j++];
        } else {
            median1 = nums1[i++];
        }
    }

    if (total % 2 === 1) {
        return median1;
    } else {
        return (median1 + median2) / 2;
    }
}

const nums1 = [3, 4];
const nums2: number[] = []; // Example: [1,2,3,4,5,9]
console.log(findMedianSortedArraysTwoPointer(nums1, nums2));