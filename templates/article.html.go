// This file is automatically generated by qtc from "article.html".
// See https://github.com/valyala/quicktemplate for details.

//line article.html:1
package templates

//line article.html:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line article.html:1
import "fmt"

//line article.html:2
import "strconv"

//line article.html:3
import "github.com/bakape/meguca/common"

//line article.html:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line article.html:5
func streamrenderArticle(qw422016 *qt422016.Writer, p common.Post, op uint64, omit, imageOmit int, subject, root string) {
	//line article.html:6
	id := strconv.FormatUint(p.ID, 10)

	//line article.html:6
	qw422016.N().S(`<article id="p`)
	//line article.html:7
	qw422016.N().S(id)
	//line article.html:7
	qw422016.N().S(`" class="glass`)
	//line article.html:7
	if p.Editing {
		//line article.html:7
		qw422016.N().S(` `)
		//line article.html:7
		qw422016.N().S(`editing`)
		//line article.html:7
	}
	//line article.html:7
	qw422016.N().S(`"><header class="spaced">`)
	//line article.html:9
	if subject != "" {
		//line article.html:9
		qw422016.N().S(`<h3>「`)
		//line article.html:11
		qw422016.E().S(subject)
		//line article.html:11
		qw422016.N().S(`」</h3>`)
		//line article.html:13
	}
	//line article.html:13
	qw422016.N().S(`<b class="name"`)
	//line article.html:14
	if p.Auth != "" {
		//line article.html:14
		qw422016.N().S(` `)
		//line article.html:14
		qw422016.N().S(`class="admin"`)
		//line article.html:14
	}
	//line article.html:14
	qw422016.N().S(`>`)
	//line article.html:15
	if p.Name != "" || p.Trip == "" {
		//line article.html:16
		if p.Name != "" {
			//line article.html:17
			qw422016.E().S(p.Name)
			//line article.html:18
		} else {
			//line article.html:18
			qw422016.N().S(`Anonymous`)
			//line article.html:20
		}
		//line article.html:21
		if p.Trip != "" {
			//line article.html:22
			qw422016.N().S(` `)
			//line article.html:23
		}
		//line article.html:24
	}
	//line article.html:25
	if p.Trip != "" {
		//line article.html:25
		qw422016.N().S(`<code>!`)
		//line article.html:27
		qw422016.E().S(p.Trip)
		//line article.html:27
		qw422016.N().S(`</code>`)
		//line article.html:29
	}
	//line article.html:30
	if p.Auth != "" {
		//line article.html:30
		qw422016.N().S(`##`)
		//line article.html:31
		qw422016.N().S(` `)
		//line article.html:31
		qw422016.E().S(p.Auth)
		//line article.html:32
	}
	//line article.html:32
	qw422016.N().S(`</b><time>`)
	//line article.html:35
	qw422016.N().S(formatTime(p.Time))
	//line article.html:35
	qw422016.N().S(`</time><nav><a href="#p`)
	//line article.html:38
	qw422016.N().S(id)
	//line article.html:38
	qw422016.N().S(`">No.</a><a class="quote" href="#p`)
	//line article.html:41
	qw422016.N().S(id)
	//line article.html:41
	qw422016.N().S(`">`)
	//line article.html:42
	qw422016.N().S(id)
	//line article.html:42
	qw422016.N().S(`</a></nav><a class="control"><svg xmlns="http://www.w3.org/2000/svg" width="8" height="8" viewBox="0 0 8 8"><path d="M1.5 0l-1.5 1.5 4 4 4-4-1.5-1.5-2.5 2.5-2.5-2.5z" transform="translate(0 1)" /></svg></a></header>`)
	//line article.html:51
	var src string

	//line article.html:52
	if p.Image != nil {
		//line article.html:53
		img := *p.Image

		//line article.html:54
		src = sourcePath(img.FileType, img.SHA1)

		//line article.html:55
		ISSrc := root + src

		//line article.html:55
		qw422016.N().S(`<figcaption class="spaced"><a class="image-toggle act" hidden></a><span class="spaced image-search-container"><a class="image-search google" target="_blank" rel="nofollow" href="https://www.google.com/searchbyimage?image_url=`)
		//line article.html:59
		qw422016.N().S(ISSrc)
		//line article.html:59
		qw422016.N().S(`">G</a><a class="image-search iqdb" target="_blank" rel="nofollow" href="http://iqdb.org/?url=`)
		//line article.html:62
		qw422016.N().S(ISSrc)
		//line article.html:62
		qw422016.N().S(`">Iq</a><a class="image-search saucenao" target="_blank" rel="nofollow" href="http://saucenao.com/search.php?db=999&url=`)
		//line article.html:65
		qw422016.N().S(ISSrc)
		//line article.html:65
		qw422016.N().S(`">Sn</a><a class="image-search desustorage" target="_blank" rel="nofollow" href="https://desuarchive.org/_/search/image/`)
		//line article.html:68
		qw422016.N().S(img.MD5)
		//line article.html:68
		qw422016.N().S(`">Ds</a><a class="image-search exhentai" target="_blank" rel="nofollow" href="http://exhentai.org/?fs_similar=1&fs_exp=1&f_shash=`)
		//line article.html:71
		qw422016.N().S(img.SHA1)
		//line article.html:71
		qw422016.N().S(`">Ex</a></span><span>(`)
		//line article.html:77
		if img.Audio {
			//line article.html:77
			qw422016.N().S(`♫,`)
			//line article.html:78
			qw422016.N().S(` `)
			//line article.html:79
		}
		//line article.html:80
		if img.Length != 0 {
			//line article.html:81
			l := img.Length

			//line article.html:82
			if l < 60 {
				//line article.html:83
				qw422016.N().S(fmt.Sprintf("0:%02d", l))
				//line article.html:84
			} else {
				//line article.html:85
				min := l / 6

				//line article.html:86
				qw422016.N().S(fmt.Sprintf("%02d:%02d", min, l-min))
				//line article.html:87
			}
			//line article.html:87
			qw422016.N().S(`,`)
			//line article.html:88
			qw422016.N().S(` `)
			//line article.html:89
		}
		//line article.html:90
		if img.APNG {
			//line article.html:90
			qw422016.N().S(`APNG,`)
			//line article.html:91
			qw422016.N().S(` `)
			//line article.html:92
		}
		//line article.html:93
		qw422016.N().S(readableFileSize(img.Size))
		//line article.html:93
		qw422016.N().S(`,`)
		//line article.html:93
		qw422016.N().S(` `)
		//line article.html:94
		qw422016.N().S(strconv.FormatUint(uint64(img.Dims[0]), 10))
		//line article.html:94
		qw422016.N().S(`x`)
		//line article.html:96
		qw422016.N().S(strconv.FormatUint(uint64(img.Dims[1]), 10))
		//line article.html:96
		qw422016.N().S(`)</span>`)
		//line article.html:99
		name := imageName(img.FileType, img.Name)

		//line article.html:99
		qw422016.N().S(`<a href="`)
		//line article.html:100
		qw422016.N().S(src)
		//line article.html:100
		qw422016.N().S(`" download="`)
		//line article.html:100
		qw422016.N().S(name)
		//line article.html:100
		qw422016.N().S(`">`)
		//line article.html:101
		qw422016.N().S(name)
		//line article.html:101
		qw422016.N().S(`</a></figcaption>`)
		//line article.html:104
	}
	//line article.html:104
	qw422016.N().S(`<div class="post-container">`)
	//line article.html:106
	if p.Image != nil {
		//line article.html:107
		img := *p.Image

		//line article.html:107
		qw422016.N().S(`<figure><a target="_blank" href="`)
		//line article.html:109
		qw422016.N().S(src)
		//line article.html:109
		qw422016.N().S(`">`)
		//line article.html:110
		if img.Spoiler {
			//line article.html:110
			qw422016.N().S(`<!-- TODO: board-specific server-side spoiler rendering --><img src="/assets/spoil/default.jpg" width="125" height="125">`)
			//line article.html:113
		} else {
			//line article.html:114
			w, h := correctDims(subject != "", img.Dims[2], img.Dims[3])

			//line article.html:114
			qw422016.N().S(`<img src="`)
			//line article.html:115
			qw422016.N().S(thumbPath(img.FileType, img.SHA1))
			//line article.html:115
			qw422016.N().S(`" width="`)
			//line article.html:115
			qw422016.N().S(w)
			//line article.html:115
			qw422016.N().S(`" height="`)
			//line article.html:115
			qw422016.N().S(h)
			//line article.html:115
			qw422016.N().S(`">`)
			//line article.html:116
		}
		//line article.html:116
		qw422016.N().S(`</a></figure>`)
		//line article.html:119
	}
	//line article.html:119
	qw422016.N().S(`<blockquote>`)
	//line article.html:121
	streambody(qw422016, p, op)
	//line article.html:121
	qw422016.N().S(`</blockquote></div>`)
	//line article.html:124
	if omit != 0 {
		//line article.html:124
		qw422016.N().S(`<span class="omit" data-omit="`)
		//line article.html:125
		qw422016.N().D(omit)
		//line article.html:125
		qw422016.N().S(`" data-image-omit="`)
		//line article.html:125
		qw422016.N().D(imageOmit)
		//line article.html:125
		qw422016.N().S(`">`)
		//line article.html:126
		qw422016.N().D(omit)
		//line article.html:126
		qw422016.N().S(` `)
		//line article.html:126
		qw422016.N().S(`post`)
		//line article.html:126
		if omit > 1 {
			//line article.html:126
			qw422016.N().S(`s`)
			//line article.html:126
		}
		//line article.html:127
		qw422016.N().S(` `)
		//line article.html:127
		qw422016.N().S(`and`)
		//line article.html:127
		qw422016.N().S(` `)
		//line article.html:127
		qw422016.N().D(imageOmit)
		//line article.html:128
		qw422016.N().S(` `)
		//line article.html:128
		qw422016.N().S(`image`)
		//line article.html:128
		if imageOmit > 1 {
			//line article.html:128
			qw422016.N().S(`s`)
			//line article.html:128
		}
		//line article.html:128
		qw422016.N().S(`omitted`)
		//line article.html:129
		qw422016.N().S(` `)
		//line article.html:129
		qw422016.N().S(`<span class="act"><a href="`)
		//line article.html:131
		qw422016.N().S(strconv.FormatUint(op, 10))
		//line article.html:131
		qw422016.N().S(`" class="history">See All</a></span></span>`)
		//line article.html:136
	}
	//line article.html:137
	if p.Backlinks != nil {
		//line article.html:137
		qw422016.N().S(`<small class="spaced">`)
		//line article.html:139
		for id, link := range p.Backlinks {
			//line article.html:140
			streampostLink(qw422016, id, link.OP, link.Board, link.OP != op)
			//line article.html:141
		}
		//line article.html:141
		qw422016.N().S(`</small>`)
		//line article.html:143
	}
	//line article.html:143
	qw422016.N().S(`</article>`)
//line article.html:145
}

//line article.html:145
func writerenderArticle(qq422016 qtio422016.Writer, p common.Post, op uint64, omit, imageOmit int, subject, root string) {
	//line article.html:145
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line article.html:145
	streamrenderArticle(qw422016, p, op, omit, imageOmit, subject, root)
	//line article.html:145
	qt422016.ReleaseWriter(qw422016)
//line article.html:145
}

//line article.html:145
func renderArticle(p common.Post, op uint64, omit, imageOmit int, subject, root string) string {
	//line article.html:145
	qb422016 := qt422016.AcquireByteBuffer()
	//line article.html:145
	writerenderArticle(qb422016, p, op, omit, imageOmit, subject, root)
	//line article.html:145
	qs422016 := string(qb422016.B)
	//line article.html:145
	qt422016.ReleaseByteBuffer(qb422016)
	//line article.html:145
	return qs422016
//line article.html:145
}