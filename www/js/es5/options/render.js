"use strict";System.register(["../../vendor/underscore","../util","./opts","lang"],function(t,e){function r(t,e){return Object.freeze(Object.defineProperties(t,{raw:{value:Object.freeze(e)}}))}function n(t,e){if(!t.length)return"";var r="";r+='<li class="tab-'+e,0===e&&(r+=" tab_sel"),r+='">';var n=!0,o=!1,l=void 0;try{for(var u,f=t[Symbol.iterator]();!(n=(u=f.next()).done);n=!0){var c=u.value;r+=i(c)}}catch(s){o=!0,l=s}finally{try{!n&&f["return"]&&f["return"]()}finally{if(o)throw l}}return 0===e&&(r+=a()),r+="</li>"}function i(t){var e="",r="shortcut"===t.type,n=t.type instanceof Array,i="checkbox"===t.type||void 0===t.type,a="number"===t.type,o="image"===t.type;r&&(e+="Alt+"),n?e+="<select":(e+="<input",i||o?e+=' type="'+(i?"checkbox":"file")+'"':a?e+=' style="width: 4em;" maxlength="4"':r&&(e+=' maxlength="1"'));var u=c(d.labels[t.id],2),f=u[0],s=u[1];if(e+=' id="'+t.id+'" title="'+s+'">',n){var v=!0,p=!1,b=void 0;try{for(var h,m=t.type[Symbol.iterator]();!(v=(h=m.next()).done);v=!0){var g=h.value;e+=l(y,g,d.modes[g]||g)}}catch(x){p=!0,b=x}finally{try{!v&&m["return"]&&m["return"]()}finally{if(p)throw b}}e+="</select>"}return e+='<label for="'+t.id+'" title="'+s+'">'+f+"</label><br>"}function a(){var t="<br>",e=["export","import","hidden"],r=!0,n=!1,i=void 0;try{for(var a,o=e[Symbol.iterator]();!(r=(a=o.next()).done);r=!0){var u=a.value,f=c(d.labels[u],2),s=f[0],y=f[1];t+='<a id="'+u+'" title="'+y+'">'+s+"</a> "}}catch(p){n=!0,i=p}finally{try{!r&&o["return"]&&o["return"]()}finally{if(n)throw i}}var b={type:"file",id:"importSettings",name:"Import Settings"};return t+=l(v,b)}var o,l,u,f,c,s,y,v,d;return t("default",function(){for(var t='<ul class="option_tab_sel">',e=d.tabs,r=[],i=function(n){if(r[n]=o(u,function(t){return t.tab===n&&(void 0===t.load||t.load)&&!t.hidden}),!r[n].length)return"continue";var i={"data-content":"tab-"+n,"class":"tab_link"};0===n&&(i["class"]+=" tab_sel"),t+=l(s,i,e[n])},a=0;a<e.length;a++){i(a)}t+='</ul><ul class="option_tab_cont">';for(var a=0;a<r.length;a++)t+=n(r[a],a);return t+="</ul>"}),{setters:[function(t){o=t.filter},function(t){l=t.parseHTML},function(t){u=t["default"]},function(t){f=t["default"]}],execute:function(){c=function(){function t(t,e){var r=[],n=!0,i=!1,a=void 0;try{for(var o,l=t[Symbol.iterator]();!(n=(o=l.next()).done)&&(r.push(o.value),!e||r.length!==e);n=!0);}catch(u){i=!0,a=u}finally{try{!n&&l["return"]&&l["return"]()}finally{if(i)throw a}}return r}return function(e,r){if(Array.isArray(e))return e;if(Symbol.iterator in Object(e))return t(e,r);throw new TypeError("Invalid attempt to destructure non-iterable instance")}}(),s=r(["<li>\n                <a ",">\n                    ","\n                </a>\n            </li>"],["<li>\n                <a ",">\n                    ","\n                </a>\n            </li>"]),y=r(['<option value="','">\n                    ',"\n                </option>"],['<option value="','">\n                    ',"\n                </option>"]),v=r(["<input ",">"],["<input ",">"]),d=f.opts}}});
//# sourceMappingURL=../maps/options/render.js.map
