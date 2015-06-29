package controllers
import "github.com/revel/revel"

func init() {
    revel.OnAppStart(InitDB) // DBやテーブルの作成
    revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
    revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
    revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)

    //revel.InterceptMethod(Application.AddUser, revel.BEFORE)
    //revel.InterceptMethod(Hotels.checkUser, revel.BEFORE)
}
