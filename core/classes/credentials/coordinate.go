package credentials

// Coordinate 表示一个六维度的坐标, 它是生成密码的6个关键参数
type Coordinate struct {
	Email    string // prop: user.email     ; 凭证持有者的邮箱地址
	Domain1  string // prop: manager.domain ; 云公钥所对应的域
	Domain2  string // prop: pass.domain    ; 凭证的工作域名
	Username string // prop: pass.name      ; 凭证的用户名
	Scene    string // prop: pass.scene     ; 凭证的使用场景
	Revision int    // prop: pass.revision  ; 凭证密码的版本
}
