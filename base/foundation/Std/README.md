1.stdout, stdin, stderr的中文名字分别是标准输出，标准输入和标准错误。
2.在默认情况下，stdout是行缓冲的，他的输出会放在一个buffer里面，只有到换行的时候，才会输出到屏幕
3.而stderr是无缓冲的，会直接输出.
4.stdout -- 标准输出设备 (printf("..")) 同 stdout。
stderr -- 标准错误输出设备
两者默认向屏幕输出。
但如果用转向标准输出到磁盘文件，则可看出两者区别。stdout输出到磁盘文件，stderr在屏幕。