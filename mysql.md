#mysql 备份cdm（Linux）

## cdm（物理级别备份）


###环境准备：

------
1.dattobd 安装 

2.dattobd 文档：https://www.notion.so/Datto-c30d1e003f9045009b4f49c18394b538

###备份原理：

--------
1.dattobd 磁盘备份（选择使用:由于增量模式由于LVM，而且LVM有局限性，只能在LVM磁盘上做操作）

2.LVM 磁盘备份（增量模块没有datobd强，局限性太大）


###流程：


1.创建 应用

-  检测mysql客户端是否存在（账号密码 方式:go-sql-driver/mysql）
    
2.创建策略

3.备份

（方式：使用dattobd cbt块级跟踪增量备份）

1.检测mysql是否存在，不存在报错

2.创建dattob快照（第一次 全量快照，第二次 增量模式转为快照模式）

 （通过判断 scripts/SnapshotLog/parameterFile.txt 文件是否存在，若存在则认为是第一次快照 ）（待修改）

（scripts/SnapshotLog/parameterFile.txt 文件中记录了 （快照 快照文件 备份地址 镜像 mysql路径 磁盘类型））（最好记录在数据库中）

3.备份（第一次dd命令,第二次update-img）
（如何判断第一次：备份盘符上是否存在datto 这个文件（需要优化））
4.sync 同步一下
5.存储创建黄金副本


    （windows）
        （方式：使用Vss快照，开启vss的监听 moinitor（监听信息,在vss的日志目录下））
        1.检测mysql是否存在，不存在报错
        2.通过mysql中（show global variables like "%datadir%"）来获得mysql路径并备份
        3.创建Vss快照（Vssagent create C:\ R:\）
        4.通过2中获得路径，并找到快照盘（R）盘中的路径。通过命令：
        （vssagent onlybackup xx(数据库本身路径） xx(快照盘符） xx(备份路径） xx(是否开启Mointor True/false） xx(vss的日志路径） ）
        *注意:是否成功，只能根据success判断，并且控制台输出进度等信息。
         若错误，则会在刚刚指定的vss目录下的error目录下，新建一个最新的日志文件，并记录错误
        5.sync 同步一下
        6.备份完成,删除快照


    4.恢复
    Linux：
    （方式：新启动进程）
        1.存储通过克隆黄金副本，并将克隆出来的磁盘挂载到客户端
        2.将dattobd 文件 挂载到客户端 （mountDatto.sh）
        3.在挂载盘符上，生成新配置文件在挂载盘符上 填写port等信息
        （*注意 这里需要填写 进程pid文件生成的路径）
        4.启动服务（mysqld --defaults-file=（新写的配置文件） --user=root）

    Windows
        （方式为：启动新服务）
        1.存储通过克隆黄金副本，并将克隆出来的磁盘挂载到客户端
        2.在挂载盘符上，生成行配置文件 填写port等信息
        3.注册服务
        （mysqld --install 服务名字 --defaults-file=新配置文件）
        4.启动服务


    5.删除恢复
    Linux:
        1.kill pid（pid 为当初备份时指向的pid文件中有存）（后续需要找好方法）
        2.卸载盘符...

    Windows:
        1.停止服务（net stop 服务名）
        2.取消注册（mysqld --remove 服务名字）
        3.卸载盘符....


    6.恢复策略
    （linux和windows 一样，都是覆盖目标机子上有的mysql）
        1.挂载磁盘...
        2.找到my.cnf（windows：my.ini） 
            命令：
            linux：默认 /etc/my.cnf （可以通过 mysql --verbose --help |grep -A 1 'Default options' 目前默认）
            windows：show global variables like  "%datadir"/"basedir"（其目录下存在my.ini 就采用）
        3.将查到的配置文件my.cnf（windows：my.ini）,修改为my.cnf.bak（windows：my.ini.bak）
        并将路径修改为挂载过来中的数据库路径（一般为 挂载磁盘/Data） 
        4.停止服务，启动服务


    7.删除恢复策略 
    （linux和windows 一样，都是停止服务，将原来的配置文件修改回来，启动服务）


##数据备份（任意时间点）

1.环境准备：

    Linux
        安装innobackupex（根据mysql版本装对应的innobackupex 待整理）
        innobackupex文档：
            https://www.percona.com/doc/percona-xtrabackup/2.1/innobackupex/creating_a_backup_ibk.html
        优势：
            1.拥有增量
            2.备份物理磁盘，绝对的快速

    Windows
        目前innobackupex 不支持windows版本，因此采用mysqldump的逻辑备份
2.备份原理（innobackupex）：

    原理：（http://blog.itpub.net/29654823/viewspace-2679287/）
    （https://cloud.tencent.com/developer/news/28511）
    （http://mysql.taobao.org/monthly/2016/03/07/）
3.流程：
 
    1.创建 应用
        1.检测mysql客户端是否存在（账号密码 方式:go-sql-driver/mysql）
    2.创建策略
    3.备份
    Linux:
        1.挂载磁盘...
        2.检测mysql是否存在，并获得mysqlbinlog的路径 在配置文件中获取 log-bin= （若为目录，则在目录下，若不为，则在datadir下生成）
        3.根据备份方式备份
            1.全量备份：
                1.全量：
                innobackupex --defult-file=（配置文件） --username=xx --password=xx --host=xx 
                --port=xx --no-timestamp=（备份到某个文件：挂载磁盘/all-bckup）

                2.记录mysqlbinlog的日志号 并记录到挂载磁盘的mysqlstartbinlog.txt （待修改）

                3.接着操作日志，将所有挂载磁盘下binlogBackup路径下所有日志清除

            2.增量备份： 
                1.刷新日志
                2.合成（通过判断incremental 文件是否存在来判断一下情况）
                    1.合成（存在增量文件）：
                    innobackupex --apply-log --redo-only （全量文件） --incremental-dir=（增量文件 需指定）

                    2.合并（不存在增量文件）：
                    innobackupex --apply-log --redo-only （全量文件）

                3.移除已经合并的增量文件
                4.增量备份：
                innobackupex --defult-file=（配置文件） --username=xx --password=xx --host=xx 
                --port=xx --no-timestamp --incremental --incremental-basedir="全量文件" 增量文件

                5.获得增量的binlog日志并备份
                6.刷新日志
                3.日志备份
                    1.找到binlog路径
                    2.备份日志
    Windows:
        1.获得mysql bin路径 （show global variables like "%basedir%"） 获得bin路径
        2.根据备份方式备份
            1.全量备份：（add-drop-database 是将删除和创建流程也备份到文件中）
                1.mysqldump -u user -p pwd -P port --databases 数据库 --add-drop-database 
                2.备份日志，找到binlog日志路径（通过配置文件找）
            2.增量备份：
                1.删除所有日志文件
                2.获得所有当前数据库的备份文件，比较备份时间，将过期的增量备份删除
                3.增量备份（备份方式和全量一样）

    4.恢复:
        1.获得mysql bin路径
        2.mysqldump 导出源数据库 database_DeleteRecover.sql（）
        3.删除原有数据库，并将全量备份导入（由于添加add-drop-database，因此会创建对应的数据库的字符集 ）
        4.过滤掉需要导入的文件
        5.导入增量
        6.导入日志（mysqlbinlog --start-datetime "xx" --stop-datetime "xx" --databse 数据库  日志文件）（开始时间是文件记录的时间）
        7.等待恢复完成

    5.删除恢复:
        1.删除数据库
        2.将恢复中导出的 database_DeleteRecover.sql 导入到数据库中

    6.恢复策略:
        目前不支持该功能

