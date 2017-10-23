统计资产信息，包括主机和显示器，可以通过post方法将主机和显示器的信息上传到mysql数据库，
实现了增删改查的功能。查询时，通过mac地址、ip地址、用户名都可以。
程序中，为了找到指定ip对应的mac地址，通过snmp协议查询核心交换机得到，这些信息保存在文件中，
所以这个文件需要定期更新。
查询方法：
     snmpwalk -v 2c -c communityname switch-ip-addr .1.3.6.1.2.1.3.1.1.2 > ipvsmac.txt

