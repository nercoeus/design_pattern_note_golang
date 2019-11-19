package template_method

import "fmt"

type Downloader interface {
	Download(uri string)
}

// 算法模板对象，包含一个 implement 接口对象（实现具体算法的子类）
type template struct {
	implement
	uri string
}

// 算法中步骤分为 download 和 save
type implement interface {
	download()
	save()
}

// 返回一个模板方法对象，具体实现子类是 impl
func newTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}

// 整个算法执行流程，分为几个步骤
func (t *template) Download(uri string) {
	t.uri = uri
	fmt.Print("prepare downloading\n")
	t.implement.download()
	t.implement.save()
	fmt.Print("finish downloading\n")
}

// 提供一个默认的 save 方法
func (t *template) save() {
	fmt.Print("default save\n")
}

// 具体实现算法步骤子类
type HTTPDownloader struct {
	*template
}

// http 子类的工厂函数
func NewHTTPDownloader() Downloader {
	downloader := &HTTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

// 子类实现具体的算法步骤
func (d *HTTPDownloader) download() {
	fmt.Printf("download %s via http\n", d.uri)
}

// http 子类实现了 save 方法
func (*HTTPDownloader) save() {
	fmt.Printf("http save\n")
}

// ftp 子类
type FTPDownloader struct {
	*template
}

// ftp 子类的工厂函数
func NewFTPDownloader() Downloader {
	downloader := &FTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

// Ftp 仅仅实现了 doenload 方法
func (d *FTPDownloader) download() {
	fmt.Printf("download %s via ftp\n", d.uri)
}

// 分别使用 HTTP 和 FTP 来下载文件
func main() {
	var downloader1 Downloader = NewHTTPDownloader()
	downloader1.Download("http://example.com/abc.zip")

	var downloader2 Downloader = NewFTPDownloader()
	downloader2.Download("ftp://example.com/abc.zip")
}

//output:
//prepare downloading
//download http://example.com/abc.zip via http
//http save
//finish downloading
//prepare downloading
//download ftp://example.com/abc.zip via ftp
//default save
//finish downloading
