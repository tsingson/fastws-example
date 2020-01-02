# wfastws-example

a example for websocket via
* [fasthttp](http://github.com/valyala/fasthttp) 
* [fasthttp router](https://github.com/fasthttp/router) 
*  [fastws ](https://github.com/fasthttp/fastws) -- -- Websocket implementation for fasthttp
*   [flatbuffers](https://github.com/google/flatbuffers) -- serialized / un-surialized for  go and javascript  
*   [websockets/ws](https://github.com/websockets/ws) -- javascript websocket client 

## 1. example IDL 
xone.fbs and JS example from [https://www.cnblogs.com/sevenstar/p/FlatBuffer.html](https://www.cnblogs.com/sevenstar/p/FlatBuffer.html), thanks

```
namespace xone.genflat;

  table LoginRequest{
  msgID:int=1;
  username:string;
  password:string;
  }

  table LoginResponse{
 msgID:int=2;
 uid:string;
 }

 //root_type非必须。

 //root_type LoginRequest;
 //root_type LoginRespons
```

## 2. flatc

generate javascript

```
flatc -s --gen-mutable ./*.fbs
```



generate golang

```
flatc  --go --gen-object-api --gen-all  --gen-compare  --raw-binary ./*.fbs
```



## 3. javascript

```
// ------------ ./jsclient/index.js

const flatbuffers = require('./flatbuffers').flatbuffers;
const xone = require('./xone_generated').xone; //Generated by `flatc`.

//-------------------------------------------
//  serialized
//-------------------------------------------
let b = new flatbuffers.Builder(1);
let username = b.createString("zlssssssssssssh");
let password = b.createString("xxxxxxxxxxxxxxxxxxx");
xone.genflat.LoginRequest.startLoginRequest(b);
xone.genflat.LoginRequest.addUsername(b, username);
xone.genflat.LoginRequest.addPassword(b, password);
xone.genflat.LoginRequest.addMsgID(b, 5);
let req = xone.genflat.LoginRequest.endLoginRequest(b);
b.finish(req); //创建结束时记得调用这个finish方法。


let uint8Array = b.asUint8Array();

console.log(uint8Array);
// console.log(b.dataBuffer() );
//-------------------------------------------
//  un-serialized
//-------------------------------------------
let bb = new flatbuffers.ByteBuffer(uint8Array);
let lgg = xone.genflat.LoginRequest.getRootAsLoginRequest(bb);


console.log("username: ", lgg.username());
console.log("password", lgg.password());
console.log("msgID: ", lgg.msgID());

```




## 4. golang 

```

// ------ ./apis/genflat/model.go

func (a *LoginRequestT) Byte() []byte {
	b := flatbuffers.NewBuilder(0)
	b.Finish(LoginRequestPack(b, a))
	return b.FinishedBytes()
}

func ByteLoginRequestT(b []byte) *LoginRequestT {
	return GetRootAsLoginRequest(b, 0).UnPack()
}


// ------- ./apis/genflat/model_test.go

func TestLoginRequestT_Byte(t *testing.T) {
	as := assert.New(t)
	// serialized
	l := &LoginRequestT{
		MsgID:    1,
		Username: "1",
		Password: "1",
	}

	b := l.Byte()

	// un-serialized 
	c := ByteLoginRequestT(b)
	if l.MsgID > 0 {
		fmt.Println(" id > ", c.MsgID, " u > ", c.Username, " pw > ", c.Password)
	}

	as.Equal(l.Password, c.Password)

}

```

## 5. Code example

```
./cmd/wsserver/main.go ----- websocket server 
./cmd/wsclient/main.go ----- websocket client
./ws/... -------------------  websocket go code for websocket handler and websocket client 
./jsclient/ws.js  ---------- javascript client code , please check-out package.json for depends
```



## 6. refence

*  [https://github.com/google/flatbuffers/issues/3781](https://github.com/google/flatbuffers/issues/3781)