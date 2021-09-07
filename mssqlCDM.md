
#Mssql CDM项目


## Mssql CDM（物理级快照备份）

### 环境准备：

----
1.windwos主要安装win.zip的包（包含vc_redist.x64.exe 等）

2.将VSSagent.exe 以及他的依赖包（AlpHaVSS.Common.dll 包拉进来）

Vss操作文档： 找鑫鑫要（VssAgent代码结构说明.docx)

### 创建应用：

----
通过 github.com/denisenkom/go-mssqldb 连接数据库(账号密码)

不采用 dbodb 的方式 因为没有长连接参数。

### 创建策略

-----
（*注意：选择一个或多个mssql数据库）

### 备份恢复流程
#### **备份流程**：

----

1. 检测lblet是否有数据库，并获得mssql的mdf和ldf文件

         sql语句:
            select type,name,physical_name as filePath,size*8 as sizeKB from 
            (数据库名字).sys.database_files order by file_id asc

2. 创建快照

         vssagent create C:\ R:\

3. 备份快照盘中，mssql的mdf和ldf以及ndf文件

         语句：vssagent onlybackup xx(数据库本身路径) xx(快照盘符) xx(备份路径) xx(是否开启Mointor True/false) xx(vss日志路径) ）
            
         ***注意：没必要开启监听，由于单文件备份，是通过比较md5值来判断是否需要增量备份，因此不需要开，vss内部也不会开启**

4. 备份用户，只备份创建的用户，不包括系统用户


5. 结束

#### **恢复流程**

-----
1. 判断mssql文件的备份文件是否存在


2. 通过mssql的创建数据库操作
   
         create databse [newName]on (filename="xx\xx.mdf"),(filename="xx\xx.ldf")...

####删除流程

-----

      1.脱机操作：(为了删除数据库的时候，不删除对应的ldf文件)
         sql语句：alter database [newName] set offline 
      2.删除数据库
         sql语句：drop databse [newName]

###恢复策略流程

----

   1. 若是第一次恢复策略，和恢复流程一样(如何判断，通过上个黄金副本是否存在)
   2. 若不是第一次恢复，则删除上次的恢复（和删除流程一样），在执行恢复流程

## Mssql 数据备份（恢复任意时间点）

###环境准备

----

目前，只测试了 mssql 2008,2012,2014。客户环境以2008,2014居多


###创建应用

----
通过 github.com/denisenkom/go-mssqldb 连接数据库（账号密码）,
不采用 dbodb 的方式 因为没有长连接参数

###创建策略

----
（*注意：选择一个或多个mssql数据库）

###备份恢复流程

----
备份原理: (https://blog.csdn.net/Hehuyi_In/article/details/90299012) 文档中可以知道，每次mssql是差量备份，目前还没有找到合适方式

####备份流程：
   1. 检测数据库是否存在，先将物理卷备份下来 
   2. 备份：
      
            1.全量备份：
               1.若目录中存在 日志目录，删除
               2.alter Database [数据库] set recovery full with no_wait //设置mssql 还原方式为 全量
               3.backup databsae [数据库] to disk =n'文件名称' with init,stats=5; //设置备份的文件名称
            2.差量备份:
               1.backup log [数据库]to disk=‘文件’ with init,stats=5;先备份日志，将备份盘中的日志刷新到最新
               2.alter Database [数据库] set recovery full with no_wait 设置数据库格式为全量
               3.若以前存在增量，则删除
            3.日志备份
               1.backup log [数据库]to disk=‘文件’ with init,stats=5;先备份日志

####恢复流程

----

1. 先将物理盘符恢复，由于逻辑名称与备份逻辑名称不同，因此目前先备份物理卷。
2. 设置数据库恢复格式为全量
3. 恢复全量和增量文件
   

      restore database [名称] from disk ='文件' with norecovery，replace //恢复增量
   
4. 恢复日志文件


      restore log [名称] from disk ='文件' with norecovery //恢复日志
   
5. 恢复日志文件到任意时间点
   

      restore log [名称] from disk ='文件' with stopAt='时间（2006-01-02 11:22:33）',recovery //恢复时间点日志 
   
####删除流程
1. 脱机操作：(为了删除数据库的时候，不删除对应的ldf文件)
   
         alter database [newName] set offline
   
2. 删除数据库 
   
         drop databse [newName]

####恢复策略
注意：目前不支持!!!!!!!!

