import dealWbi from './dealWbi.js';

export default async function get(url, params) {
  const newUrl = url + '?' + await dealWbi(params);
  console.log(newUrl);
  const res = await fetch(newUrl, {
    headers: {
      "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36",
      Cookie: "buvid3=3D2BCBC6-D95A-00A4-9E0A-BA1BB9B14C8E80954infoc; b_nut=1696159880; i-wanna-go-back=-1; b_ut=7; _uuid=A2A4762B-A89D-91D4-8103D-D22DDFEAE82280810infoc; buvid4=12C023C6-4C02-2CF9-C192-30D24B5925C081860-023100119-7qfszcJeM4hXYWeQpF4pdA%3D%3D; header_theme_version=CLOSE; CURRENT_FNVAL=4048; DedeUserID=434100110; DedeUserID__ckMd5=d5026cbd6ab92890; rpdid=|(ku|u~)mR~Y0J'uYmYJJRuu~; buvid_fp_plain=undefined; enable_web_push=DISABLE; LIVE_BUVID=AUTO4417028970004144; CURRENT_QUALITY=80; hit-dyn-v2=1; fingerprint=0c5cecea22eb663dc7701128d29bdda3; PVID=2; bili_ticket=eyJhbGciOiJIUzI1NiIsImtpZCI6InMwMyIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDg3Njk1NzEsImlhdCI6MTcwODUxMDMxMSwicGx0IjotMX0.TM45NVv_RgHnxcY1tgoXWWU0JCf8wElA7CEXGa3aIBc; bili_ticket_expires=1708769511; home_feed_column=4; buvid_fp=bed0aec5a3cfca0874dac1c8f17b288c; SESSDATA=b1b9f38d%2C1724246810%2C9a502%2A21CjBhRdDC4vZQIzyIh9y_amB3Bel2j-PxNJp--HsAw1lL65ZLad4GUQCxQGk_5r0GpusSVmRUUENFOVZRcFFmLUQ0R3EtVU5XQlo5TDdVRjdnVExZWldNMU11cEZtcFljS2FtdTd5NFFpOXFYc0NvX1Iycm80MzhvbkZKeG5PM3M1N2FqOE1reUl3IIEC; bili_jct=3da59da9c9d63d7f81a18b6238fcfb61; sid=89kp6o76; bp_t_offset_434100110=901551521365229671; b_lsid=B91025957_18DDA57401F; browser_resolution=697-672"
    }
  });
  return await res.json();
}