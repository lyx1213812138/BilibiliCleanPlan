import get from './get.js'

const vmid = 434100110;

async function getSubscriptTags() {
  const url = 'https://api.bilibili.com/x/relation/tags';
  return (await get(url, {vmid})).data;//.map(item => {item.name, item.tagid});
  /*
  特别关注：-10
  短篇休闲：97606352
  知识：97607696
  长篇休闲：97616784
  知识：97607696
  */
}
// getSubscriptTags().then(res => console.log(res));

async function getSubscriptUpByTag(tagid) {
  const UrlSubscript = 'https://api.bilibili.com/x/relation/tag';
  // 这样写最多只能get前50个
  return (await get(UrlSubscript, {tagid})).data;
}
const tagList = [-10, 97606352, 97607696, 97616784, 97607696];
const upList = [];
tagList.forEach(tagid => {
  getSubscriptUpByTag(tagid).then(res => upList.push(...res.map(item => [item.mid, item.uname])));
});
