# HomeSweetHome

通过读取 Warframe 的游戏日志，判断本次任务是否存在玩家想要的地形

# 免责声明

仅供学习交流，严禁用于商业用途，请于24小时内删除  
请勿将本程序用于非法用途，由此产生的一切后果与作者无关  
用户应对本程序中的内容自行加以判断，并承担因使用本程序引起的所有风险  

# 作者的话

- 本程序具有一定风险，请谨慎使用

- 日志中具有玩家的游戏 ID，分享游戏日志给他人时请谨慎

- 如果喜欢本程序，麻烦给个 Star 吧

- 祝大家
  - 每天登录都有 75% 折扣
  - 每次执刑官、衰退室都能出 TAU 石头
  - 每天的突击都出传说核心
  - 开出的紫卡都是极品
  - 开核桃全是金


# TODO

- [ ] 优化读取日志
> 逆序读取的时候就把地形信息存储起来，等获取完具体任务之后再对地形进行循环判断，从而只需一次逆序读取

- [x] 支持其他星图
> 需要大家自行编辑 yaml 配置文件，自行查找星图、任务、地形

- [ ] 支持查找金箱子
> 目前还不知道日志中是否有金箱子的记录。如果有，他应该不是地形，那就需要寻找指定资源的功能

- [ ] 分析任务的地形分布
> 在日志中具有地形分布的记录，例如 `generating layout with segments: SCI[CO]CI[CIC]CEP`，但是暂时不清楚具体是怎么样的

# 打赏

本程序完全 **免费开源** ，请您确认是否自愿打赏

[爱发电](https://afdian.net/a/Helzoph)

![赞赏码](/pic/微信赞赏码.png)