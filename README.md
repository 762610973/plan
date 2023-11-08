<!-- TOC -->
* [PLAN](#plan)
  * [Star](#star)
  * [Books](#books)
    * [深入理解计算机系统](#深入理解计算机系统)
    * [现代操作系统](#现代操作系统)
    * [计算机网络自顶向下](#计算机网络自顶向下)
    * [数据库系统概论](#数据库系统概论)
  * [Rust](#rust)
    * [深入浅出Rust](#深入浅出rust)
  * [Go](#go)
    * [Article](#article)
    * [Cache](#cache)
    * [Source](#source)
    * [设计模式](#设计模式)
      * [刘丹冰](#刘丹冰)
      * [鸟窝](#鸟窝)
  * [前端](#前端)
    * [HTML](#html)
    * [CSS](#css)
    * [JavaScript](#javascript)
    * [Vue](#vue)
  * [Docker](#docker)
    * [深入浅出Docker](#深入浅出docker)
  * [K8s](#k8s)
  * [汇编](#汇编)
  * [Redis](#redis)
  * [MySQL](#mysql)
    * [MySQL实战45讲](#mysql实战45讲)
    * [高性能MySQL](#高性能mysql)
  * [ES](#es)
  * [Nginx](#nginx)
  * [ETCD](#etcd)
  * [CPP](#cpp)
  * [算法](#算法)
    * [Hello Algorithm](#hello-algorithm)
    * [LeetCode](#leetcode)
      * [HOT100](#hot100)
      * [代码随想录](#代码随想录)
      * [剑指offer](#剑指offer)
      * [SQL](#sql)
<!-- TOC -->
# PLAN
> 记录定制的计划、进度、笔记

##  Star

- **<span style="font-size: 20px;">[Emoji](https://gist.github.com/rxaviers/7360908)</span>**	
- **<span style="font-size: 20px;">[Emoji](https://emojixd.com/)</span>**	
- **<span style="font-size: 20px;">[Restful](https://restfulapi.cn/)</span>**	

## Books

### 深入理解计算机系统
- 2023.10.04:open_book: 第一章: 计算机系统漫游
- 2023.10.31:open_book: 第二章: 信息存储
### 现代操作系统

### 计算机网络自顶向下
- 2023.09.25:open_book: 第一章


### 数据库系统概论

## Rust
> **<span style="font-size: 18px;">[Rust镜像代理](https://rsproxy.cn/)</span>**
> 
> **<span style="font-size: 18px;">[MSVC](https://visualstudio.microsoft.com/zh-hans/visual-cpp-build-tools/)</span>**

### 深入浅出Rust
- 2023.07.24:blush: 第一章
- 2023.07.25:blush: 第二章:变量声明
- 2023.07.26:blush: 第二章:基本数据类型+复合数据类型
- 2023.07.28:blush: 第三章:语句+表达式
- 2023.07.29:blush: 第四章:函数
- 2023.07.31:blush: 第五章:`trait`成员方法+静态方法
- 2023.08.09:blush: 第五章:`trait`
- 2023.08.11:blush: 第六章:`array-slice-string`
- 2023.08.12:blush: 第七章:`pattern-destructure`
- 2023.08.12:blush: 第八章:类型系统初探
- 2023.08.18:blush: 第九章:Rust宏初探, 太强大了!
- 2023.09.15:blush: 第十章: 内存管理基础, 第十一章: 所有权和移动语义
- 2023.09.15:blush: 第十二章: 借用和生命周期

## Go

### Article
- 2023.07.24:blush: [uber-go-guide](https://github.com/xxjwxc/uber_go_guide_cn)
- 2023.08.14:blush: [go的net/http有哪些值得关注的细节](https://mp.weixin.qq.com/s/QfeycEFqeqqhRKrYYL5mGA)
- 2023.08.18:blush: [泛型](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzU2ODc4NzUxMg==&action=getalbum&album_id=2218002393592627201&scene=173&from_msgid=2247485263&from_itemidx=1&count=3&nolastread=1#wechat_redirect)
- 2023.10.03:blush: [slice](https://www.bilibili.com/video/BV1Bz4y1s7g1)
- 2023.10.03:blush: [slice-case](https://mp.weixin.qq.com/s/uNajVcWr4mZpof1eNemfmQ)
- 2023.11.07:blush: [曹大带我学Go](https://mp.weixin.qq.com/s/RmEtL869Sd-FjB_73tFSaw): 调度本质, goroutine执行顺序
- 2023.11.08:blush: [曹大带我学Go](https://mp.weixin.qq.com/s/RmEtL869Sd-FjB_73tFSaw): 初识AST的威力, 如何用汇编打同事的脸, 哪里来的goexit, 如何优雅地指定配置项

### Cache
- 2023.09.06:blush: init three projects
- 2023.09.19:blush: go-cache部分源码解读
- 2023.09.20:blush: go-cache源码阅读完成

### Source
- 2023.08.16:money_mouth_face: errors
- 2023.08.16:money_mouth_face: Context初探([参考视频](https://www.bilibili.com/video/BV1EA41127Q3))
- 2023.09.20:money_mouth_face: Context用例
- 2023.09.23:money_mouth_face: cancelCtx源码阅读完成
- 2023.09.25:money_mouth_face: gin源码阅读`mode.go,errors.go`
- 2023.09.26:money_mouth_face: `reflect testcase`
- 2023.09.05:money_mouth_face: 无缓冲channel读写顺序(终于明白之前不理解的一个点了)
- 2023.09.29:money_mouth_face: interface implement verify.
- 2023.09.30:money_mouth_face: net/http source code read.
- 2023.10.01:money_mouth_face: net/http source code read(learned a lot).
- 2023.10.01:money_mouth_face: channel case.
- 2023.10.31:money_mouth_face: `sync.Once`源码解读.
- 2023.11.06:money_mouth_face: `sync.WaitGroup`源码解读.
- 2023.11.07:money_mouth_face: `sync.Map`源码解读.

### 设计模式
#### 刘丹冰
- 2023.08.20:star2: 简单工厂模式
- 2023.08.21:star2: 工厂方法模式,抽象工厂方法模式
- 2023.08.22:star2: 单例模式: 饿汉+懒汉
- 2023.08.22:star2: 代理模式,装饰器模式
- 2023.08.23:star2: 适配器模式
- 2023.08.23:star2: 外观模式
- 2023.09.07:star2: 模板方法模式
- 2023.09.07:star2: 命令模式
- 2023.09.09:star2: 观察者模式
#### 鸟窝
- 2023.09.13:star2: 原型模式
- 2023.09.13:star2: 生产者-消费者模式
- 2023.09.13:star2: Go可以使用设计模式, 但绝不是<<设计模式>>中的那样, 真实世界的Go设计模式-工厂模式

## 前端
> **<span style="font-size: 18px;">[html+css](https://www.bilibili.com/video/BV1p84y1P7Z5)</span>**

### HTML
- 2023.07.24:blush:List and Table
- 2023.07.26:blush:Table and Common attributes
- 2023.07.27, 2023.7.28:blush:Form 
- 2023.07.29:blush:basically complete html4 and start css2
- 2023.08.27:blush:html5简介,新增布局标签,状态标签,列表标签
- 2023.08.28:blush:html5新增表单控件,type属性,视频标签,音频标签,全局属性,兼容性处理
### CSS
- 2023.07.31:blush:css style position
- 2023.07.31:blush:css 样式优先级和基本选择器
- 2023.08.08:blush:css子代选择器
- 2023.08.10:blush:css兄弟+属性选择器+伪类选择器
- 2023.08.11:blush:伪类选择器,否定伪类,目标伪类,语言伪类,UI伪类,伪元素选择器
- 2023.08.12:blush:选择器优先级
- 2023.08.13:blush:三大特性+像素
- 2023.08.14:blush:列表,表格,背景
- 2023.08.15:blush:背景,鼠标,长度单位
- 2023.08.16:blush:元素的显示模式
- 2023.08.17:blush:盒子模型的组成部分,内容区,内边距,边框,默认宽度
- 2023.08.18:blush:盒子的外边距,margin注意事项,塌陷问题,合并问题
- 2023.08.20:blush:样式继承,元素的空白问题,默认样式,如何隐藏,内容溢出,布局技巧
- 2023.08.21:blush:浮动简介
- 2023.08.22:blush:元素浮动之后的特点,影响,解决影响
- 2023.08.23:blush:浮动布局练习
- 2023.08.25:blush:固定定位,浮动定位,粘性定位,绝对定位
- 2023.08.26:blush:定位的层级,特殊应用,布局(常用类名,版心,重置默认样式)
- 2023.08.26:blush:尚品汇顶部导航条,头部,主导航,内容区,秒杀,楼层,页脚
- 2023.09.02:blush:css3新增盒子属性,背景属性
- 2023.09.04:blush:css3新增边框属性
- 2023.09.04:blush:css3新增文本属性: 文本阴影,换行,溢出,修饰,描边
- 2023.09.04:blush:css3新增渐变: 线性,径向,重复渐变
- 2023.09.05:blush:css3新增2D变换: 位移,缩放,旋转,扭曲,多重变换,变换原点
- 2023.09.10:blush:css3新增3D变换: 空间与景深, 透视点位置, 位移, 缩放, 旋转, 多重变换, 背部
- 2023.09.10:blush:css3新增过渡
- 2023.09.11:blush:css3动画, 多列布局
- 2023.09.12:blush:css3伸缩盒模型: 简介, 容器和项目, 主轴方向, 换行方式, 对齐方式, 侧轴, 基准长度, 元素水平垂直居中, 伸缩性
- 2023.09.12:blush:css3响应式布局, BFC

### JavaScript
> [JavaScript](https://www.w3school.com.cn/js/index.asp)
> [JavaScript](https://www.w3school.com.cn/js/index_pro.asp)
- 2023.09.05:dart: js的基本使用
- 2023.09.13:dart: js的输出, 语句, 语法, 注释, 变量, Let, const, 运算符


### Vue
> **<span style="font-size: 18px;">[React](https://www.bilibili.com/video/BV1Zy4y1K7SH)</span>**
> **<span style="font-size: 18px;">[React](https://cn.vuejs.org/)</span>**

- 2023.09.22:camping: Vue简介, 环境搭建, 模板语法, 数据绑定, el与data的写法, MVVM模型, 数据代理
- 2023.09.23:camping: 事件处理
- 2023.09.28:camping: 计算属性, 监视属性, 绑定样式
- 2023.09.28:camping: vue终结:sob:

## Docker
### 深入浅出Docker
> [](./books/深入浅出Docker.pdf)

- 2023.10.02:whale2: 容器发展, 简介, 安装, 纵观
- 2023.10.04:whale2: Docker Engine
- 2023.10.04:whale2: Docker Image
- 2023.10.07:whale2: Docker Container
- 2023.10.11:whale2: 应用容器化
- 2023.10.13:whale2: Docker Compose
- 2023.10.17:whale2: Docker Swarm

## K8s

## 汇编
>  [汇编](https://www.bilibili.com/video/BV1Wu411B72F)、王爽老师的书

- 2023.07.22:keyboard: 第一节概论、环境搭建
- 2023.10.09:keyboard: 寄存器, mov和add指令, 内存的分段表示. Debug使用, CS,IP和代码段 
- 2023.10.10:keyboard: jmp指令, 内存中字的存储.

## Redis

## MySQL
### MySQL实战45讲
- 2023.11.02:floppy_disk: 第一讲和第二讲: SQL如何执行, update如何执行.
- 2023.11.05:floppy_disk: 第三讲: 事务隔离, 为什么你改了我还看不见.
- 2023.11.06:floppy_disk: 第四讲, 第五讲: 深入浅出索引.
- 2023.11.06:floppy_disk: 第六讲: 全局锁和表锁: 给表加个字段怎么有这么多阻碍.
- 2023.11.08:floppy_disk: 第七讲: 行锁功过: 怎么减少行锁对性能的影响.
### 高性能MySQL


## ES

- 2023.07.28:tiger:安装`elasticsearch`以及`kibana`和环境
- 2023.09.11:tiger:geek time elasticsearch简介
- 2023.09.16:tiger:geek time 文档, 索引, 分片

## Nginx

- 2023.09.04:watermelon: 环境搭建
- 2023.09.04:watermelon: 最小核心配置文件
- 2023.09.05:watermelon: 域名解析, 多hosts
- 2023.09.07:watermelon: 负载均衡和反向代理的基本概念与基本配置
- 2023.09.18:watermelon: 负载均衡策略, 动静分离

## ETCD
> 极客时间实战课

- 2023.09.12:cloud: 开篇词

## CPP
[](./CPP)
> C++ Primer

- 2023.09.27:atom_symbol: hello world
- 2023.09.28:atom_symbol: begin
- 2023.09.29:atom_symbol: 基本内置类型, 变量, 复合类型(引用和指针)

## 算法
### Hello Algorithm
> **<span style="font-size: 18px;">[hello-algo](https://www.hello-algo.com/)</span>**
- 2023.10.01:wind_chime:算法复杂度分析
- 2023.10.02:wind_chime:数据结构
- 2023.10.03:wind_chime:数组与链表
- 2023.10.04:wind_chime:栈与队列
- 2023.10.05:wind_chime:哈希表
- 2023.10.07:wind_chime:树
- 2023.10.08:wind_chime:堆
- 2023.10.09:wind_chime:图
- 2023.10.09:wind_chime:搜索
- 2023.10.09:wind_chime:排序
- 2023.10.09:wind_chime:分治, 动态规划, 贪心, 回溯

### LeetCode

#### HOT100
#### 代码随想录
#### 剑指offer
#### SQL
- 2023.10.16::rabbit2:175组合两个表, 176第二高的薪水.