package mediator

import (
	"fmt"
	"strings"
)

type CDDriver struct {
	Data string
}

func (c *CDDriver) ReadData() {
	c.Data = "music,image"

	fmt.Printf("Ⅰ ：CD 先获取数据，然后通知中介者\n")
	GetMediatorInstance().changed(c)
}

type CPU struct {
	Video string
	Sound string
}

func (c *CPU) Process(data string) {
	sp := strings.Split(data, ",")
	c.Sound = sp[0]
	c.Video = sp[1]

	fmt.Printf("Ⅲ ：CUP 对数据进行处理，并接着通知中介者\n")
	GetMediatorInstance().changed(c)
}

type VideoCard struct {
	Data string
}

func (v *VideoCard) Display(data string) {
	v.Data = data
	fmt.Printf("Ⅵ ：收到中介者传来的数据，播放画面\n")
	GetMediatorInstance().changed(v)
}

type SoundCard struct {
	Data string
}

func (s *SoundCard) Play(data string) {
	s.Data = data
	fmt.Printf("Ⅴ ：收到中介者传来的数据，播放声音\n")
	GetMediatorInstance().changed(s)
}

type Mediator struct {
	CD    *CDDriver
	CPU   *CPU
	Video *VideoCard
	Sound *SoundCard
}

var mediator *Mediator

func GetMediatorInstance() *Mediator {
	if mediator == nil {
		mediator = &Mediator{}
	}
	return mediator
}

func (m *Mediator) changed(i interface{}) {
	switch inst := i.(type) {
	case *CDDriver:
		fmt.Println("Ⅱ ：中介者收到 CD 消息，通知 CPU 进行处理")
		m.CPU.Process(inst.Data)
	case *CPU:
		fmt.Println("Ⅳ ：中介者收到 CPU 的消息，分别通知 Sound 和 Video 对象播放视频")
		m.Sound.Play(inst.Sound)
		m.Video.Display(inst.Video)
	}
}

func main() {
	mediator := GetMediatorInstance()
	mediator.CD = &CDDriver{}
	mediator.CPU = &CPU{}
	mediator.Video = &VideoCard{}
	mediator.Sound = &SoundCard{}

	//Tiggle
	mediator.CD.ReadData()
}

//output:
//Ⅰ ：CD 先获取数据，然后通知中介者
//Ⅱ ：中介者收到 CD 消息，通知 CPU 进行处理
//Ⅲ ：CUP 对数据进行处理，并接着通知中介者
//Ⅳ ：中介者收到 CPU 的消息，分别通知 Sound 和 Video 对象播放视频
//Ⅴ ：收到中介者传来的数据，播放声音
//Ⅵ ：收到中介者传来的数据，播放画面
