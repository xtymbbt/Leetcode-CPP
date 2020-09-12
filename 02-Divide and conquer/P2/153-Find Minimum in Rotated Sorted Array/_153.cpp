//
// Created by Bridge_Wang on 2020/9/12.
//

#include <iostream>
#include <string>
#include <vector>
#include <map>
#include <unordered_map>
using namespace std;

class Solution {
public:
    int findMin(vector<int>& nums) {
        int min = nums[0];
        for (int i = 0; i < nums.size(); ++i) if (min > nums[i]) return nums[i];
        return min;
    }
};