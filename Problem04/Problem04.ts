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

const nums1 = [3, 4];
const nums2: number[] = []; // Example: [1,2,3,4,5,9]
console.log(findMedianSortedArrays(nums1, nums2));