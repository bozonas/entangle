(this.webpackJsonpclient=this.webpackJsonpclient||[]).push([[0],{300:function(e,t,a){e.exports=a.p+"static/media/logo.5d5d9eef.svg"},350:function(e,t,a){e.exports=a(579)},355:function(e,t,a){},358:function(e,t,a){},579:function(e,t,a){"use strict";a.r(t);var n=a(0),r=a.n(n),c=a(82),o=a.n(c),l=(a(355),a(356),a(357),a(333)),i=a(71),u=(a(358),a(621)),s=a(622),m=a(300),p=a.n(m),d={h1:{marginTop:"3em"}},h=function(e){var t=e.children;return r.a.createElement(r.a.Fragment,null,r.a.createElement(u.a,{as:"h1",content:"Entangle",style:d.h1,textAlign:"center"}),r.a.createElement("img",{src:p.a,className:"App-logo",alt:"logo"}),r.a.createElement(s.a,null,t))},f=a(24),E=a.n(f),g=a(42),v=a(37),b=a(324),w=a(321),y=a(585),j=a(623),O=a(95),x=a(580),S=a(624),k=a(217),A=a(335),C=a(123),F=a.n(C),N=a(177),T=a.n(N),V="http://sleepy-earth-62648.herokuapp.com/api";function L(e){for(var t="ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",a="",n=0;n<e;n++)a+=t.charAt(Math.floor(Math.random()*t.length));return a}var P=function(e){Object(n.useEffect)((function(){C({name:"secret"},{required:!0})}),[]);var t=Object(n.useState)(""),a=Object(v.a)(t,2),c=a[0],o=a[1],l=Object(n.useState)(!1),i=Object(v.a)(l,2),u=i[0],s=i[1],m=Object(n.useState)({form:!0,copy:!1}),p=Object(v.a)(m,2),d=p[0],h=p[1],f=Object(A.a)(),C=f.register,N=f.errors,P=f.handleSubmit,R=f.setValue,U=f.triggerValidation,B=function(){var t=Object(g.a)(E.a.mark((function t(a){var n,r,c,l,i;return E.a.wrap((function(t){for(;;)switch(t.prev=t.next){case 0:return s(!0),n=L(10),r=L(20),c=F.a.AES.encrypt(a.secret,r).toString(),l={key:n,ciphertext:c},t.prev=5,t.next=8,T.a.post("".concat(V,"/message"),l);case 8:i=t.sent,console.log(i),o("".concat(window.location.href).concat(n,"-").concat(r)),h({form:!1,copy:!1}),setTimeout((function(){h({form:!1,copy:!0})}),500),t.next=18;break;case 15:t.prev=15,t.t0=t.catch(5),e.history.push("/errorPage");case 18:s(!1);case 19:case"end":return t.stop()}}),t,null,[[5,15]])})));return function(e){return t.apply(this,arguments)}}();return r.a.createElement(r.a.Fragment,null,r.a.createElement(b.a,null,r.a.createElement(b.a.Row,{centered:!0},r.a.createElement(b.a.Column,{width:8,textAlign:"center"},r.a.createElement(w.a,{visible:d.copy,animation:"fade up",duration:500,de:!0},r.a.createElement(r.a.Fragment,null,r.a.createElement(y.a,{fluid:!0,defaultValue:c,action:{color:"teal",labelPosition:"right",icon:"copy",content:"Copy",onClick:function(){navigator.clipboard.writeText(c),Object(k.toast)({type:"success",icon:"copy",title:"Copied",description:"URL has been copied to your clipboard",animation:"fade up",time:3e3})}}}))),r.a.createElement(w.a,{visible:d.form,animation:"fade up",duration:500},r.a.createElement(j.a,{onSubmit:P(B)},r.a.createElement(j.a.TextArea,{rows:5,name:"secret",placeholder:"Enter your secret",onChange:function(){var e=Object(g.a)(E.a.mark((function e(t,a){var n,r;return E.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return n=a.name,r=a.value,R(n,r),e.next=4,U(n);case 4:case"end":return e.stop()}}),e)})));return function(t,a){return e.apply(this,arguments)}}(),error:!!N.secret}),r.a.createElement(O.a,null,"Save secret")))))),r.a.createElement(k.SemanticToastContainer,{position:"bottom-right"}),r.a.createElement(x.a,{active:u,page:!0},r.a.createElement(S.a,{inverted:!0},"Loading")))},R=a(584),U=function(e){var t=Object(n.useState)(""),a=Object(v.a)(t,2),c=a[0],o=a[1],l=Object(n.useState)(!0),i=Object(v.a)(l,2),u=i[0],s=i[1];return Object(n.useEffect)((function(){s(!0);var t=e.match.params.id.split("-"),a=Object(v.a)(t,2),n=a[0],r=a[1];T()("".concat(V,"/message/").concat(n)).then((function(e){var t=e.data.ciphertext,a=F.a.AES.decrypt(t,r).toString(F.a.enc.Utf8);o(a),s(!1)})).catch((function(t){e.history.push("/notFound")}))}),[]),r.a.createElement(r.a.Fragment,null,r.a.createElement(b.a,null,r.a.createElement(b.a.Row,{centered:!0},r.a.createElement(b.a.Column,{width:8,textAlign:"center"},r.a.createElement(j.a,null,r.a.createElement(R.a,{rows:5,name:"secret",defaultValue:c}))))),r.a.createElement(x.a,{active:u,page:!0},r.a.createElement(S.a,{inverted:!0},"Loading")))},B=function(){return r.a.createElement("div",{className:"ui statistic"},r.a.createElement("div",{className:"value"},"404"),r.a.createElement("div",{className:"label"},"Error"))},J=function(){return r.a.createElement("h1",null,"Upsy.... something went horribly wrong")},M=function(){return r.a.createElement(l.a,null,r.a.createElement("div",{className:"App"},r.a.createElement(h,null,r.a.createElement(i.c,null,r.a.createElement(i.a,{exact:!0,path:"/",component:P}),r.a.createElement(i.a,{path:"/notFound",component:B}),r.a.createElement(i.a,{path:"/errorPage",component:J}),r.a.createElement(i.a,{path:"/:id",component:U})))))};Boolean("localhost"===window.location.hostname||"[::1]"===window.location.hostname||window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/));o.a.render(r.a.createElement(M,null),document.getElementById("root")),"serviceWorker"in navigator&&navigator.serviceWorker.ready.then((function(e){e.unregister()}))}},[[350,1,2]]]);
//# sourceMappingURL=main.4106639b.chunk.js.map