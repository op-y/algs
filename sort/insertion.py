#!/usr/bin/python3

def insertion_sort(arr):
    l = len(arr)
    for i in range(l):
        idx = i - 1
        current = arr[i]
        print(idx, current, 'ok1')
        while idx >= 0 and arr[idx] > current:
            arr[idx+1] = arr[idx]
            print(idx, current, 'ok2')
            idx -= 1
        arr[idx+1] = current
        print(arr, 'ok3')
    return arr


if __name__ == '__main__':
    arr = [54, 26, 90, 70, 46, 30, 46, 50, 20]
    print("old list: ", arr)
    insertion_sort(arr)
    print("new list: ", arr)


