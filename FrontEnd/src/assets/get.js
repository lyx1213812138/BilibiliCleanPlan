import dealWbi from './dealWbi.js';

export default async function get(url, params) {
  const newUrl = url + '?' + await dealWbi(params);
  console.log(newUrl);
  const res = await fetch(newUrl, {
    headers: {
      "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36",
      Cookie: "buvid3=3D2BCBC6-D95A-00A4-9E0A-BA1BB9B14C8E80954infoc; b_nut=1696159880; i-wanna-go-back=-1; b_ut=7; _uuid=A2A4762B-A89D-91D4-8103D-D22DDFEAE82280810infoc; buvid4=12C023C6-4C02-2CF9-C192-30D24B5925C081860-023100119-7qfszcJeM4hXYWeQpF4pdA%3D%3D; header_theme_version=CLOSE; CURRENT_FNVAL=4048; DedeUserID=434100110; DedeUserID__ckMd5=d5026cbd6ab92890; rpdid=|(ku|u~)mR~Y0J'uYmYJJRuu~; buvid_fp_plain=undefined; enable_web_push=DISABLE; LIVE_BUVID=AUTO4417028970004144; CURRENT_QUALITY=80; hit-dyn-v2=1; FEED_LIVE_VERSION=V8; fingerprint=2a50c62182d7aca939dcc647c7bf5eb4; buvid_fp=2a50c62182d7aca939dcc647c7bf5eb4; PVID=2; bili_ticket=eyJhbGciOiJIUzI1NiIsImtpZCI6InMwMyIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjMxMDI4NzgsImlhdCI6MTcyMjg0MzYxOCwicGx0IjotMX0.pWohPpCz5hr180Q-ODJ6It0XzkldYUM6MLWMhBRztrs; bili_ticket_expires=1723102818; SESSDATA=fb621017%2C1738395681%2C58677%2A81CjA8INPkSzFViOsnnfbR7_F2sbiSuh8LuVJ6Q7mBiHM9pL5oKDzMD-DflICPb3sPjBMSVmJkQ1RmeFJDY09fQWVYc3ZVa1JLVUxnNndja2MwaWx2NHFOZXgwa0d0Q0pqNDdYUUUzeWt0cjZYOE52czlXYWVKQU4zWU1Jd3FaNHp1MzJEY1V3aTdnIIEC; bili_jct=b2df4c4e750c50295ab11af8ba02a100; sid=6jb764hi; b_lsid=51094C2B8_19120718321; home_feed_column=4; browser_resolution=1280-661; bp_t_offset_434100110=962162245865832457"
    }
  });
  return await res.json();
}