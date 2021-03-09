# cobra使用手册
# 1.安装
    go get -v -u [github.com/spf13/cobra/cobra](http://github.com/spf13/cobra/cobra)
# 2.初始化一个项目
    cobra init --pkg-name demo
# 3.增加一个方法
    cobra add name
# 4.获取参数
    myCmd.Flags().StringVarP(&path, "path", "p", "", "the log file dir path")
# 5.运行一个程序
    go build -o myctl main.go
    长的参数用两个-，短的用一个-
    ./myctl apply --path /var/log或者./myctl apply -p /var/log
    
    