# make_data_set_so-vits-svc

make data set for so-vits-svc

## Language

go and python

go responsible for scraping data, python is responsible for preprocessing the data

## Function

go:

1. [x] Grab all the works of Kugou's personal homepage and store them scientifically locally
2. [x] Restore the file name that has been processed by UltimateVocalRemover
3. [x] Some utility functions such as getting the intersection of files in two folders and copying them into the other folders

python:

4. [x] Scientific segmentation of audio
5. [ ] Solve the problem of multiple speakers with the help of Alibaba Cloud DAMO Academy's speaker recognition model
6. [ ] Have a human examine the dataset

## QA

The code for the dev branch is messy, don't expect too much until there is no main branch

## Copyright

The crawler itself does not have any copyright risk, and this project assumes that you have obtained written permission or that you are the author

If you've violated the copyright owner, remove content that you haven't authorized

## Law

依据现行的(2002年1月1日)《计算机软件保护条例》

>第十七条 为了学习和研究软件内含的设计思想和原理，通过安装、显示、传输或者存储软件等方式使用软件的，可以不经软件著作权人许可，不向其支付报酬。

来自大自然的馈赠