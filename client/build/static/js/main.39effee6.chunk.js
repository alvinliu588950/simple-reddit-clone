(window.webpackJsonp=window.webpackJsonp||[]).push([[0],{18:function(e,t,a){e.exports=a(47)},24:function(e,t,a){},25:function(e,t,a){e.exports=a.p+"static/media/logo.5d5d9eef.svg"},26:function(e,t,a){},47:function(e,t,a){"use strict";a.r(t);var n=a(0),c=a.n(n),o=a(16),s=a.n(o),i=(a(24),a(3)),l=a(4),r=a(6),m=a(5),u=a(7),d=(a(25),a(26),a(27),a(17)),p=a(2),h=a(9),f=a.n(h),v=function(e){function t(e){var a;return Object(i.a)(this,t),(a=Object(r.a)(this,Object(m.a)(t).call(this,e))).state={topics:[],content:""},a.handleInputChange=a.handleInputChange.bind(Object(p.a)(Object(p.a)(a))),a.addTopics=a.addTopics.bind(Object(p.a)(Object(p.a)(a))),a}return Object(u.a)(t,e),Object(l.a)(t,[{key:"handleInputChange",value:function(e){var t=e.target,a=t.value,n=t.name;this.setState(Object(d.a)({},n,a))}},{key:"componentDidMount",value:function(){this.fetchTopics()}},{key:"render",value:function(){var e={marginBottom:"10px"};return c.a.createElement("div",null,c.a.createElement("div",{className:"columns"},c.a.createElement("div",{className:"column is-6 is-offset-3"},c.a.createElement("h1",{className:"title has-text-centered"},"Topics"))),c.a.createElement("div",{className:"columns"},c.a.createElement("div",{className:"column is-6 is-offset-3"},c.a.createElement("form",null,c.a.createElement("div",{className:"field"},c.a.createElement("label",{className:"label"},"New Topic"),c.a.createElement("div",{className:"control"},c.a.createElement("input",{className:"input",type:"text",placeholder:"Text input",name:"content",checked:this.state.content,onChange:this.handleInputChange}))),c.a.createElement("div",{className:"field"},c.a.createElement("div",{className:"control has-text-right"},c.a.createElement("button",{className:"button is-link",type:"button",onClick:this.addTopics},"Add")))))),c.a.createElement("div",{className:"columns"},c.a.createElement("div",{className:"column is-6 is-offset-3"},this.state.topics.map(function(t,a){return c.a.createElement("div",{className:"card",key:a,style:e},c.a.createElement("div",{className:"card-content"},c.a.createElement("p",null,t.Content)))}))))}},{key:"fetchTopics",value:function(){var e=this;f.a.get("/api/v1/topics").then(function(t){console.log(t),e.setState({topics:t.data.topics})}).catch(function(e){console.log(e)})}},{key:"addTopics",value:function(){var e=this;f()({method:"post",url:"/api/v1/topics/add",data:{content:e.state.content}}).then(function(t){e.fetchTopics()}).catch(function(e){console.log(e)})}}]),t}(n.Component),b=function(e){function t(){return Object(i.a)(this,t),Object(r.a)(this,Object(m.a)(t).apply(this,arguments))}return Object(u.a)(t,e),Object(l.a)(t,[{key:"render",value:function(){return c.a.createElement("div",{className:"App"},c.a.createElement(v,null))}}]),t}(n.Component);Boolean("localhost"===window.location.hostname||"[::1]"===window.location.hostname||window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/));s.a.render(c.a.createElement(b,null),document.getElementById("root")),"serviceWorker"in navigator&&navigator.serviceWorker.ready.then(function(e){e.unregister()})}},[[18,1,2]]]);
//# sourceMappingURL=main.39effee6.chunk.js.map