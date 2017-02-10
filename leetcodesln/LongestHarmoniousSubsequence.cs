﻿using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace leetcodesln
{
    public class LongestHarmoniousSubsequence
    {
        public int FindLHS(int[] nums)
        {
            var excemption = new List<int>();
            var maxlength = 0;
            for (int i = 0; i < nums.Length; i++)
            {
                if (excemption.Contains(nums[i])) continue;
                if (nums.Contains(nums[i] - 1))
                {
                    var min = nums[i] - 1;
                    var max = nums[i];
                    var localmax = 0;
                    for (int j = i; j < nums.Length; j++)
                    {
                        if (min <= nums[j] && nums[j] <= max)
                        {
                            localmax++;
                            
                        }
                    }
                    excemption.Add(nums[i]);
                    maxlength = maxlength > localmax ? maxlength : localmax;
                }
                if (nums.Contains(nums[i] + 1))
                {
                    var min = nums[i];
                    var max = nums[i]+1;
                    var localmax = 0;
                    for (int j = i; j < nums.Length; j++)
                    {
                        if (min <= nums[j] && nums[j] <= max)
                        {
                            localmax++;
                            
                        }
                    }
                    excemption.Add(nums[i]);
                    maxlength = maxlength > localmax ? maxlength : localmax;
                }
            }
            return maxlength;
        }
    }
}