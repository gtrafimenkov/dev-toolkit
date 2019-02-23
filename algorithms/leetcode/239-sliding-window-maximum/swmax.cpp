// SPDX-License-Identifier: MIT
// Copyright (c) 2019 Gennady Trafimenkov

#include <vector>
#include <iostream>

using namespace std;

class Solution {
public:
    vector<int> maxSlidingWindow(vector<int>& nums, int k) {
        vector<int> results;

        if (nums.size() == 0) {
            return results;
        }

        int length = nums.size();

        int i = 1;
        int prevMax;

        // calculating the first max
        {
            int max = nums[0];
            for (; i < k; i++) {
                if (nums[i] > max) {
                    max = nums[i];
                }
            }
            results.push_back(max);
            prevMax = max;
        }

        for (; i < length; i++) {
            int leaving = nums[i-k];
            if (leaving < prevMax) {
                // no need to recalculate max of the rest of the window
                if (nums[i] > prevMax) {
                    prevMax = nums[i];
                }
            } else {
                int max = nums[i-k+1];
                for (int j = i - k + 2; j <= i; j++) {
                    if (nums[j] > max) {
                        max = nums[j];
                    }
                }
                prevMax = max;
            }
            results.push_back(prevMax);
        }
        return results;
    }

    vector<int> maxSlidingWindowSimple(vector<int>& nums, int k) {
        vector<int> results;

        if (nums.size() > 0) {
            int end = nums.size() - k;
            for (int start = 0; start <= end; start++) {
                int max = nums[start];
                int windowEnd = start + k;
                for (int j = start + 1; j < windowEnd; j++) {
                    if (nums[j] > max) {
                        max = nums[j];
                    }
                }
                results.push_back(max);
            }
        }

        return results;
    }
};

void printVector(vector<int> v) {
    for(int n : v) {
        std::cout << n << ' ';
    }
    std::cout << endl;
}

int main() {
    vector<int> v1 = {1, 3, -1, -3, 5, 3, 6, 7};
    vector<int> v2;
    Solution sol;
    printVector(sol.maxSlidingWindow(v1, 3));
    printVector(sol.maxSlidingWindow(v2, 3));
    return 0;
}
