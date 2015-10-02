MomentAPI
=========
## 运行 ##
```
go get github.com/soyking/MomentAPI
```
数据库：MomentAPI.mongo


## 朋友圈操作 ##

/moment
- `POST`
>发布朋友圈，数据格式：
>```
>{
>      "UserId": "user1",		// 发布者id
>      "Type": 0,				// 朋友圈类型，纯文字、视频、分享
>      "Text": "it is rainy",	// 文字内容
>      "OtherData": "",		    // 除了文字外的内容，这里统一用字符串表示
>      "Source": "",			    // 朋友圈来源
>}
```
返回数据格式：（下面其他POST、DELETE方法返回格式也是如下）
```
{Result: success}	//失败为error
```
- `GET`
>获取用户id在指定时间戳之后的朋友圈，用来显示个人朋友圈主页，参数：
```
/moment?UserId=user1&Timestamp=100
```
返回数据格式（按时间戳降序）：
```
{
  "Result": "success",							// 获取结果，错误为error
  "Moments": [									// 多条朋友圈内容，意义同上
    {
      "MomentId": "560bfe69e6e90f2e8a000001",
      "UserId": "user1",
      "Type": 0,
      "Text": "it is rainy",
      "OtherData": "",
      "Source": "",
      "Timestamp": 150
    }
  ]
}
```
- `DELETE`
>通过朋友圈id和用户id删除朋友圈，数据格式：
```
{
	"MomentId":"560c8ea3e6e90f0cc1000001"，
    "UserId":"user1"
}
```

/moment/pull
- `GET`
>通过用户id和时间戳拉取好友朋友圈，参数：
```
/moment/pull?UserId=user3&Timestamp=200
```
返回数据格式同上

/moment/exist
- `POST`
>通过朋友圈id查询是否存在，用于更新，参数：
```
/moment/exist?MomentId=560c8944e6e90f5631000003
```
返回数据格式：
```
{Result:exist}	//不存在 not exist；错误 error
```

## 评论操作 ##

/comment
- `POST`
>发布评论，数据格式：
```
{
      "MomentId": "560bfc56e6e90f5178000001",	// 朋友圈id
      "FromUserId": "user1",					// 发布者id
      "ToUserId": "",							// 回复对象id，空表示没有回复对象
      "Text": "yeah,cold",					    // 回复内容
}
```
- `GET`
>通过朋友圈id和用户id获取的评论，用户id用来剔除非好友的评论，参数：
```
/comment?MomentId=560bfc56e6e90f5178000001&UserId=user3
```
返回数据格式（按时间戳升序）：
>```
{
  "Result": "success",
  "Comments": [									// 多条评论
    {
      "CommentId": "560c0ce0e6e90f6001000002",
      "MomentId": "560bfc56e6e90f5178000001",
      "FromUserId": "user1",
      "ToUserId": "",
      "Text": "yeah,cold",
      "Timestamp": 160
    }
  ]
}
```
- `DELETE`
>通过评论id和用户id删除评论，数据格式：
```
{
	"CommentId":"560e5338e6e90f1482000004",
    "UserId":"user1"
}
```

## 点赞操作 ##

/like
- `POST`
>点赞，数据格式：
```
{
      "MomentId": "560bfc56e6e90f5178000001",	// 朋友圈id
      "UserId": "user2",						// 发布者id
}
```
- `GET`
>通过朋友圈id和用户id获取赞，用户id用来剔除非好友的赞，参数：
```
/like?MomentId=560bfc56e6e90f5178000001&UserId=user2
```
返回数据格式（按时间戳升序）：
```
{
  "Result": "success",
  "Likes": [									 //多条赞
    {
      "LikeId": "560c116ce6e90f1841000001",
      "MomentId": "560bfc56e6e90f5178000001",
      "UserId": "user2",
      "Timestamp": 160
    }
  ]
}
```
- `DELETE`
>通过赞id和用户id取消赞，数据格式:
```
{
	"LikeId":"560e5879e6e90f1557000001",
    "UserId":"user1"
}
```

## 好友操作 ##

/block
- `POST`
>FollowId屏蔽FollowedId，数据格式：
```
{
	"FollowId":"user1",
    "FollowedId":"user2"
}
```
- `DELETE`
> FollowId取消屏蔽FollowedId，数据格式同上

/unshare
- `POST`
>FollowId不让FollowedId看自己朋友圈，数据格式同上
- `DELETE`
> FollowId取消不让FollowedId看自己朋友圈，数据格式同上
