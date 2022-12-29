package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/longda/markov"
	"golang.org/x/net/websocket"
)

const listenAddres = "localhost:8000"

type socket struct {
	io.Reader
	io.Writer
	done chan bool
}

func (s socket) Close() error {
	s.done <- true
	return nil
}

var chain = markov.NewChain(2)

func main() {
	http.HandleFunc("/", rootHandler)
	http.Handle("/socket", websocket.Handler(socketHandler))
	err := http.ListenAndServe(listenAddres, nil)
	if err != nil {
		log.Fatal(err)
	}
}

var rootTemplate = template.Must(template.New("root").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8"/>
<script>

 var msgInput, sendBtn, rspOut, websocket

 function onMessage(event) {
	console.debug("WebSocket message received:", event);
 }
 function onClose(event) {
		console.log("The connection has been closed successfully.");
 }


	function initBindings() {
		msgInput = document.getElementById("chat-input")
		sendBtn = document.getElementById("send-chat")
		rspOut = document.getElementById("response-wrapper")

		sendBtn.onclick = sendBtnClickHandler
		msgInput.onkeydown = function( event) { if (event.keyCode == 13) sendBtn.click();}
		initWebSocket()
	}

	function initWebSocket() {
		websocket = new WebSocket("ws://{{.}}/socket");
	  websocket.onopen = function(evt) { onOpen(evt)}
	  websocket.onmerror = function(evt) { onError(evt)}
	  websocket.onmessage = function(evt) { onMessage(evt)}
	  websocket.close = function(evt) { onClose(evt)};
	}

	function onMessage(evnt) {
		writeToScreen('<span style="color: blue;">RESPONSE: ' + evnt.data+'</span>');
	}

	function onError(evnt) {
		writeToScreen('<span style="color: blue;">Error:</span>' + evnt.data);
	}
	function onClose(evnt) {
		writeToScreen("DISCONNECTED");
	}
	function onOpen(evnt) {
		writeToScreen("CONNECTED");
	}

	function sendBtnClickHandler(event) {
		doSend(msgInput.value)
		msgInput.value =""
	}

	function doSend(message) {
		writeToScreen("SENT:" + message)
		websocket.send(message)
	}


	function writeToScreen(message) {
		var pre = document.createElement("p")
		pre.style.wordWrap = "break-word"
		pre.innerHTML = message
		rspOut.appendChild(pre)
	}

	window.addEventListener("load", initBindings, false)
</script>
</head>
<body>
 <div>
  <div>
	<label for="chat-input">Message:</label>

 	<input type="text" id="chat-input"/> <input  value="Send" type="button" id="send-chat"/>
	</div>
	<div id="response-wrapper">

	</div>
 </div>
</body>
</html>
`))

func rootHandler(w http.ResponseWriter, r *http.Request) {
	rootTemplate.Execute(w, listenAddres)
}

func socketHandler(ws *websocket.Conn) {
	r, w := io.Pipe()
	go func() {
		_, err := io.Copy(io.MultiWriter(w, chain), ws)
		w.CloseWithError(err)
	}()
	sockConn := socket{r, ws, make(chan bool)}
	go match(sockConn)
	<-sockConn.done
}

var partner = make(chan io.ReadWriteCloser)

func match(c io.ReadWriteCloser) {
	fmt.Println(c, "Waiting for a partner...")
	select {
	case partner <- c:
		fmt.Println("I am executing bRo!")
	case p := <-partner:
		fmt.Println("I am chatting bRo!")
		chat(p, c)
	case <-time.After(5 * time.Second):
		fmt.Println("I am gonna chat with Bot bRo!")
		chat(Bot(), c)
	}
}

func chat(a, b io.ReadWriteCloser) {
	fmt.Println(a, "Found one! Say hi.")
	fmt.Println(b, "Found one! Say hi.")
	errc := make(chan error, 1)
	go copy(a, b, errc)
	go copy(b, a, errc)
	if err := <-errc; err != nil {
		log.Println(err)
	}
	a.Close()
	b.Close()
}

func copy(w io.Writer, r io.Reader, errc chan<- error) {
	_, err := io.Copy(w, r)
	errc <- err
}

// bot

func Bot() io.ReadWriteCloser {
	r, out := io.Pipe()
	return bot{r, out}
}

type bot struct {
	io.ReadCloser
	out io.Writer
}

func (b bot) Write(buf []byte) (int, error) {
	go b.speak()
	return len(buf), nil
}

func (b bot) speak() {
	time.Sleep(time.Second)
	msg := chain.Generate(10) // at most 10 random words
	b.out.Write([]byte(msg))
}
