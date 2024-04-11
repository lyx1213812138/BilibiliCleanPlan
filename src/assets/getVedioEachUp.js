import get from './get.js';

const UrlUp = 'https://api.bilibili.com/x/space/wbi/arc/search';
const UrlSeason = 'https://api.bilibili.com/x/polymer/web-space/seasons_archives_list';

export default async function getVedioEachUp(isSeason, mid, season_id) {
  // https://socialsisteryi.github.io/bilibili-API-collect/docs/user/space.html#%E6%9F%A5%E8%AF%A2%E7%94%A8%E6%88%B7%E6%8A%95%E7%A8%BF%E8%A7%86%E9%A2%91%E6%98%8E%E7%BB%86
  // https://socialsisteryi.github.io/bilibili-API-collect/docs/video/collection.html#%E8%8E%B7%E5%8F%96%E8%A7%86%E9%A2%91%E5%90%88%E9%9B%86%E4%BF%A1%E6%81%AF
  
  const data = isSeason ? 
    (await get(UrlUp, {mid})).data.list.vlist 
    : (await get(UrlSeason, {mid, season_id})).data.archives;
  // get data[i].comment length pic(图片) title in a list
  const usefulData = data.map(item => {
    return {
      bvid: item.bvid,
      comment: item.comment,
      length: isSeason ? item.duration : dealToSecond(item.length),
      pic: item.pic,
      title: item.title,
      play: isSeason ? item.stat.view : item.play,
    }
  });
  return usefulData;
}

function dealToSecond(str) {
  const time = str.split(':');
  return time[0] * 60 + time[1] * 1;
}

// test().then(res => console.log(res));
