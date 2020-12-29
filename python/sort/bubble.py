#!/usr/bin/python3

def bubble_sort(arr):
    n = len(arr)
    for i in range(n - 1):
        count = 0
        for j in range(0, n - 1 - i):
            if arr[j] > arr[j + 1]:
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
                count += 1
        if 0 == count:
            break
    return arr


if __name__ == '__main__':
    arr = [54, 26, 90, 70, 46, 30, 46, 50, 20]
    print("old list: ", arr)
    bubble_sort(arr)
    print("new list: ", arr)

