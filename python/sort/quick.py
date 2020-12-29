#!/usr/bin/python3

def quick_sort(arr, left=None, right=None):
    left = 0 if not isinstance(left, int) else left
    right = len(arr) - 1 if not isinstance(right, int) else right
    if left < right:
        pidx = partition(arr, left, right)
        quick_sort(arr, left, pidx-1)
        quick_sort(arr, pidx+1, right)
    return arr


def partition(arr, left, right):
    pivot = left
    idx = pivot + 1
    i = idx
    while i <= right:
        if arr[i] < arr[pivot]:
            arr[i], arr[idx] = arr[idx], arr[i]
            idx += 1
        i += 1
    arr[pivot], arr[idx-1] = arr[idx-1], arr[pivot]
    return idx - 1


def quickthreeway_sort(arr, left=None, right=None):
    pass


if __name__ == '__main__':
    arr = [54, 26, 90, 70, 46, 30, 46, 50, 20]
    print("old list: ", arr)
    quick_sort(arr)
    print("new list: ", arr)
    # quickthreeway_sort(arr)
    # print("three new list: ", arr)
