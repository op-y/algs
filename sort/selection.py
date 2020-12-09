#!/usr/bin/python3

def selection_sort(arr):
    l = len(arr)
    for i in range(l - 1):
        minIdx = i
        for j in range(i + 1, l):
            if arr[j] < arr[minIdx]:
                minIdx = j
        if i != minIdx:
            arr[i], arr[minIdx] = arr[minIdx], arr[i]
    return arr


if __name__ == '__main__':
    arr = [54, 26, 90, 70, 46, 30, 46, 50, 20]
    print("old list: ", arr)
    selection_sort(arr)
    print("new list: ", arr)
