1. 技术站 react + go + 可选nginx

2. 功能模块
    - 2.1 用户
        - 2.1.1 用户注册(头像上传)(可选邮箱或手机号绑定+支持验证码)                
        - 2.1.2 用户登录，进入个人中心，可进行修改
        - 2.1.3 管理员登录，进入个人中心，且可进入后台，进行数据管理。
    - 2.2 视频
        - 2.2.1 上传功能，在个人中心进行视频上传，视频上传后，管理员进行审核
        - 2.2.2 首页，视频筛选，分类，播放
    - 2.3 评论
        - 2.3.1 登录后，可进行评论
    - 2.4 关注/站内信/日志功能　待做

3. 模型
    - 用户(id, nickname, account, password, email, phone, isAdmin, photoFile(头像), createdTime)(微信)
    - 视频(id, userId, title, desc, category, label, rm, duration(时长), videoFile, createdTime, pass)
    - 评论(id, objId, model, context, createdTime)            

4. TODO
    - 4.1 用户关注
    - 4.2 用户消息
    - 4.3 视频弹幕
    - 4.4 用户空间大小，视频上传大小
    - 4.5 直播