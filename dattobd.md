* 安装指定版本最小系统


* 关闭selin
aa
  vi /etc/selinux/config
SELINUX=disabled
重启

* 修改yum安装不删除rpm包

vi /etc/yum.conf
keepcache=1

* 升级kernel

将准备好的kernel文件夹拷贝至目标机
进入文件夹执行: yum localinstall --nogpgcheck *.rpm
重启
查看kernel: uname -r，如果正确安装应该显示: 2.6.32-754.23.1.el6.x86_64或者3.10.0-1062.4.1.el7.x86_64

* 导出升级kernel所需的包

新建文件夹: mkdir /root/kernel
cd /var/cache/yum/x86_64/6/
cp base/packages/* /root/kernel/
cp extras/packages/* /root/kernel/
cp updates/packages/* /root/kernel/
将/root/kernel中的rpm再导出到原先安装的kernel文件夹中

* 安装datto源

yum localinstall https://cpkg.datto.com/datto-rpm/repoconfig/datto-el-rpm-release-$(rpm -E %rhel)-latest.noarch.rpm
安装后如果yum运行报错则修改: /etc/yum.repos.d/epel.repo，将#baseurl=去掉注释，mirrorlist=注释掉

* 安装datto

yum install dkms-dattobd dattobd-utils
重启

* 验证datto

dbdctl setup-snapshot /dev/mapper/VolGroup-lv_root /.datto 5 （不会报错）
mount -o,ro,norecovery /dev/datto5 /root/test （Contos6）
mount -o,ro,norecovery,nouuid /dev/datto5 /root/test （Contos7）
挂载成功说明datto安装成功
卸载快照盘:
umount -f /dev/datto5
dbdctl destroy 5

* 导出datto包

新建文件夹: mkdir /root/datto
cd /var/cache/yum/x86_64/6/
cp base/packages/* /root/datto/
cp datto-rpm/packages/* /root/datto/
cp epel/packages/* /root/datto/
cp extras/packages/* /root/datto/
cp updates/packages/* /root/datto/

* 至此kernel文件夹和datto文件夹就可以在新机上离线安装