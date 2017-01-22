var d = false;
var loading = false;
var r;
var q = document.location.search;
var s = 0, c = 50;
if(q != "") {
	if(q.match(/s=(.*)&?/)) s = parseInt(q.match(/s=(.*)&?/)[1]);
	if(q.match(/c=(.*)&?/)) c = parseInt(q.match(/c=(.*)&?/)[1]);
}
function goto(b) {
	if(b) s += document.getElementsByClassName('box').length;
	else s -= c;
	var q = s <= 0 ? '' : '?s=' + s;
	if(q != '') q += c == 50 ? '' : '&c=' + c;
	else q = c == 50 ? '' : '?c=' + c;
	var l = document.location
		document.location = l.protocol + '//' + l.host + l.pathname + q;
}
function shImg() {
	var p = document.documentElement.scrollTop || document.body.scrollTop;
	var t = p - window.innerHeight * 0.5;
	var b = p + window.innerHeight * 1.5;
	if(t == 0 && b == 0) {
		var a = document.getElementsByClassName('hidden');
		for(i = 0, m = a.length; i < a; i++)
			a[i].src = a[i].alt;
		return
	}
	if(b > (document.body.scrollHeight || document.documentElement.scrollHeight))
		loadNextPage();
	var a = document.getElementsByClassName('hidden');
	var tmp = [];
	for(i = a.length-1; 0 <= i; i--)
		if(t < a[i].offsetParent.offsetTop && a[i].offsetParent.offsetTop < b) {
			tmp.push(a[i].getElementsByTagName('img')[0]);
			a[i].className = 'img';
		}
	for(i = 0, max = tmp.length; i < max; i++)
		tmp[i].setAttribute('src', tmp[i].getAttribute('alt'));
}
function rsImg(i) {
	i.setAttribute('style', 'max-height: 100%; max-width: 100%;');
}
function loadNextPage() {
	if(loading) return;
	loading = true;
	var t = s + document.getElementsByClassName('box').length;
	var q = t <= 0 ? '' : '?s=' + t;
	r = new XMLHttpRequest();
	r.open("GET", location.protocol + '//' + location.host + location.pathname + q, true);
	r.onreadystatechange=writeBody;
	r.send(null);
}
function writeBody() {
	if(r.readyState == 4 && r.status == 200) {
		var p = new DOMParser();
		var o = p.parseFromString(r.responseText, "text/html");
		o.body.removeChild(o.getElementById('control'));
		var a = o.getElementsByClassName('box');
		for(i=0,m=a.length; i<m; i++)
			document.body.appendChild(a[i].cloneNode(true));
	}
	if(r.status != 404)
		loading = false;
}
function rm() {
	document.body.style.overflow = 'auto';
	document.getElementById("lbox").style.display = 'none';
}
function lbox(o) {
	var w = document.getElementById("lbox");
	w.style.display = 'inline';
	var g = w.getElementsByTagName('img');
	for(i=0; i < g.length; i++)
		g[i].src = o.alt;
	document.body.style.overflow = 'hidden';
}
function stop(e) {
	e.stopPropagation();
}
function nextImg(e) {
	stop(e);
	var l = document.getElementsByClassName('limg');
	var i = document.getElementsByTagName('img');
	for(j = 1, m = i.length-1; j < m; j++)
		if(i[j].src == l[0].src) {
			l[0].src = i[j+1].alt;
			if(l[0].src != "")
				window.scrollTo(0,i[j+1].offsetParent.offsetParent.offsetTop);
			return;
		}
}
function prevImg(e) {
	stop(e);
	var l = document.getElementsByClassName('limg');
	var i = document.getElementsByTagName('img');
	for(j = i.length-2; 0 < j; j--)
		if(i[j].src == l[0].src) {
			l[0].src = i[j-1].alt;
			if(l[0].src != "")
				window.scrollTo(0,i[j-1].offsetParent.offsetParent.offsetTop);
			return;
		}
}
function toggleDirs() {
	var d = document.getElementById('dirs');
	if((d.style.display || 'none') == 'none')
		d.style.display = 'block';
	else
		d.style.display = 'none';
}
document.addEventListener('DOMContentLoaded', function() {
	document.addEventListener('scroll', shImg);
	document.documentElement.addEventListener('load', shImg);
}, false);
