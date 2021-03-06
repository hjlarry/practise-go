package handler

import "net/http"

func IsTokenValid(token string) bool {
	if len(token) != 40 {
		return false
	}
	// TODO: 判断token的时效性，是否过期
	// TODO: 从数据库表tbl_user_token查询username对应的token信息
	// TODO: 对比两个token是否一致
	return true
}

func HttpInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			_ = r.ParseForm()
			username := r.Form.Get("username")
			token := r.Form.Get("token")

			//验证登录token是否有效
			if len(username) < 3 || !IsTokenValid(token) {
				// w.WriteHeader(http.StatusForbidden)
				// token校验失败则跳转到登录页面
				http.Redirect(w, r, "/static/view/signin.html", http.StatusFound)
				return
			}
			h(w, r)

		})
}
