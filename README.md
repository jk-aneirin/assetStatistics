主机和显示器信息采集<br>
功能介绍：<br>
采集员工姓名、所属部门、主机Mac地址和显示器SN到数据库系统，通过Web API接口展示这些信息；<br>
可以通过员工姓名、Mac地址或Ip地址来查询，当使用Ip地址查询时，如果数据库没有对应记录，则返回Mac地址。<br>
不足：<br>
非Windows主机，信息采集脚本不能使用；<br>
Windows系统非Windows 10系统采集脚本可能会运行异常（主机型号为Dell OptiPlex 7040、7050、7060，显示器型号为Dell U2412M、U2717D）；<br>
如果主机有虚拟网口，Mac地址可能会采集错误，需要人为更改；<br>
只能采集显示器的SN信息，新购入的显示器同时存在Service Tag和SN，Service Tag信息不能采集，采集到的SN信息中间有部分缺失，但是不会影响显示器SN信息的唯一性。<br>
注意：<br>
- 文件main/conf/app.conf必须包含copyrequestbody = true，否则不能post数据
- models.Entryinfo必须包含Id属性，即使将Name设置为主键也不行。否则不能在数据库中成功的创建表，应该是个bug
- beego orm只能通过主键删除表记录，不能使用名字这个非主键来删除，不是很便利
