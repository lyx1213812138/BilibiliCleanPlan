import{d as p,r as s,c as d,w as m,o as i,a as g,b as f,e as v,f as _,g as h}from"./index-InYg6qn1.js";const b={class:"w-full flex justify-center mt-10"},S=p({__name:"Select",setup(w){let o=[];const l=s([]),t=s(JSON.parse(localStorage.getItem("selectedVgroup")||"[]").map(e=>e.mid)),r=async()=>{o=await v(),l.value=o.sort((e,a)=>a.label-e.label).map(e=>({value:e.mid,label:e.uname||e.title}))??[]},n=d(()=>(console.log("%c [ vg ]-47","font-size:13px; background:#09a78a; color:#4debce;",t.value),t.value.map(e=>o.find(a=>a.mid===e))));return m(t,()=>{localStorage.setItem("selectedVgroup",JSON.stringify(n.value))}),i(()=>{r()}),(e,a)=>{const c=_("a-transfer");return h(),g("div",b,[f(c,{"show-search":"","one-way":"",simple:"",data:l.value,"model-value":t.value,"onUpdate:modelValue":a[0]||(a[0]=u=>t.value=u),title:["未选择","已选择"],"source-input-search-props":{placeholder:"source search"},"target-input-search-props":{placeholder:"target search"}},null,8,["data","model-value"])])}}});export{S as default};
