package docs

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/swagger"
)

const (
    Rootinfo string = `{"apiVersion":"1.0.0","swaggerVersion":"1.2","apis":[{"path":"/moment","description":""},{"path":"/comment","description":""},{"path":"/like","description":""},{"path":"/block","description":""},{"path":"/unshare","description":""}],"info":{"title":"beego Test API","description":"beego has a very cool tools to autogenerate documents for your API","contact":"astaxie@gmail.com","termsOfServiceUrl":"http://beego.me/","license":"Url http://www.apache.org/licenses/LICENSE-2.0.html"}}`
    Subapi string = `{"/block":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/block","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/","description":"","operations":[{"httpMethod":"POST","nickname":"block","type":"","summary":"block someone","parameters":[{"paramType":"body","name":"body","description":"\"FollowId block FollowedId\"","dataType":"UserRelation","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"{Result:success}","responseModel":""},{"code":200,"message":"{Result:error}","responseModel":""}]}]},{"path":"/","description":"","operations":[{"httpMethod":"DELETE","nickname":"unblock","type":"","summary":"cancel block someone","parameters":[{"paramType":"body","name":"body","description":"\"FollowId cancel block FollowedId\"","dataType":"UserRelation","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"{Result:success}","responseModel":""},{"code":200,"message":"{Result:error}","responseModel":""}]}]}]},"/comment":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/comment","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/","description":"","operations":[{"httpMethod":"GET","nickname":"get","type":"","summary":"get comments by MomentId and UserId","parameters":[{"paramType":"query","name":"MomentId","description":"\"The moment you want to query\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"query","name":"UserId","description":"\"Your UserId\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"controllers.CommentResult","responseModel":"CommentResult"},{"code":200,"message":"{\"Result\": \"error\",\"Comments\": null}","responseModel":""}]}]},{"path":"/","description":"","operations":[{"httpMethod":"POST","nickname":"publish","type":"","summary":"publish comment","parameters":[{"paramType":"body","name":"body","description":"\"The comment content\"","dataType":"Comment","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"{Result:success}","responseModel":""},{"code":200,"message":"{Result:error}","responseModel":""}]}]},{"path":"/","description":"","operations":[{"httpMethod":"DELETE","nickname":"delete","type":"","summary":"delete the comment","parameters":[{"paramType":"body","name":"body","description":"\"The comment you want to delete\"","dataType":"CommentDeleteInfo","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"{Result:success}","responseModel":""},{"code":200,"message":"{Result:error}","responseModel":""}]}]}],"models":{"CommentResult":{"id":"CommentResult","properties":{"Comments":{"type":"array","description":"","items":{"$ref":"\u0026{models Comment}"},"format":""},"Result":{"type":"string","description":"","format":""}}}}},"/like":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/like","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/","description":"","operations":[{"httpMethod":"GET","nickname":"get","type":"","summary":"get likes by MomentId and UserId","parameters":[{"paramType":"query","name":"MomentId","description":"\"The moment you want to query\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"query","name":"UserId","description":"\"Your UserId\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"controllers.LikeResult","responseModel":"LikeResult"},{"code":200,"message":"{\"Result\": \"error\",\"Likes\": null}","responseModel":""}]}]},{"path":"/","description":"","operations":[{"httpMethod":"POST","nickname":"post","type":"","summary":"like","parameters":[{"paramType":"body","name":"body","description":"\"The like content\"","dataType":"Like","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"{Result:success}","responseModel":""},{"code":200,"message":"{Result:error}","responseModel":""}]}]},{"path":"/","description":"","operations":[{"httpMethod":"DELETE","nickname":"cancel","type":"","summary":"cancel the like","parameters":[{"paramType":"body","name":"body","description":"\"The like you want to cancel\"","dataType":"LikeDeleteInfo","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"{Result:success}","responseModel":""},{"code":200,"message":"{Result:error}","responseModel":""}]}]}],"models":{"LikeResult":{"id":"LikeResult","properties":{"Likes":{"type":"array","description":"","items":{"$ref":"\u0026{models Like}"},"format":""},"Result":{"type":"string","description":"","format":""}}}}},"/moment":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/moment","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/","description":"","operations":[{"httpMethod":"POST","nickname":"publish","type":"","summary":"publish moment","parameters":[{"paramType":"body","name":"body","description":"\"The moment content\"","dataType":"Moment","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"{Result:success}","responseModel":""},{"code":200,"message":"{Result:error}","responseModel":""}]}]},{"path":"/","description":"","operations":[{"httpMethod":"GET","nickname":"get","type":"","summary":"get one's moments by UserId and Timestamp","parameters":[{"paramType":"query","name":"UserId","description":"\"Your UserId\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"query","name":"Timestamp","description":"\"Last query time\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"controllers.MomentResult","responseModel":"MomentResult"},{"code":200,"message":"{\"Result\": \"error\",\"Moments\": null}","responseModel":""}]}]},{"path":"/","description":"","operations":[{"httpMethod":"DELETE","nickname":"delete","type":"","summary":"delete the moment","parameters":[{"paramType":"body","name":"body","description":"\"The moment you want to delete\"","dataType":"MomentDeleteInfo","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"{Result:success}","responseModel":""},{"code":200,"message":"{Result:error}","responseModel":""}]}]},{"path":"/pull","description":"","operations":[{"httpMethod":"GET","nickname":"pull","type":"","summary":"pull moments by UserId and Timestamp","parameters":[{"paramType":"query","name":"UserId","description":"\"Your UserId\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"query","name":"Timestamp","description":"\"Last query time\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"controllers.MomentResult","responseModel":"MomentResult"},{"code":200,"message":"{\"Result\": \"error\",\"Moments\": null}","responseModel":""}]}]},{"path":"/exist","description":"","operations":[{"httpMethod":"GET","nickname":"exist","type":"","summary":"find out the moment exist or not","parameters":[{"paramType":"query","name":"MomentId","description":"\"The moment you want to query\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"{Result:(not)exist}","responseModel":""},{"code":200,"message":"{Result:error}","responseModel":""}]}]}],"models":{"MomentResult":{"id":"MomentResult","properties":{"Moments":{"type":"array","description":"","items":{"$ref":"\u0026{models Moment}"},"format":""},"Result":{"type":"string","description":"","format":""}}}}},"/unshare":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/unshare","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/","description":"","operations":[{"httpMethod":"POST","nickname":"unshare","type":"","summary":"unshare with someone","parameters":[{"paramType":"body","name":"body","description":"\"FollowId unshare with FollowedId\"","dataType":"UserRelation","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"{Result:success}","responseModel":""},{"code":200,"message":"{Result:error}","responseModel":""}]}]},{"path":"/","description":"","operations":[{"httpMethod":"DELETE","nickname":"share","type":"","summary":"cancel unshare someone","parameters":[{"paramType":"body","name":"body","description":"\"FollowId cancel unshare FollowedId\"","dataType":"UserRelation","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"{Result:success}","responseModel":""},{"code":200,"message":"{Result:error}","responseModel":""}]}]}]}}`
    BasePath string= "/v1"
)

var rootapi swagger.ResourceListing
var apilist map[string]*swagger.ApiDeclaration

func init() {
	if beego.EnableDocs {
		err := json.Unmarshal([]byte(Rootinfo), &rootapi)
		if err != nil {
			beego.Error(err)
		}
		err = json.Unmarshal([]byte(Subapi), &apilist)
		if err != nil {
			beego.Error(err)
		}
		beego.GlobalDocApi["Root"] = rootapi
		for k, v := range apilist {
			for i, a := range v.Apis {
				a.Path = urlReplace(k + a.Path)
				v.Apis[i] = a
			}
			v.BasePath = BasePath
			beego.GlobalDocApi[strings.Trim(k, "/")] = v
		}
	}
}


func urlReplace(src string) string {
	pt := strings.Split(src, "/")
	for i, p := range pt {
		if len(p) > 0 {
			if p[0] == ':' {
				pt[i] = "{" + p[1:] + "}"
			} else if p[0] == '?' && p[1] == ':' {
				pt[i] = "{" + p[2:] + "}"
			}
		}
	}
	return strings.Join(pt, "/")
}
