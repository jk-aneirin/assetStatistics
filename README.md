注意：
文件main/conf/app.conf必须包含copyrequestbody = true，否则不能post数据
models.Entryinfo必须包含Id属性，即使将Name设置为主键也不行。否则不能在数据库中成功的创建表，应该是个bug
beego orm只能通过主键删除表记录，不能使用名字这个非主键来删除，不是很便利
