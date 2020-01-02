
# flatbuffers 在 websocket 中互通的示例


## 0. 简要说明
为某个开源项目增加 websocket 对接, 写了这个示例

代码中 javascript 对 flatbuffers 的序列化/反序列化, 查了一天资料, 嗯哼, 最终完成了.
看代码吧.........


## 1. 使用代码库

示例代码使用了以下开源库
* [fasthttp](http://github.com/valyala/fasthttp) 
* [fasthttp router](https://github.com/fasthttp/router) 
*  [fastws ](https://github.com/fasthttp/fastws) ----  fasthttp 实现的 websocket 库
*   [flatbuffers](https://github.com/google/flatbuffers) ---- flatbuffers 高效反序列化通用库, 用在 go语言/javascript 
*   [websockets/ws](https://github.com/websockets/ws) ---- javascript websocket 通用库

## 1. flatbuffers  IDL 示例 
xone.fbs 示例来自 [https://www.cnblogs.com/sevenstar/p/FlatBuffer.html](https://www.cnblogs.com/sevenstar/p/FlatBuffer.html), 感谢!!

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

## 2. flatc 编译代码

生成 javascript

```
flatc -s --gen-mutable ./*.fbs
```



生成 golang

```
flatc  --go --gen-object-api --gen-all  --gen-compare  --raw-binary ./*.fbs
```

## 3. 主要代码说明

```
./cmd/wsserver/main.go ----- websocket server 
./cmd/wsclient/main.go ----- websocket client
./ws/... -------------------  websocket go code for websocket handler and websocket client 
./jsclient/ws.js  ---------- javascript client code , please check-out package.json for depends
```




## 4. javascript 序列化/反序列化

**请注意代码注释中的--------- 特别注意这一行**

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


let uint8Array = b.asUint8Array();   // ------------- 特别注意这一行

console.log(uint8Array);
// console.log(b.dataBuffer() );
//-------------------------------------------
//  un-serialized
//-------------------------------------------
let bb = new flatbuffers.ByteBuffer(uint8Array);  //-------------- 特别注意这一行
let lgg = xone.genflat.LoginRequest.getRootAsLoginRequest(bb);


console.log("username: ", lgg.username());
console.log("password", lgg.password());
console.log("msgID: ", lgg.msgID());

```



## 5.  golang 中对 flatbuffers 的序列化/反序列化

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

## 6. websocket 代码
```

ws.onmessage = (event) => {
    //-------------------------------------------------------------------
    //   read from websocket and un-serialized via flatbuffers
    //--------------------------------------------------------------------
    let aa = str2ab(event.data);
    let bb = new flatbuffers.ByteBuffer(aa);
    let lgg = xone.genflat.LoginRequest.getRootAsLoginRequest(bb);
    let pw = lgg.password();

    if (typeof pw === 'string') {
        console.log("----------------------------------------------");

        console.log("username: ", lgg.username());
        console.log("password", lgg.password());
        console.log("msgID: ", lgg.msgID());
    } else {
        console.log("=================================");
        console.log(event.data);
    }


    // console.log(`Roundtrip time: ${Date.now() }` , ab2str(d ));

    setTimeout(function timeout() {
    //-------------------------------------------------------------------
    //   serialized via flatbuffers and send to websocket 
    //--------------------------------------------------------------------
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

        ws.send(uint8Array);
    }, 500);
};

function str2ab(str) {
    let array = new Uint8Array(str.length);
    for (let i = 0; i < str.length; i++) {
        array[i] = str.charCodeAt(i);
    }
    return array
}

```



## 6. 参考

*  [https://github.com/google/flatbuffers/issues/3781](https://github.com/google/flatbuffers/issues/3781)

## 8. License
MIT

-----

code  by [tsingson](https://tsingson.github.io)

![tsingson-logo](README.assets/tsingson-logo.png)

