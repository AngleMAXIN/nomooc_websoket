## 为竞赛参与者提供持久连接，已检验参与者是否中断考试

- ### WeSocket

- ### Mysql

- ### GO

  ------


### 具体实现

- 接收来自客户端发送的uid、contestID参数，创建WebSocket连接

```go

func wsHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var (
		wsConn    *websocket.Conn
		err       error
		uID       int
		contestID int
	)

	// 取出uID
	if uID, err = strconv.Atoi(p.ByName("uID")); err != nil {
		return
	}

	// 取出contestID
	if contestID, err = strconv.Atoi(p.ByName("contestID")); err != nil {
		return
	}

	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}

	_ = webconn.NewConnection(wsConn, uID, contestID)
}
```

- ###### 监听连接状态

  一旦连接终止，立即向KillChan里发送当前连接用户信息，通知修改当前用户信息

  ```go
  func (conn *CustomerConnection) ListenLoop() {
  	var err error
  	for {
  		if _, _, err = conn.wsConn.ReadMessage(); err != nil {
  			// 用户突然断开链接
  			if websocket.IsCloseError(err, websocket.CloseGoingAway) {
  				log.Printf("Connection Close: user %d, contest %d\n", conn.user.uID, conn.user.contestID)
  				killChan <- conn.user
  			}
  			conn.wsConn.Close()
  			return
  		}
  	}
  }
  ```

- 监听killchan中状态，一旦有用户进来，立即开启协程，通知数据库修改用户状态，chan 是buffer channel，防止之前的代码阻塞

  ```go
  func ListenkillSignal() {
  	killChan = make(chan userOne, 10)
  	for v := range killChan {
  		go db.MarkUserStatus(v.uID, v.contestID)
  	}
  }
  ```

- 事先初始化Prepare

  ```go
  
  const (
  	sqlStr = `UPDATE contest_partner SET is_disabled='1' WHERE user_id = ? and contest_id = ?;`
  )
  
  var (
  	dbConn *sql.DB
  	err    error
  	stmUp  *sql.Stmt
  )
  
  func init() {
      //此处应该getenv
  	if dbConn, err = sql.Open("mysql", "root:maxinz@tcp(localhost:3306)/oj_database?charset=utf8"); err != nil {
  		panic(err.Error())
  	}
  
  	if stmUp, err = dbConn.Prepare(sqlStr); err != nil {
  		log.Printf("Prepare update SQl Failed: %s\n", err)
  		return
  	}
  
  }
  ```

- 修改用户状态

  ```go
  // MarkUserStatus 用来更新用户的竞赛权限
  func MarkUserStatus(uID, contestID int) error {
  	var (
  		err error
  	)
  	if _, err = stmUp.Exec(uID, contestID); err != nil {
  		log.Printf("Exec update userid: %d and contest_id: %d Failed\n", uID, contestID)
  		return err
  	}
  	log.Printf("Update user %d, contest %d\n", uID, contestID)
  	return nil
  }
  
  ```


