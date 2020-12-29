#!/usr/bin/python3

def insertion_sort(arr):
    l = len(arr)
    for i in range(l):
        idx = i - 1
        current = arr[i]
        while idx >= 0 and arr[idx] > current:
            arr[idx+1] = arr[idx]
            idx -= 1
        arr[idx+1] = current
    return arr


def shell_sort(arr):
    gap = 1
    while (gap < len(arr) / 3):
        gap = gap * 3 + 1
    while gap > 0:
        for i in range(gap, len(arr)):
            tmp = arr[i]
            j = i - gap
            while j >= 0 and arr[j] > tmp:
                arr[j+gap] = arr[j]
                j -= gap
            arr[j+gap] = tmp
        gap = gap // 3


if __name__ == '__main__':
    arr = [54, 26, 90, 70, 46, 30, 46, 50, 20]
    print("old list: ", arr)
    insertion_sort(arr)
    print("insertion new list: ", arr)
    shell_sort(arr)
    print("shell new list: ", arr)


