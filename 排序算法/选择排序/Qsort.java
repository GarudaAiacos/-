package com.zzm;

public class Qsort {

    public static void main(String[] args) {
        int[] arr = {6, 1, 2, 7, 9, 3, 4, 5, 10, 8};
        quickSort(arr, 0, arr.length-1);
        for (int i=0; i<arr.length; i++) {
            System.out.println(arr[i]);
        }
    }

    private static void quickSort(int[] arr, int start, int end) {
        int i, j, temp, t;
        if(start > end) {
            return;
        }
        i = start;
        j = end;
        //基准值
        temp = arr[start];

        while (i<j) {
            //先从右边开始找比arr[temp]小的值
            while (temp <= arr[j] && i<j) {
                j--;
            }
            //从左边找比arr[temp]大的值
            while (temp >= arr[i] && i<j) {
                i++;
            }

            //如果已经找到，则交换两个的值
            if(i<j) {
                t = arr[j];
                arr[j] = arr[i];
                arr[i] = t;
            }
        }
        //j和i相遇,将基准值和j所对应的值交换，完成第一轮交换
        arr[start] = arr[i];
        arr[i] = temp;

        quickSort(arr, start, j-1);
        quickSort(arr, j+1, end);
    }
}