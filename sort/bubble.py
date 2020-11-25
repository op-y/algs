#!/usr/bin/python3

def bubble_sort(alist):
    n = len(alist)
    for i in range(n - 1):
        count = 0
        for j in range(0, n - 1 - i):
            if alist[j] > alist[j + 1]:
                alist[j], alist[j + 1] = alist[j + 1], alist[j]
                count += 1
        if 0 == count:
            break


if __name__ == '__main__':
    alist = [54, 26, 90, 70, 46, 30, 46, 50, 20]
    print("old list: ", alist)
    bubble_sort(alist)
    print("new list: ", alist)

