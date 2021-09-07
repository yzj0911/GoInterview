## 简介

datto可以在首次备份后记录块变化，并在之后的备份时根据块变化进行快速备份。这些操作可在不影响服务器运行的情况下进行。

## 安装

---

1. [在线安装](https://www.notion.so/Datto-c30d1e003f9045009b4f49c18394b538)

   [INSTALL.md](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/ad31577b-5038-4a9d-b566-7662f21d8081/INSTALL.md)

2. 离线安装

   [dbdctl.8](https://www.notion.so/dbdctl-8-aa1779dd9fa748198110bdaf1d3a2bb1)

   [datto 离线安装](https://www.notion.so/datto-d27bf82e5a3e4654b4825e5d29f58a40)

## 使用

---

1. 安装驱动及其相关工具
2. 创建快照

    ```bash
    # dbdctl setup-snapshot [-c <cache size>] [-f <fallocate>] <block device> <cow file path> <minor>
    dbdctl setup-snapshot /dev/sda1  /.datto 0
    ```

   该命令对根目录的磁盘/dev/sda1创造一个快照(`/dev/datto0`)，并生成在根目录生成一个COW(copy-on-write)文件/.datto。
   
   0 为 minor，是用于标识快照的唯一符

3. 将该快照进行备份

    ```bash
    dd if=/dev/datto0 of=/backups/sda1-bkp bs=1Mj
    ```

   此时快照已被备份至磁盘/backups/sda1-bkp

4. 将快照设置成增加模式(incremental mode)

    ```bash
    dbdctl transition-to-incremental 0
    ```

   此时快照(`/dev/datto0`)会进入增加模式，datto驱动会追踪块变化的地址，并且记录至COW文件/.datto。

5. 正常使用系统。在第一次备份后（步骤3）并当快照设置成增加模式（步骤4）后。
6. 将增加模式的快照设置成快照模式（snapshot mode)

    ```bash
    dbdctl transition-to-snapshot /.datto1 0
    ```

   该命令将新生成一个COW文件/.datto1，原来的COW文件/.datto保存了自从上次快照后的磁盘块变化，以用于**增量备份**。

7. 将变化更新至备份

    ```bash
    update-img /dev/datto0 /.datto /backup/sda1-bkp
    ```

   此时update-img工具来更新备份，用到了快照(`/dev/datto0`)此时快照内容已经是步骤6时的内容，即包含了新的修改内容，COW文件(`/.datto`)保存了块变化列表和原始备份(`/backups/sd1-bkp`)。update-img通过查看COW文件的块变化地址，直接复制快照内容至备份盘以此来快速更新备份数据。

8. 清除不需要的文件

    ```bash
    rm /.datto
    ```

9. 回到步骤4继续重复

   ![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/a6027c7d-0c8b-4566-89b5-0e2d7096c66f/Untitled.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/a6027c7d-0c8b-4566-89b5-0e2d7096c66f/Untitled.png)

## 相关链接

---

[INSTALL.md](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/ad31577b-5038-4a9d-b566-7662f21d8081/INSTALL.md)

[dbdctl.8](https://www.notion.so/dbdctl-8-aa1779dd9fa748198110bdaf1d3a2bb1)

[datto/dattobd](https://github.com/datto/dattobd)