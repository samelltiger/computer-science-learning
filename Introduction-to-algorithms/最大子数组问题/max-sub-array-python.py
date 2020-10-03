'''
寻找A的和最大的非空连续子数组。我们称这样的连续子数组为最大子数组
'''
import math


# 找出穿过中点的最大子数组
def find_max_crossing_subarray(A, low, mid, high):
    left_sum = -math.inf
    _sum = 0
    max_left = -1
    for i in range(mid, low - 1, -1):
        _sum += A[i]
        if _sum > left_sum:
            left_sum = _sum
            max_left = i

    right_sum = -math.inf
    _sum = 0
    max_right = -1
    for i in range(mid + 1, high + 1):
        _sum += A[i]
        if _sum > right_sum:
            right_sum = _sum
            max_right = i

    return max_left, max_right, left_sum + right_sum


def find_maximum_subarray(A, low, high):
    if low == high:
        return low, high, A[low]
    else:
        mid = int((low + high) / 2)
        left_low, left_high, left_sum = find_maximum_subarray(A, low, mid)
        right_low, right_high, right_sum = find_maximum_subarray(A, mid + 1, high)
        cross_low, cross_high, cross_sum = find_max_crossing_subarray(A, low, mid, high)

        if left_sum >= right_sum and left_sum >= cross_sum:
            return left_low, left_high, left_sum

        if right_sum >= left_sum and right_sum >= cross_sum:
            return right_low, right_high, right_sum

        if cross_sum >= left_sum and cross_sum >= right_sum:
            return cross_low, cross_high, cross_sum


if __name__ == "__main__":
    a = [13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7]
    low = 0
    high = len(a) - 1
    mid = int((low + high) / 2)
    # result = find_max_crossing_subarray(a, low, mid, high)
    result = find_maximum_subarray(a, low, high)
    print(result)
    print(a[result[0]:result[1] + 1])
