package main

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
	"github.com/kataras/iris/middleware/logger"
	recover2 "github.com/kataras/iris/middleware/recover"
	"mime/multipart"
	"strings"
	"time"
)

const maxSize = 5 << 20

func main() {
	fmt.Println("直接开启Iris的Go的web编程")
	app := iris.New()

	app.Logger().SetLevel("debug")

	app.Use(recover2.New())
	app.Use(logger.New())

	app.PartyFunc("/api/system", func(system router.Party) {
		system.Get("/information", systemInformationHandler)
		system.Get("/name", systemNameHandler)
	})

	app.PartyFunc("/api/user", func(users router.Party) {
		users.Get("/information", userInformationHandler)
		users.Get("/{userId:int}/status", userStatusHandler)
		users.Get("/{userId:int}/birthday", userBirthdayHandler)
	})

	app.PartyFunc("/api/video", func(video router.Party) {
		video.Post("/upload", iris.LimitRequestBodySize(maxSize), uploadVideoHandler)
	})

	app.PartyFunc("/api/rattrap/algorithm", func(algorithm router.Party) {
		algorithm.Get("", algorithmPageDataListHandler)
	})

	app.PartyFunc("/api/rattrap/chinese-ancient-article", func(chineseAncientPoems router.Party) {
		chineseAncientPoems.Get("/poems", chineseAncientPoemsListHandler)
		chineseAncientPoems.Get("/types",chineseAncientArticleTypeListHandler)
		chineseAncientPoems.Get("/{articleType:int}/articles",ancientArticleListHandler)
	})

	_ = app.Run(iris.Addr(":12309"), iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}), iris.WithoutServerError(iris.ErrServerClosed))
}

func systemInformationHandler(ctx iris.Context) {
	ctx.JSON(iris.Map{"name": "Iris-Web-系统"})
}

func systemNameHandler(ctx iris.Context) {
	ctx.JSON(iris.Map{"systemName": "Go-Iris系统"})
}

func userInformationHandler(ctx iris.Context) {
	ctx.JSON(iris.Map{"userName": "rootAdmin", "age": 20, "birthday": "1990-10-07"})
}

func userStatusHandler(ctx iris.Context) {
	userId := ctx.Params().Get("userId")
	ctx.JSON(iris.Map{"status": "OK", "userId": userId})
}

func userBirthdayHandler(ctx iris.Context) {
	userId := ctx.Params().Get("userId")
	ctx.JSON(iris.Map{"userId": userId, "birthday": "1990-10-06"})
}

func uploadVideoHandler(ctx iris.Context) {
	ctx.UploadFormFiles("F:/", beforeSave)
}

func beforeSave(ctx iris.Context, file *multipart.FileHeader) {
	ip := ctx.RemoteAddr()
	ip = strings.Replace(ip, ".", "_", -1)
	ip = strings.Replace(ip, ":", "_", -1)

	file.Filename = ip + "-" + file.Filename
}

func goLanguageFormat() {
	fmt.Println("This isa f嗲吗格式化的实验工作")
}

type JsonTime time.Time

func (this JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

func algorithmPageDataListHandler(ctx iris.Context) {
	results := []AlgorithmVo{{1, "冒泡排序", JsonTime(time.Now()), "smart mouse"}, {2, "插入排序", JsonTime(time.Now()), "Mouse"}, {3, "归并排序", JsonTime(time.Now()), "liuxu"}}
	bytes, _ := json.Marshal(results)
	ctx.JSON(iris.Map{"data": string(bytes), "message": "请求成功", "code": 0})
}

type AlgorithmVo struct {
	Id         int64    `json:"id"`
	Title      string   `json:"title"`
	CreateTime JsonTime `json:"createTime"`
	CreateBy   string   `json:"createBy"`
}

func chineseAncientPoemsListHandler(ctx iris.Context) {
	results := []ChineseAncientPoemVo{{1, "静夜思", "李白", []string{"白日依山尽", "黄河入海流", "欲穷千里目", "更上一层楼"}},
		{2, "锦瑟", "李商隐", []string{"锦瑟无端五十弦，一弦一柱思华年。", "庄生晓梦迷蝴蝶，望帝春心托杜鹃.", "沧海月明珠有泪，蓝田日暖玉生烟", "此情可待成追忆，只是当时已惘然"}},
		{3, "蜀相", "杜甫", []string{"丞相祠堂何处寻，锦官城外柏森森", "映阶碧草自春色，隔叶黄鹂空好音", "三顾频烦天下计，两朝开济老臣心", "出师未捷身先死，长使英雄泪满襟"}}}

	bytes, _ := json.Marshal(results)
	ctx.JSON(iris.Map{"data": string(bytes), "message": "请求成功", "code": 0})
}

type ChineseAncientPoemVo struct {
	Id           int64    `json:"id"`
	PoemName     string   `json:"poemName"`
	Author       string   `json:"author"`
	LineContents []string `json:"lineContents"`
}

func chineseAncientArticleTypeListHandler(ctx iris.Context){
	results := []ChineseAncientArticleTypeVo{
		{1,"诗经","诗经就是诗歌中的圣经"},
		{2,"楚辞","楚国的地盘"},
		{3,"汉赋","我大汉天朝的赋多么的气魄"},
		{4,"唐诗","封建社会的巅峰"},
		{5,"宋词","精神丰富的时代"},
		{6,"元曲","摇滚时代的来临"},
		{7,"明清小说","长篇集锦"},
	}

	bytes, _ := json.Marshal(results)
	ctx.JSON(iris.Map{"data":string(bytes),"message":"请求成功","code":0})
}

type ChineseAncientArticleTypeVo struct{
	Id int32 `json:"id"`
	TypeName string `json:"typeName"`
	Details string `json:"details"`
}

type AncientArticleVo struct{
	Id int64 `json:"id"`
	ArticleName string `json:"articleName"`
	ArticleType int32 `json:"articleType"`
}

func ancientArticleListHandler(ctx iris.Context){
	articleType := ctx.Params().Get("articleType")
	print("获取到的类型为:",articleType)
	results := []AncientArticleVo{
		{1,"关关雎鸠",1},
		{2,"南风",1},
		{3,"秦风",1},
		{4,"静夜思",4},
		{5,"出塞",4},
		{6,"悯农",4},
	}

	bytes,_  := json.Marshal(results)
	ctx.JSON(iris.Map{"data":string(bytes),"message":"请求成功","code":0})
}
